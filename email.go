package email_client

import (
	b64 "encoding/base64"
	"fmt"
	"time"

	mail "github.com/xhit/go-simple-mail"
)

func (e *email) SetCreeds(
	emailHost string,
	emailPort int,
	emailUser string,
	emailPassword string,
) {
	e.emailHost = emailHost
	e.emailPort = emailPort
	e.emailUser = emailUser
	e.emailPassword = emailPassword
}

func (e *email) SetBody(body string) {
	e.body = body
}

func (e *email) SetSubject(subject string) {
	e.subject = subject
}

func (e *email) SetTo(to string) {
	e.to = to
}

func (e *email) SetFrom(from string) {
	e.from = from
}

func (e *email) SetAttachments(attachments map[string]string) {
	e.attachments = attachments
}

func (e *email) Send() {
	server := mail.NewSMTPClient()
	server.Host = e.emailHost
	server.Port = e.emailPort
	server.Username = e.emailUser
	server.Password = e.emailPassword
	server.Encryption = mail.EncryptionTLS
	server.KeepAlive = false
	server.ConnectTimeout = 100 * time.Second
	server.SendTimeout = 100 * time.Second
	smtpClient, err := server.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	email := mail.NewMSG()
	email.SetFrom(e.from + " <" + e.emailUser + ">").
		AddTo(e.to).
		SetSubject(e.subject)

	email.SetBody(mail.TextHTML, e.body)

	for name, body := range e.attachments {
		email.AddAttachmentBase64(b64.StdEncoding.EncodeToString([]byte(body)), name)
	}

	email.Send(smtpClient)
}
