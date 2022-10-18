package email

import (
	"gopkg.in/gomail.v2"
	"pandownload/settings"
)

var (
	d             *gomail.Dialer
	emailDefaults = make(map[string][]string, 2)
)

func Init() {
	email := settings.EmailSetting()
	d = gomail.NewDialer(email.Host, email.Port, email.From, email.AccessToken)
	emailDefaults["From"] = append(emailDefaults["From"], email.From)
	emailDefaults["Subject"] = append(emailDefaults["Subject"], email.Subject)
	return
}

func SendmailTo(To []string, body string) error {
	m := gomail.NewMessage()
	emailDefaults["To"] = To
	m.SetBody("text/plain", body)
	m.SetHeaders(emailDefaults)
	err := d.DialAndSend(m)
	return err
}
