package email

import "net/smtp"

// func sendEmailHTML(subject string, templatePath string, to []string, code string) {

// 	from := "abdumalikusmonov66@gmail.com"
// 	password := "cplcvssmktfgfkzq"

// 	var body bytes.Buffer

// 	t, err := template.ParseFiles(templatePath)

// 	t.Execute(&body, struct{ Code string }{Code: code})

// 	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
// 	msg := []byte(subject + mime + body.String())

// 	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

// 	err = smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("Email Sent Successfully!")

// }

func SendMail(to []string, message []byte) error {
	from := "gofurovmurtazoxon@gmail.com"
	password := "ihtizrusjypnultt"

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	return err
}
