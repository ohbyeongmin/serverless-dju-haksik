package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/ohbyeongmin/daejeon-haksik/constants"
	"github.com/xuri/excelize/v2"
)

const (
	minRowNum      int    = 4
	maxRowNum      int    = 23
	minColNum      int    = 2
	maxColNum      int    = 7
	Filename       string = "diet.xlsx"
	menuObjectName string = "menuObject"
)

type menu map[time.Weekday][]string

type menutable struct {
	table map[constants.LunOrDin]menu
}

var mt menutable

func HandleErr(err error) {
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
}

func InitMenu() {
	mt = menutable{
		table: make(map[constants.LunOrDin]menu),
	}
	mt.table[constants.LUNCH] = make(menu)
	mt.table[constants.DINNER] = make(menu)

	mt.parseMenuFile()
}

func (m *menutable) parseMenuFile() {
	f, err := excelize.OpenFile("/tmp/diet.xlsx")
	HandleErr(err)
	sheetName := f.GetSheetList()[0]

	rows, err := f.GetRows(sheetName)

	if err != nil {
		fmt.Println(err)
		return
	}

	lunchOrDinner := constants.LUNCH
	for i, row := range rows {
		if i < minRowNum {
			continue
		}
		if i > maxRowNum {
			break
		}
		for j, colCell := range row {
			if j < minColNum || j >= maxColNum {
				if colCell == "석  식" {
					lunchOrDinner = constants.DINNER
				}
				continue
			}
			mt.table[lunchOrDinner][time.Weekday(j-1)] = append(mt.table[lunchOrDinner][time.Weekday(j-1)], colCell)
		}
	}
}

func WriteFile() {
	f, err := os.Create(fmt.Sprintf("/tmp/%s", menuObjectName))
	HandleErr(err)
	defer f.Close()
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(mt.table)
	f.Write(buffer.Bytes())
}

func ReadFile() {
	f, _ := os.Open(menuObjectName)
	defer f.Close()
	var m menutable = menutable{}
	var buffer bytes.Buffer
	dec := gob.NewDecoder(&buffer)
	buffer.ReadFrom(f)
	dec.Decode(&m.table)
	fmt.Println(m.table[constants.LUNCH][time.Friday])
}

func DownloadDietFile() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	HandleErr(err)

	file, err := os.Create("/tmp/diet.xlsx")
	HandleErr(err)
	defer file.Close()

	client := s3.NewFromConfig(cfg)

	bucket := os.Getenv("bucket")

	downloader := manager.NewDownloader(client)
	_, err = downloader.Download(context.TODO(), file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(Filename),
	})
	HandleErr(err)
}

func DownloadObjectFile() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	HandleErr(err)

	file, err := os.Create(menuObjectName)
	HandleErr(err)
	defer file.Close()

	client := s3.NewFromConfig(cfg)

	bucket := os.Getenv("bucket")

	downloader := manager.NewDownloader(client)
	_, err = downloader.Download(context.TODO(), file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(menuObjectName),
	})
	HandleErr(err)
}

type S3PutObjectAPI interface {
	PutObject(ctx context.Context,
		params *s3.PutObjectInput,
		optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

func PutFile(c context.Context, api S3PutObjectAPI, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return api.PutObject(c, input)
}

func UploadFileToS3() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	HandleErr(err)
	client := s3.NewFromConfig(cfg)

	file, err := os.Open(fmt.Sprintf("/tmp/%s", menuObjectName))
	HandleErr(err)
	defer file.Close()

	bucket := os.Getenv("bucket")
	input := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(menuObjectName),
		Body:   file,
	}

	_, err = PutFile(context.TODO(), client, input)
	HandleErr(err)
}

type Test struct {
	Test string `yaml:"test" json:"test"`
}

var testMessage Test

func LambdaHandler() (Test, error) {
	DownloadDietFile()
	InitMenu()
	WriteFile()
	UploadFileToS3()
	testMessage.Test = "menu lambda test"
	return testMessage, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
