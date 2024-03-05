.PHONY:
broker:
	@docker compose \
		-f ./compose.yaml \
		up -d broker

.PHONY: test
test: broker
	@echo "Running the tests"
	@go test ./... -coverprofile=coverage_sheet.md
	@docker compose \
		-f ./compose.yaml \
		down broker

.PHONY: run
run:
	@docker compose \
		-f ./build/compose.yaml \
		up --build

.PHONY: coverage
coverage: test
	@go tool cover -html=./utils/coverage_sheet.md

.PHONY: env
env: ./config/.env.tmpl
	cp ./config/.env.tmpl ./config/.env