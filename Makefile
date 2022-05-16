export PROJECT_ENV=dev
PROJECT_NAME=passwordless_verifier
DOCKER_COMPOSE_FILE=./build/docker/docker-compose.yml
DOCKER_COMPOSE_CMD=docker-compose -f ${DOCKER_COMPOSE_FILE} -p ${PROJECT_NAME}

POSTGRES_URL?=postgres://root:root@localhost:26257/postgres?sslmode=disable

.PHONY: all
all:

##
# Common section
##

.PHONY: add-pre-commit
add-pre-commit:
	lefthook add pre-commit

.PHONY: db-migrate
db-migrate:
	goose -dir ./db/migrations postgres ${POSTGRES_URL} up

##
# Go section
##

.PHONY: deps
deps:
	go mod tidy && go mod vendor && go mod verify

.PHONY: gen
gen:
	go generate ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: run
run:
	go run ./cmd/app

.PHONY: test
test:
	go test -cover ./...

.PHONY: tools
tools:
	cd ./tools && ./install.sh

##
# Docker section
##

.PHONY: container-start
container-start:
	${DOCKER_COMPOSE_CMD} up -d ${CONTAINER}

.PHONY: container-stop
container-stop:
	${DOCKER_COMPOSE_CMD} stop ${CONTAINER} && ${DOCKER_COMPOSE_CMD} rm --force ${CONTAINER}

.PHONY: jaeger-start
jaeger-start:
	CONTAINER=jaeger $(MAKE) container-start

.PHONY: jaeger-stop
jaeger-stop:
	CONTAINER=jaeger $(MAKE) container-stop

.PHONY: cockroach-start
cockroach-start:
	CONTAINER=cockroach $(MAKE) container-start

.PHONY: cockroach-stop
cockroach-stop:
	CONTAINER=cockroach $(MAKE) container-stop

.PHONY: docker-start
docker-start: jaeger-start cockroach-start

.PHONY: docker-stop
docker-stop: cockroach-stop jaeger-stop

.PHONY: docker-prune
docker-prune:
	docker system prune --volumes
