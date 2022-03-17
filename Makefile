dev: 
	reflex -r '\.c' -s -- go run main.go

build:
	go build -o email-cli main.go

install:
	make build
	mv email-cli /usr/bin/

remove:
	rm /usr/bin/email-cli

lint:
	gofmt -w *.go 
	gofmt -w */*.go
