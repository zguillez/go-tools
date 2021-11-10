package email

import (
	"fmt"
	"net/smtp"

	"github.com/fatih/color"

	"zguillez.io/gotools/system"
)

func Email(emails []string, subject string, text string, from string, password string, verbose bool) {
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

	from = fmt.Sprintf("From: %v", from)
	to := ""
	for i := range to {
		if i == 0 {
			to = fmt.Sprintf("%v", emails[i])
		} else {
			to = fmt.Sprintf("%v,%v", to, emails[i])
		}
	}
	to = fmt.Sprintf("To: %v", to)
	subject = fmt.Sprintf("Subject: %v", subject)

	if verbose {
		color.Yellow("[email] [to] %v", to)
		color.Yellow("[email] [subject] %v", subject)
		color.Cyan("[email] [message] %v", text)
	}

	message := []byte(fmt.Sprintf("%v\n%v\n%v\n\n%v", from, to, subject, text))
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	err := smtp.SendMail(smtpServer.Address(), auth, from, emails, message)
	system.CheckError(err)
	if verbose {
		color.Green("[email] sended")
	}
}

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}
