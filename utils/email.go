package utils

import (
	"fmt"
	"net/smtp"
)

// TODO: ホストの情報は環境変数から取得に変更
func SendVerificationEmail(email, token string) error {
	// SMTPサーバーの設定
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	from := "your-email@gmail.com"    // 送信者のメールアドレス
	password := "your-email-password" // 送信者のメールアドレスのパスワード

	// 受信者のメールアドレス
	to := []string{email}

	// メールの内容
	subject := "Email Verification"
	body := fmt.Sprintf("Please verify your email by clicking the link: http://localhost:8080/verify?token=%s", token)
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
