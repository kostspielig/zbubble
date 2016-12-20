TARGET = zbubble

.PHONY: all exec run fmt test

all: build exec

exec:
	@./$(TARGET)

run:
	go run main.go

test:
	go test -timeout=5s github.com/kostspielig/zbubble/...

build:
	go build \
		-o $(TARGET)

fmt:
	@go fmt ./...

buildweb:
	gopherjs build -o web/zbubble.js
