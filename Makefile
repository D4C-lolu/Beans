hello:
	echo "Hello"

build:
	go build -o bin/main cmd/main.go

run:
	go run cmd/main.go

all:
	hello build