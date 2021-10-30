.PHONY: build clean

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/main menu.go

test:
	aws lambda invoke --region=ap-northeast-2 --function-name=menu response.json