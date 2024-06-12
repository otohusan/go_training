package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendVerificationEmail(email, token string) error {
	// 環境変数からSMTP設定を読み込む
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	// 受信者のメールアドレス
	to := []string{email}

	// メールの内容
	subject := "本登録のご案内<自動送信>"
	// TODO: ここでローカルホストのURLを使ってるから直す
	body := fmt.Sprintf("Konwalk(コンウォーク)に仮登録いただききありがとうございます。\n下記のリンクをクリックし、本登録を完了させてください。\n\nhttp://localhost:8080/verify?token=%s", token)
	message := fmt.Sprintf("Subject: %s\n\n%s", subject, body)

	// 認証情報
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// メールの送信
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
