.PHONY: run build test vet tidy docker-build docker-up clean

run:
	go run ./cmd/server

build:
	go build -o todo-api ./cmd/server

test:
	go test -race -count=1 ./...

vet:
	go vet ./...

tidy:
	go mod tidy

docker-build:
	docker build -t todo-api:local .

docker-up:
	docker compose up --build

clean:
	rm -f todo-api
