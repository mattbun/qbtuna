.PHONY: all
all: tidy fmt build

.PHONY: build
build:
	go build main.go

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: fmt
fmt:
	go fmt

.PHONY: run
run:
	go run main.go

.PHONY: docker-image
docker-image:
	docker build -t qbtuna:latest .
