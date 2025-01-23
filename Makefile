.PHONY: compose
compose:
	@docker compose up -d --build

.PHONE: protoc
protoc:
	@bin/protoc
