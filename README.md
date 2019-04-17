# 使用Go语言通过SMTP协议实现发送邮件功能  

## example  

```go
package main

import (
    "bytes"
    "log"
    "net/smtp"
)

func main() {
    // Connect to the remote SMTP server.
    c, err := smtp.Dial("smtp.gmail.com:465")
    if err != nil {
        log.Fatal(err)
    }
    defer c.Close()
    // Set the sender and recipient.
    c.Mail("dhal.asitk@gmail.com")
    c.Rcpt("asitdhal_tud@outloo.com")
    // Send the email body.
    wc, err := c.Data()
    if err != nil {
        log.Fatal(err)
    }
    defer wc.Close()
    buf := bytes.NewBufferString("This is the email body.")
    if _, err = buf.WriteTo(wc); err != nil {
        log.Fatal(err)
    }
}

/*
Don’t try this. This doesn’t work for Gmail. Gmail needs password for login and SSL/TLS for end-to-end security.
*/
```



