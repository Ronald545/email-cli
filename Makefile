dev: 
	reflex -r '\.c' -s -- go run main.go

build:
	go build -o email-cli main.go

install:
	mv email-cli /usr/bin/

remove:
	rm /usr/bin/email-cli
