.PHONY: compose
compose:
	@docker compose up -d --build

.PHONY: compose-fast
compose-fast:
	@docker compose up -d

.PHONY: test
test:
	@go test ./...
