package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	menuObjectName string = "menuObject"
)

type menu map[time.Weekday][]string

type menutable struct {
	table map[LunOrDin]menu
}

func DownloadObjectFile() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	HandleErr(err)

	file, err := os.Create(fmt.Sprintf("/tmp/%s", menuObjectName))
	HandleErr(err)
	defer file.Close()

	client := s3.NewFromConfig(cfg)

	bucket := os.Getenv("bucket")
	key := fmt.Sprintf("data/%s", menuObjectName)
	downloader := manager.NewDownloader(client)
	_, err = downloader.Download(context.TODO(), file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	HandleErr(err)
}

func ReadFile() *menutable {
	f, _ := os.Open(fmt.Sprintf("/tmp/%s", menuObjectName))
	defer f.Close()
	var m menutable = menutable{}
	var buffer bytes.Buffer
	dec := gob.NewDecoder(&buffer)
	buffer.ReadFrom(f)
	dec.Decode(&m.table)
	return &m
}

func HandleSkillServer(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var weekDay time.Weekday
	params := BotRequestType{
		Action: ActionType{
			Params: make(map[string]string),
		},
	}
	err := json.Unmarshal([]byte(req.Body), &params)
	HandleErr(err)

	dateTag, day := ParseSysdate(params.Action.Params["sys_date"])
	lunchOrDinner, err := StringToLunOrDin(params.Action.Params["lunch_dinner"])
	HandleErr(err)

	t := time.Now()
	loc, _ := time.LoadLocation("Asia/Seoul")
	kst := t.In(loc)
	if day != 0 {
		weekDay = DayToWeekday(day)
	} else if dateTag == "today" {
		weekDay = kst.Weekday()
	} else if dateTag == "tomorrow" {
		weekDay = kst.Add(time.Hour * 24).Weekday()
	} else if dateTag == "Monday" {
		weekDay = time.Monday
	} else if dateTag == "Tuesday" {
		weekDay = time.Tuesday
	} else if dateTag == "Wednesday" {
		weekDay = time.Wednesday
	} else if dateTag == "Thursday" {
		weekDay = time.Thursday
	} else if dateTag == "Friday" {
		weekDay = time.Friday
	}

	DownloadObjectFile()
	m := ReadFile()

	bodyResponse := GetOneMenuReasponse(m.table[lunchOrDinner][weekDay], lunchOrDinner)
	response, err := json.Marshal(bodyResponse)
	HandleErr(err)

	res := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(response),
	}

	return res, nil

}
