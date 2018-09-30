package mail

import (
	"eFuture/config"
	"log"
	"net/smtp"
	"strings"
)

// Send a mail with `subject` and `content` to `receivers`.
func SendMail(subject string, receivers []string, content string) bool {
	isSent := true
	conf := config.Configuration
	auth := smtp.PlainAuth("", conf.Username, conf.Password, conf.Smtp)
	nickname := "eFuture"
	contentType := "\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n"
	msg := []byte("To: " + strings.Join(receivers, ",") + "\r\n" +
		"From: " + nickname + "<" + conf.Username + ">\r\n" +
		"Subject: " + subject +
		contentType +
		content,
	)
	err := smtp.SendMail(conf.Smtp+":"+conf.SmtpPort, auth, conf.Username, receivers, msg)
	if err != nil {
		log.Printf("Failed to send mail: %v", err)
		isSent = false
	}
	return isSent
}

// This is a example.
//func main() {
//	SendMail("test email", []string{"liuqi0315@gmail.com"}, "This is a test email, and current time is: "+time.Now().String())
//}
