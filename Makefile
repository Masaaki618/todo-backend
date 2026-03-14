.PHONY: up down build rebuild logs

up:
	docker compose up

down:
	docker compose down

build:
	docker compose build

rebuild:
	docker compose up --build

logs:
	docker compose logs -f app
