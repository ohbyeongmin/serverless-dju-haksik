.PHONY: build clean test lambda

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/main menu.go

test:
	go test -v -coverprofile cover.out ./... && go tool cover -html=cover.out

lambda:
	aws lambda invoke --region=ap-northeast-2 --function-name=menu response.json
