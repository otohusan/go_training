# マイグレーションを実行するタスク
migrate-up:
	go run cmd/migrate/main.go up

# マイグレーションをロールバックするタスク
migrate-down:
	go run cmd/migrate/main.go down

# 適切なDBPathが返ってくるか試せる
migrate-test:
	go run cmd/migrate/main.go test

# 新しいマイグレーションファイルを作成するタスク
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq "$$name"

# 開発段階でのサーバ起動
run-dev:
	go run cmd/main.go

# 製品段階でのサーバ起動
run-release:
	GIN_MODE=release go run cmd/main.go


