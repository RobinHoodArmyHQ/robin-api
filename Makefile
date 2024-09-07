-include .env.development
export

dev:
	docker-compose -f docker-compose.yaml up -d --build

clean:
	docker ps -q --filter "name=${SERVICE_NAME}" | xargs docker container stop

logs:
	docker ps --filter "name=${SERVICE_NAME}$$" --format "{{.ID}}" | xargs docker logs -f --tail 10

migrate:
	docker run -v ${PWD}/db:/migrations --network host \
		migrate/migrate \
		-path=/migrations/ \
		-database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}" \
		up

migrate-revert:
	@read -p "Enter the number of migrations to revert: " count; \
	docker run -v ${PWD}/db:/migrations --network host \
		migrate/migrate \
		-path=/migrations/ \
		-database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}" \
		down $$count

create-migration:
	docker run -v ${PWD}/db:/migrations --network host \
		migrate/migrate \
		create -ext sql -dir /migrations -seq -digits 3 $(name)

test-data:
	docker run -v ${PWD}/testdata/db:/migrations --network host \
		migrate/migrate \
		-path=/migrations/ \
		-database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}?x-migrations-table=testdata_migration" \
		up