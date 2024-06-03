include .env

# マイグレーションを実行するタスク
migrate-up:
	migrate -database ${POSTGRESQL_URL} -path migrations up

# マイグレーションをロールバックするタスク
migrate-down:
	migrate -database ${POSTGRESQL_URL} -path migrations down

# 新しいマイグレーションファイルを作成するタスク
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq "$$name"


