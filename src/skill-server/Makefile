.PHONY: build clean test lambda

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/main main.go request-from-bot.go  response.go skill_server.go utils.go service.go

test:
	go test -v -coverprofile cover.out ./...

lambda:
	aws lambda invoke --region=ap-northeast-2 --function-name=skillserver response.json
