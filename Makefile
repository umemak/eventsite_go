.PHONY: build
build:
	docker compose build

.PHONY: up
up:
	docker compose up -d --remove-orphans --wait
	npx prisma migrate dev

.PHONY: down
down:
	docker compose down --volumes --remove-orphans

.PHONY: dev
dev:
	docker compose exec dev /bin/bash
