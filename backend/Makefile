.PHONY: tests
tests:
	@echo "Running the tests"
	@docker compose \
		-f ./build/compose.yaml \
		up simulation --build -d
	@go test ./... -coverprofile=./tools/coverage_sheet.md -v
	@docker compose \
		-f ./build/compose.yaml \
		down simulation

.PHONY: run
run:
	@docker compose \
		-f ./build/compose.yaml \
		--env-file ./config/.env \
		up app metabase --build

.PHONY: migrations
migrations:
	@docker compose \
		-f ./build/compose.yaml \
		up migrations --build

.PHONY: metabase
metabase:
	@docker compose \
		-f ./build/compose.yaml \
		up metabase --build

.PHONY: coverage
coverage: test
	@go tool cover -html=./tools/coverage_sheet.md

.PHONY: env
env: ./config/.env.develop.tmpl
	cp ./config/.env.develop.tmpl ./config/.env

.PHONY: app
app:
	@docker compose \
		-f ./build/compose.yaml \
		--env-file ./config/.env \
		up app --build

.PHONY: infra
infra:
	$(START_LOG)
	@docker compose -f ./deployments/compose.yaml up --build
	$(END_LOG)