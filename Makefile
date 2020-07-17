up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

logs:
	docker logs -f app

fmt:
	go fmt ./...