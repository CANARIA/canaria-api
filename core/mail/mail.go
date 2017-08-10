package mail

import (
	"github.com/CANARIA/canaria-api/core/config"
	"github.com/CANARIA/canaria-api/core/message"
	"github.com/CANARIA/canaria-api/core/model"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
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
func (mail *Mail) Send(ctx context.Context) {
	log.Infof(ctx, "Mail To: %s\n", mail.To)
	log.Infof(ctx, "Mail CC: %s\n", mail.Cc)
	log.Infof(ctx, "Mail BCC: %s\n", mail.Bcc)
	log.Infof(ctx, "Mail FROM: %s\n", mail.From)
	log.Infof(ctx, "Mail SUBJECT: %s\n", mail.Subject)
	log.Infof(ctx, "Mail BODY: %s\n", mail.Body)

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
