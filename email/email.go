package email

import (
	"fmt"
	"net/smtp"

	"github.com/fatih/color"
	"github.com/zguillez/go-tools/core"

	"github.com/zguillez/go-tools/system"
)

func Email(emails []string, subject string, text string) {
	smtpServer := smtpServer{host: core.EmailSMTP, port: core.EmailPORT}

	from := fmt.Sprintf("From: %v", core.EmailFROM)
	to := ""
	for i := range emails {
		if i == 0 {
			to = fmt.Sprintf("%v", emails[i])
		} else {
			to = fmt.Sprintf("%v,%v", to, emails[i])
		}
	}
	to = fmt.Sprintf("To: %v", to)
	subject = fmt.Sprintf("Subject: %v", subject)

	if core.Verbose {
		color.Yellow("[email] [to] %v", to)
		color.Yellow("[email] [subject] %v", subject)
		color.Cyan("[email] [message] %v", text)
	}

	message := []byte(fmt.Sprintf("%v\n%v\n%v\n\n%v", from, to, subject, text))
	auth := smtp.PlainAuth("", from, core.EmailPASS, smtpServer.host)
	err := smtp.SendMail(smtpServer.Address(), auth, from, emails, message)
	system.CheckError(err)
	system.Echo(core.Verbose, color.Green, "[email] sended")
}

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}
