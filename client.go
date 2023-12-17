package email_client

type email struct {
	emailHost     string
	emailPort     int
	emailUser     string
	emailPassword string
	to            string
	body          string
	attachments   map[string]string
	subject       string
	from          string
}

func NewEmailClient() *email {
	return &email{}
}
