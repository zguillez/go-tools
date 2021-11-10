package email

import (
	"fmt"
	"net/smtp"

	"github.com/fatih/color"

	"github.com/zguillez/go-tools/system"
)

func Email(to []string, subject string, text string, verbose bool) {
	from := "api.zguillez.io@gmail.com"
	// password := "apizguillez2016"
	password := "kevnihpjqokaylln"
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

	_from := fmt.Sprintf("From: %v", from)
	to_ := ""
	for i := range to {
		if i == 0 {
			to_ = fmt.Sprintf("%v", to[i])
		} else {
			to_ = fmt.Sprintf("%v,%v", to_, to[i])
		}
	}
	_to := fmt.Sprintf("To: %v", to_)
	_subject := fmt.Sprintf("Subject: %v", subject)

	if verbose {
		color.Yellow("[email] [to] %v", to_)
		color.Yellow("[email] [subject] %v", subject)
		color.Cyan("[email] [message] %v", text)
	}

	message := []byte(fmt.Sprintf("%v\n%v\n%v\n\n%v", _from, _to, _subject, text))
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
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
