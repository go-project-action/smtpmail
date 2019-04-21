package test

import (
	"fmt"
	"log"
	"net/smtp"
)

func Example() {
	// 连接远程的smtp服务器.
	c, err := smtp.Dial("mail.example.com:25")
	if err != nil {
		log.Fatal(err)
	}

	// 首先，保存发送方和接收方
	if err := c.Mail("mail.example.com"); err != nil {
		log.Fatal(err)
	}

	if err := c.Rcpt("mail.example.com"); err != nil {
		log.Fatal(err)
	}

	// 发送email数据.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintf(wc, "This is the email body")
	if err != nil {
		log.Fatal(err)
	}

	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 关闭连接
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
}

var (
	from       = "gopher@example.net"
	msg        = []byte("dummy message")
	recipients = []string{"foo@example.com"}
)

func ExamplePlainAuth() {

	// hostname is used by PlainAuth to validate the TLS certificate.
	hostname := "mail.example.com"
	auth := smtp.PlainAuth("", "user@example.com", "password", hostname)
	err := smtp.SendMail(hostname+":25", auth, from, recipients, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleSendMail() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"recipient@example.net"}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")

	err := smtp.SendMail("mail.example.com:25", auth, "sender@example.org", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
