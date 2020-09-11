.PHONY: build
build:
	go build -o ./bin/vel ./cmd/vel/main.go

.PHONY: run
run:
	go run cmd/vel/main.go