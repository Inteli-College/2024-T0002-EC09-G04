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
		--env-file ./config/.env \
		up simulation app metabase --build

.PHONY: migrations
migrations:
	@docker compose \
		-f ./build/compose.yaml \
		up migrations --build

.PHONY: coverage
coverage: test
	@go tool cover -html=./utils/coverage_sheet.md

.PHONY: env
env: ./config/.env.develop.tmpl
	cp ./config/.env.develop.tmpl ./config/.env