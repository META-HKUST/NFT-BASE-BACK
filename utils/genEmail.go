package utils

import (
	"github.com/matcornic/hermes/v2"
	"github.com/wneessen/go-mail"
	"log"
)

var (
	FrontLink string = ("http://unifit.ust.hk:8889/api/v2/user/activate?token=")

	SubjectName  string = ("Activate your uNiFiT account!")
	SubjectName2 string = ("Reset your password")
	logo         string = ("https://unift-1312994969.cos.ap-guangzhou.myqcloud.com/unifit/logo_0725.png?q-sign-algorithm=sha1&q-ak=AKIDBD4i9ML5aswlLgmfJisnTt30f6JJ6duu&q-sign-time=1658731709%3B1745131709&q-key-time=1658731709%3B1745131709&q-header-list=host&q-url-param-list=&q-signature=596694ad14acc8698b0e1fe7d6d15ac1131c4019")

	SenderName   string      = ("uNiFiT Team")
	Sender       string      = ("contact@unifit.art")
	SenderServer string      = ("smtp.feishu.cn")
	SenderPort   int         = 587
	SenderPasswd string      = ("eWwBcBBqeQIHCft0")
	WithSSL      mail.Option = mail.WithSSL()
)

func Email(ReceiverName string, ReceiverMail string, token string) error {
	//---------------------------------form html-----------------------------------
	h := hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: SenderName,
			Link: FrontLink + token,
			// Optional product logo
			Logo: logo,
		},
	}
	email := hermes.Email{
		Body: hermes.Body{
			Name: ReceiverName,
			Intros: []string{
				"Welcome to uNiFiT! We're very excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To verify your email address, please click here:",
					Button: hermes.Button{
						Text: "Get Verified",
						Link: FrontLink + token,
					},
				},
			},
			Outros: []string{
				"Then you can sign in to your account at https://unifit.ust.hk/login. ",
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		return err
	}
	//-----------------------------------------------------------------------
	//---------------------send email body-----------------------------------
	//-----------------------------------------------------------------------
	// Create a new mail message
	m := mail.NewMsg()

	if err := m.FromFormat(SenderName, Sender); err != nil {
		return err
	}
	if err := m.To(ReceiverMail); err != nil {
		return err
	}

	// Set a subject line
	m.Subject(SubjectName)

	// Add your mail message to body
	m.SetBodyString(mail.TypeTextHTML, emailBody)

	host := SenderServer
	c, err := mail.NewClient(host,
		mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithUsername(Sender),
		mail.WithPassword(SenderPasswd), mail.WithTLSPolicy(mail.TLSMandatory), mail.WithPort(SenderPort))

	if err != nil {
		log.Println(err)
		return err
	}
	// Now that we have our client, we can connect to the server and send our mail message
	// via the convenient DialAndSend() method. You have the option to Dial() and Send()
	// seperately as well
	if err := c.DialAndSend(m); err != nil {
		return err
	}
	log.Println("Activation email successfully sent to", ReceiverName)
	return nil
}

func ResetEmail(ReceiverName string, ReceiverMail string, code string) error {
	h := hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: SenderName,
			// Optional product logo
			Logo: logo,
		},
	}
	email := hermes.Email{
		Body: hermes.Body{
			Name: ReceiverName,
			Intros: []string{
				"Your verification code is: ",
				code,
			},
			Outros: []string{
				"The verification code is valid for 15 minutes, please do not disclose it to others!",
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}
	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		return err
	}
	//-----------------------------------------------------------------------
	//---------------------send email body-----------------------------------
	//-----------------------------------------------------------------------
	// Create a new mail message
	m := mail.NewMsg()

	if err := m.FromFormat(SenderName, Sender); err != nil {
		return err
	}
	if err := m.To(ReceiverMail); err != nil {
		return err
	}

	// Set a subject line
	m.Subject(SubjectName)

	// Add your mail message to body
	m.SetBodyString(mail.TypeTextHTML, emailBody)

	host := SenderServer
	c, err := mail.NewClient(host,
		mail.WithSMTPAuth(mail.SMTPAuthLogin), mail.WithUsername(Sender),
		mail.WithPassword(SenderPasswd), mail.WithTLSPolicy(mail.TLSMandatory), mail.WithPort(SenderPort))

	if err != nil {
		return err
	}
	// Now that we have our client, we can connect to the server and send our mail message
	// via the convenient DialAndSend() method. You have the option to Dial() and Send()
	// seperately as well
	if err := c.DialAndSend(m); err != nil {
		return err
	}
	log.Println("Reset email successfully sent to", ReceiverName)
	return nil
}
