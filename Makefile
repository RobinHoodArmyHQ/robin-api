SERVICE_NAME=robin-api

dev:
	docker-compose -f docker-compose.yaml up -d --build

clean:
	docker ps -q --filter "name=${SERVICE_NAME}" | xargs docker container stop

logs:
	docker ps --filter "name=${SERVICE_NAME}$$" --format "{{.ID}}" | xargs docker logs -f --tail 10
