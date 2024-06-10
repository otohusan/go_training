package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func loadEnvAndMigrate(action string) {
	// モードを尋ねる
	var mode string
	fmt.Print("Enter mode (development/release): ")
	fmt.Scan(&mode)

	// 環境変数ファイルの設定
	if mode != "release" && mode != "development" {
		fmt.Print("入力が正しくありません")
		return
	}
	envFile := ".env.development"
	if mode == "release" {
		envFile = ".env.release"
	}

	// .env ファイルを読み込む
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}

	// 環境変数からPOSTGRESQL_URLを取得
	postgresURL := os.Getenv("POSTGRESQL_URL")
	if postgresURL == "" {
		log.Fatalf("POSTGRESQL_URL is not set in %s file", envFile)
	}

	if action == "test" {
		fmt.Print(postgresURL)
		return
	}

	// migrate コマンドを実行
	cmd := exec.Command("migrate", "-database", postgresURL, "-path", "migrations", action)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("migrate command failed: %v", err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/migrate/main.go [up|down]")
		return
	}
	action := os.Args[1]
	if action != "up" && action != "down" && action != "test" {
		fmt.Println("Invalid action. Use 'up' or 'down'.")
		return
	}
	loadEnvAndMigrate(action)
}
