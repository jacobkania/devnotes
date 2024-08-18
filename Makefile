#####################
#### Primary Commands
#####################

init: fetch-dependencies-go

run: build-go-local run-go-local

build: build-go-local

test: test-go

migrate: create-migration

#####################
#### Application Commands
#####################

fetch-dependencies-go:
	go get ./...

test-go:
	go test -race ./...

# Local
run-go-local:
	bin/dn

build-go-production:
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o bin/dn .

build-go-local:
	go build -o bin/dn ./cmd/term

run-go-built: build-go-local
	./bin/dn

# Migration Commands

create-migration:
	@read -p "Enter migration name: " name; \
	timestamp=$$(date +%s); \
	mkdir -p migrate; \
	touch migrate/$${timestamp}_$${name}.up.sql; \
	touch migrate/$${timestamp}_$${name}.down.sql; \
	echo "Migration files created: migrate/$${timestamp}_$${name}.up.sql, migrate/$${timestamp}_$${name}.down.sql"

