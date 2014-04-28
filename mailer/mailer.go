package mailer

import (
	"errors"
	"net/smtp"
	"github.com/scorredoira/email"
	//"../logger"
	"log"
)

func SendMail(subject string, mailbody string, to []string, pathtoattachment string) {
	// Send mail with attachment. If you don't want to send attachment just pass "".	

	panic(errors.New("mailer.SendMail not implemented"))
	//sendmail(subject, mailbody, "cannot-reply@cogmed.com", to, "cannot-reply@cogmed.com", "7384C9", "smtp.gmail.com", pathtoattachment)
}

func SendError(mailbody string) {
	// Send error message and attach the geomys.log log file
	panic(errors.New("mailer.SendError not implemented."))
	//sendmail("Geomys error", mailbody, "cannot-reply@cogmed.com", []string{"michael.smietana@pearson.com"}, "cannot-reply@cogmed.com", "7384C9", "smtp.gmail.com", "geomys.log")
}

type SendMailCreds struct {
	From 				string
	To 					[]string
	Cc 					[]string
	Bcc 				[]string
	Subject 			string
	Message 			string
	SMTPSrvAddr 		string
	SMTPSrvPort			string
	AuthUsrName 		string
	AuthPwd 			string
	PathToAttachment 	string
}

//func sendmail(subject string, body string, from string, to []string, username string, password string, serveraddress string, pathtoattachment string) {
func sendmail(creds SendMailCreds) error {
	// send email

	m := email.NewMessage(creds.Subject, creds.Message)
	m.From = creds.From
	m.To = creds.To
	if len(creds.Cc) != 0 {
		m.Cc = creds.Cc
	}
	if len(creds.Bcc) != 0 {
		m.Bcc = creds.Bcc
	}
	if creds.PathToAttachment != "" {
		err := m.Attach(creds.PathToAttachment)
		if err != nil {
			//logger.LogWarning("error when attaching mail attachment. error: ", err)
			log.Println("error when attaching mail attachment. error: ", err)
		}
	}

	if creds.SMTPSrvPort == "" {
		creds.SMTPSrvPort = "587"
	}
	err := email.Send((creds.SMTPSrvAddr + ":" + creds.SMTPSrvPort), smtp.PlainAuth("", creds.AuthUsrName, creds.AuthPwd, creds.SMTPSrvAddr), m)
	if err != nil {
		//logger.LogWarning("error when sending mail. error: ", err)
		log.Println("error when sending mail. error: ", err)
		return err
	}
	return nil
}
