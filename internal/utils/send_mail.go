package utils

import "net/smtp"

func SendEmail(msg []byte, email string) error {
	// from := os.Getenv("MAIL")
	// password := os.Getenv("PASSWD")

	from := "phuocnghiaqs7@gmail.com"
	password := "lewxbqwmtbvayvek"
	// toList is list of email address that email is to be sent.
	toList := []string{email}

	// host is address of server that the
	// sender's email address belongs,
	// in this case its gmail.
	// For e.g if your are using yahoo
	// mail change the address as smtp.mail.yahoo.com
	host := "smtp.gmail.com"

	// Its the default port of smtp server
	port := "587"

	// We can't send strings directly in mail,
	// strings need to be converted into slice bytes

	// PlainAuth uses the given username and password to
	// authenticate to host and act as identity.
	// Usually identity should be the empty string,
	// to act as username.
	auth := smtp.PlainAuth("", from, password, host)

	// SendMail uses TLS connection to send the mail
	// The email is sent to all address in the toList,
	// the body should be of type bytes, not strings
	// This returns error if any occurred.
	err := smtp.SendMail(host+":"+port, auth, from, toList, msg)
	// handling the errors
	if err != nil {
		return err
	}

	return nil
}
