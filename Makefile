scripts_dir := ./scripts

all: init install

dev:
	docker compose -f ./deployment/compose.development.yml up --build

install:
	go mod download
	go mod verify
	pnpm --dir web install

init:
	$(scripts_dir)/init.sh
