.PHONY: build clean test lambda

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/main crawling.go
	cp -r config bin/

test:
	go test ./... -cover

lambda:
	aws lambda invoke --region=ap-northeast-2 --function-name=crawling response.json