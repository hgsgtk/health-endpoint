.PHONY: up
up:
	docker-compose up -d

.PHONY: build
build:
	docker-compose build health-api
