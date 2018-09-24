package mail

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"strings"
)

type configuration struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Smtp     string `json:"smtp"`
	SmtpPort string `json:"smtp_port"`
}

var conf configuration

func init() {
	config := os.Getenv("EFUTURE_CONFIG")
	if config == "" {
		config = "/data/eFuture/config.json"
	}
	bytes, e := ioutil.ReadFile(config)
	if e != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", e)
		return
	}
	json.Unmarshal(bytes, &conf)
}

func SendMail(subject string, receivers []string, content string) bool {
	isSent := true
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
		fmt.Printf("Failed to send mail: %v", err)
		isSent = false
	}
	return isSent
}

//func main() {
//	SendMail("test email", []string{"liuqi0315@gmail.com"}, "This is a test email, and current time is: "+time.Now().String())
//}
