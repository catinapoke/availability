PROJECT_NAME=availability_checker

run-build:
	docker compose -p $(PROJECT_NAME) up --build

run:
	docker compose -p $(PROJECT_NAME) up -d

stop:
	docker compose -p $(PROJECT_NAME) down