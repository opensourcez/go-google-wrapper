package googlewrapper

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

type Email struct {
	FromName  string
	FromEmail string
	ToName    string
	ToEmail   string
	Subject   string
	Message   string
	Token     string
	Secret    string
}

func GmailConfigFromFile(file string, scope ...string) (config *oauth2.Config, err error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return config, err
	}
	config, err = google.ConfigFromJSON(dat, gmail.GmailSendScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	return config, err
}

func GetGmailService(secret string, token string) (*gmail.Service, error) {
	config, err := GmailConfigFromFile(secret)
	if err != nil {
		panic(err)
	}
	cl := GetClient(config, token)
	return gmail.New(cl)
}
func (em *Email) SendMailFromEmail() error {

	gmailService, err := GetGmailService(em.Secret, em.Token)
	if err != nil {
		panic(err)
	}
	from := mail.Address{em.FromName, em.FromEmail}
	to1 := mail.Address{em.ToName, em.ToEmail}

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to1.String()
	header["Subject"] = "=?utf-8?B?" + base64.RawURLEncoding.EncodeToString([]byte(em.Subject)) + "?="
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html"
	header["Content-Transfer-Encoding"] = "base64"

	var msg string
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + em.Message

	gmsg := gmail.Message{
		Raw: base64.RawURLEncoding.EncodeToString([]byte(msg)),
	}

	_, err = gmailService.Users.Messages.Send("me", &gmsg).Do()
	if err != nil {
		log.Printf("em %v, err %v", gmsg, err)
		return err
	}
	return err
}
