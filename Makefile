.PHONY: run build docker-build docker-run

run:
	go run ./cmd/gatewayx

mock:
	go run ./cmd/mock

build:
	go build -o bin/gatewayx ./cmd/gatewayx

docker-build:
	docker build -t gatewayx .

docker-run:
	docker run -p 8080:8080 gatewayx