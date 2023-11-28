boot:
	docker-compose up -d --remove-orphans
down:
	docker-compose down -v --remove-orphans
.PHONY: boot down