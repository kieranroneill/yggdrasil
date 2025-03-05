.PHONY: install

all: install

dev:
	docker compose -f compose.development.yml up

install:
	@go install github.com/conventionalcommit/commitlint@latest
