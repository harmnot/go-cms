dev:
	go run ./server/main.go
local:
	./go.sh dev
start:
	./go.sh prod
build:
	go build .