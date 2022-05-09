.PHONY: all
all:

##
# Common section
##

.PHONY: add-pre-commit
add-pre-commit:
	lefthook add pre-commit

##
# Go section
##

.PHONY: deps
deps:
	go mod tidy && go mod vendor && go mod verify

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -cover ./...

.PHONY: tools
tools:
	cd ./tools && ./install.sh
