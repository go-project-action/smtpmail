package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type Mail struct {
	senderId string   //发送信息的id
	toIds    []string //接收信息的id
	subject  string
	body     string
}

type SmtpServer struct {
	host string
	port string
}

func (s *SmtpServer) ServerName() string {
	return s.host + ":" + s.port
}

func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From:%s\r\n", mail.senderId)
	if len(mail.toIds) > 0 {
		message += fmt.Sprintf("To:%s\r\n", strings.Join(mail.toIds, ";"))
	}

	message += fmt.Sprintf("Subject:%s\r\n", mail.subject)
	message += "\r\n" + mail.body

	return message
}

func main() {
	mail := Mail{}
	mail.senderId = "zoctopus@qq.com"
	mail.toIds = []string{"782435935@qq.com"}
	mail.subject = "This is the email subject"
	mail.body = "hello world"

	messageBody := mail.BuildMessage()

	smtpServer := SmtpServer{host: "smtp.qqmail.com", port: "465"}
	log.Println(smtpServer.host)

	auth := smtp.PlainAuth("", mail.senderId, "password", smtpServer.host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.host,
	}

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		log.Panic(err)
	}

	//1,使用Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	//2,添加发送者和接收者
	if err = client.Mail(mail.senderId); err != nil {
		log.Panic(err)
	}
	for _, k := range mail.toIds {
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	//3,传数据
	w, err := client.Data()
	if err != nil {
		log.Panic()
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		log.Panic()
	}

	err = w.Close()
	if err != nil {
		log.Panic()
	}

	client.Quit()

	log.Println("Mail sent success!")

}
