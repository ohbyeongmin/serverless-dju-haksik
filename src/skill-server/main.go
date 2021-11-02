package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func testLogs() {
	log.Println("test logs")
}

func main() {
	testLogs()
	lambda.Start(HandleSkillServer)
}
