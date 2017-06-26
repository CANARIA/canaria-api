package mail

import (
	"fmt"

	"github.com/CANARIA/canaria-api/core/config"
	"github.com/CANARIA/canaria-api/core/message"
	"github.com/CANARIA/canaria-api/core/model"
)

type (
	Mail struct {
		To      string
		From    string
		Cc      string
		Bcc     string
		Subject string
		Body    string
	}
)

func BuildPreRegisterMail(preAccount model.PreAccount, url string) *Mail {

	body := message.PRE_REGISTER_MAIL_BODY + url
	return &Mail{
		To:      preAccount.MailAddress,
		From:    config.GetMailAddress(),
		Subject: message.PREREGISTER_SUBJECT,
		Body:    body,
	}

}

func BuildRegisteredMail(preAccount *model.PreAccount) *Mail {
	return &Mail{
		To:      preAccount.MailAddress,
		From:    config.GetMailAddress(),
		Subject: message.REGISTER_SUBJECT,
		Body:    message.REGISTER_MAIL_BODY,
	}
}

func (mail *Mail) Send() {
	fmt.Printf("Mail To: %s\n", mail.To)
	fmt.Printf("Mail CC: %s\n", mail.Cc)
	fmt.Printf("Mail BCC: %s\n", mail.Bcc)
	fmt.Printf("Mail FROM: %s\n", mail.From)
	fmt.Printf("Mail SUBJECT: %s\n", mail.Subject)
	fmt.Printf("Mail BODY: %s\n", mail.Body)

	// Connect to the remote SMTP server.
	// c, err := smtp.Dial("localhost:25")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // Set the sender and recipient.
	// c.Mail("sender@localhost")     // メールの送り主を指定
	// c.Rcpt("ganimata39@gmail.com") // 受信者を指定
	// // c.Rcpt("yyyyyy@gmail.com") // Ccにしたい場合も同様に指定

	// // Send the email body.
	// wc, err := c.Data()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer wc.Close()
	// //ToにするかCcにするかBccにするかはDATAメッセージ次第
	// buf := bytes.NewBufferString("To:ganimata39@gmail.com")
	// buf.WriteString("\r\n") // DATA メッセージはCRLFのみ
	// // buf.WriteString("Cc:yyyyyy@gmail.com")
	// // buf.WriteString("\r\n")
	// buf.WriteString("Subject:this is Subject") //件名
	// buf.WriteString("\r\n")
	// buf.WriteString("This is the email body.")
	// if _, err = buf.WriteTo(wc); err != nil {
	// 	log.Fatal(err)
	// }

	// c.Quit() //メールセッションの終了
}
