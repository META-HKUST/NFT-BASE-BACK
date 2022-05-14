package utils

import (
	"github.com/matcornic/hermes/v2"
	"github.com/wneessen/go-mail"
	"log"
)

var (
	FrontLink string = ("http://localhost:8888/api/v1/users/activate?token=")

	SubjectName string = ("Fire Bird Account Confirmation")
	logo        string = ("https://thumbs.dreamstime.com/z/phoenix-fenix-fire-bird-cartoon-character-fantasy-magic-creature-red-burning-plumage-steaming-wings-fairytale-animal-182971826.jpg")

	SenderName   string      = ("MetaUST Account Center")
	Sender       string      = ("1721062927@qq.com")
	SenderServer string      = ("smtp.qq.com")
	SenderPort   int         = 465
	SenderPasswd string      = ("rqdkrrkzsubnedfg")
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
				"Welcome to Fire Bird NFT! We're very excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with Fire Brid, please click here:",
					Button: hermes.Button{
						Text: "Confirm your account",
						Link: FrontLink + token,
					},
				},
			},
			Outros: []string{
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
		mail.WithPassword(SenderPasswd), mail.WithTLSPolicy(mail.TLSMandatory), mail.WithPort(SenderPort), WithSSL)

	if err != nil {
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
