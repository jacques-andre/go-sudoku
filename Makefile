build:
	go build -o bin/game game.go

run:
	go run game.go

test:
	go test -v ./...
