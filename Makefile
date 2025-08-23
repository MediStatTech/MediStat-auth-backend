VERSION=v1.0.2

# SQLC generation
sqlc-run:
	sqlc generate

sqlc-verify:
	sqlc verify

# Install dependencies
install-tools:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Новый способ через отдельный репозиторий
sqlc-generate:
	go install github.com/MediStatTech/MediStat-generate-model/cmd/sqlc-gen@latest
	sqlc-gen

# Clean
clean:
	rm -rf bin/
	rm -rf internal/repository/*.go

# Linting
lint:
	golangci-lint run

# Release
release:
	echo "Only create it after you push the changes to the repository # master"
	go mod tidy; git add .; git commit -m "Release $(VERSION)"; git push origin master; git tag $(VERSION); git push origin $(VERSION);

.PHONY: build run migrate-up migrate-down migrate-create sqlc-run sqlc-generate sqlc-verify sqlc-config sqlc-repo generate install-tools clean test test-coverage lint fmt release