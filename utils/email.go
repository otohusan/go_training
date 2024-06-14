package utils

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

func SendVerificationEmail(email, token string) error {
	// 環境変数からSMTP設定を読み込む
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	myURL := os.Getenv("MY_URL")

	// 受信者のメールアドレス
	to := []string{email}

	// メールの内容
	subject := "本登録のご案内<自動送信>"
	// TODO: ここでローカルホストのURLを使ってるから直す
	body := fmt.Sprintf("Konwalk(コンウォーク)に仮登録いただききありがとうございます。\n下記のリンクをクリックし、本登録を完了させてください。\n\n%s/verify?token=%s", myURL, token)

	// メールのメッセージを作成（エンコーディングの設定を含む）
	header := make(map[string]string)
	header["From"] = from
	header["To"] = strings.Join(to, ",")
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"UTF-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// 認証情報
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// メールの送信
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
