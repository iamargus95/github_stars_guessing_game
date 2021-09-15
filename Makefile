all: build

build: main.go
	go build -o githubGuessStars.o

clean:
	rm githubGuessStars.o
	go clean

run:
	go run .

test:
	go test