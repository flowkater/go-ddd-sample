DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
DB_CONN=postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable
ENV=dev

TEST_DB_HOST=test_db
TEST_DB_PORT=5432
TEST_DB_USER=postgres
TEST_DB_PASSWORD=postgres
TEST_DB_NAME=postgres
TEST_DB_CONN=postgres://${TEST_DB_USER}:${TEST_DB_PASSWORD}@${TEST_DB_HOST}:${TEST_DB_PORT}/${TEST_DB_NAME}?sslmode=disable


.PHONY: run
run:
	docker-compose up --build -d

.PHONY: start
start:
	docker-compose exec -e ENV=${ENV} -e DB_HOST=${DB_HOST} -e DB_PORT=${DB_PORT} -e DB_USER=${DB_USER} -e DB_PASSWORD=${DB_PASSWORD} -e DB_NAME=${DB_NAME} app realize start --run 

.PHONY: db-console
db-console:
	docker-compose exec db psql ${DB_USER} ${DB_PASSWORD}

.PHONY: generate
generate:
	docker-compose exec app go generate ./...

.PHONY: migrate-create
migrate-create:
	docker-compose exec app migrate create -ext sql -dir migrations ${FILENAME}

.PHONY: migrate-up
migrate-up:
	docker-compose exec app migrate --source file://migrations --database ${DB_CONN} up

.PHONY: migrate-down
migrate-down:
	docker-compose exec app migrate --source file://migrations --database ${DB_CONN} down 1

.PHONY: migrate-force
migrate-force:
	docker-compose exec app migrate --source file://migrations --database ${DB_CONN} force ${VERSION}

.PHONY: test-migrate-up
test-migrate-up:
	docker-compose exec app migrate --source file://migrations --database ${TEST_DB_CONN} up

.PHONY: test-migrate-down
test-migrate-down:
	docker-compose exec app migrate --source file://migrations --database ${TEST_DB_CONN} down 1

.PHONY: test-migrate-force
test-migrate-force:
	docker-compose exec app migrate --source file://migrations --database ${TEST_DB_CONN} force ${VERSION}

.PHONY: test-convey
test-convey:
	docker-compose exec -e ENV=${ENV} -e TEST_DB_HOST=${TEST_DB_HOST} -e TEST_DB_PORT=${TEST_DB_PORT} -e TEST_DB_USER=${TEST_DB_USER} -e TEST_DB_PASSWORD=${TEST_DB_PASSWORD} -e TEST_DB_NAME=${TEST_DB_NAME} app ../go/bin/goconvey -host 0.0.0.0

.PHONY: test-all
test-all:
	docker-compose exec -e ENV=${ENV} -e TEST_DB_HOST=${TEST_DB_HOST} -e TEST_DB_PORT=${TEST_DB_PORT} -e TEST_DB_USER=${TEST_DB_USER} -e TEST_DB_PASSWORD=${TEST_DB_PASSWORD} -e TEST_DB_NAME=${TEST_DB_NAME} app go test -v ./...

.PHONY: test
test:
	docker-compose exec -e ENV=${ENV} -e TEST_DB_HOST=${TEST_DB_HOST} -e TEST_DB_PORT=${TEST_DB_PORT} -e TEST_DB_USER=${TEST_DB_USER} -e TEST_DB_PASSWORD=${TEST_DB_PASSWORD} -e TEST_DB_NAME=${TEST_DB_NAME} app go test -v ./${PACKAGE}

# .PHONY: local-test
# local-test:
# 	sudo -E ~/go/bin/goconvey