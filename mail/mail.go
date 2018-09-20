//package mail
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"strings"
	"time"
)

type configuration struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Smtp     string `json:"smtp"`
}

var conf configuration

func init() {
	//workDir, _ := os.Getwd()
	bytes, e := ioutil.ReadFile("/home/xlui/eFuture/config.json")
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
	contentType := "Content-Type: text/plain; charset=UTF-8"
	msg := []byte(
		"To: " + strings.Join(receivers, ",") + "\r\n" +
		"From: " + nickname + "<" + conf.Username + ">\r\n" +
		"Subject: " + subject + "\r\n" +
		contentType + "\r\n\r\n" +
		content,
	)
	err := smtp.SendMail(conf.Smtp+":587", auth, conf.Username, receivers, msg)
	if err != nil {
		fmt.Printf("Failed to send mail: %v", err)
		isSent = false
	}
	return isSent
}

func main() {
	SendMail("test email", []string{"liuqi0315@gmail.com"}, "This is a test email, current is: " + time.Now().String())
}
