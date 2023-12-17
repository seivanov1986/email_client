package email_client

type Email interface {
	SetCreeds(
		emailHost string,
		emailPort int,
		emailUser string,
		emailPassword string,
	)
	SetBody(body string)
	SetSubject(subject string)
	SetTo(to string)
	SetFrom(from string)
	SetAttachments(attachments map[string]string)
	Send()
}
