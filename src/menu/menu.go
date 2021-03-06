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
	"github.com/xuri/excelize/v2"
)

const (
	minRowNum      int    = 4
	maxRowNum      int    = 23
	minColNum      int    = 2
	maxColNum      int    = 7
	filename       string = "diet.xlsx"
	menuObjectName string = "menuObject"
)

type LunOrDin int

const (
	LUNCH LunOrDin = iota
	DINNER
)

type menu map[time.Weekday][]string

type menutable struct {
	table map[LunOrDin]menu
}

type IConvertFilepath interface {
	convertFilepath(file string) string
}

type AWSLambdaFilepath struct{}

func (AWSLambdaFilepath) convertFilepath(file string) string {
	return fmt.Sprintf("/tmp/%s", file)
}

type AWSS3Filepath struct{}

func (AWSS3Filepath) convertFilepath(file string) string {
	return fmt.Sprintf("data/%s", file)
}

var mt *menutable

func HandleErr(err error) {
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
}

func InitMenu() *menutable {
	m := &menutable{
		table: make(map[LunOrDin]menu),
	}
	m.table[LUNCH] = make(menu)
	m.table[DINNER] = make(menu)
	return m
}

func (m *menutable) parseMenuFile(file string, target IConvertFilepath) {
	f, err := excelize.OpenFile(target.convertFilepath(file))
	HandleErr(err)
	sheetName := f.GetSheetList()[0]

	rows, err := f.GetRows(sheetName)

	if err != nil {
		fmt.Println(err)
		return
	}

	lunchOrDinner := LUNCH
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
					lunchOrDinner = DINNER
				}
				continue
			}
			m.table[lunchOrDinner][time.Weekday(j-1)] = append(m.table[lunchOrDinner][time.Weekday(j-1)], colCell)
		}
	}
}

func WriteFile(file string, target IConvertFilepath) {
	f, err := os.Create(target.convertFilepath(file))
	HandleErr(err)
	defer f.Close()
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(mt.table)
	f.Write(buffer.Bytes())
}

func DownloadDietFile(file string, target, keyPath IConvertFilepath) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	HandleErr(err)

	f, err := os.Create(target.convertFilepath(file))
	HandleErr(err)
	defer f.Close()

	client := s3.NewFromConfig(cfg)

	bucket := os.Getenv("bucket")
	key := keyPath.convertFilepath(file)

	downloader := manager.NewDownloader(client)
	_, err = downloader.Download(context.TODO(), f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
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

func UploadFileToS3(file string, target, keyPath IConvertFilepath) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	HandleErr(err)
	client := s3.NewFromConfig(cfg)

	f, err := os.Open(target.convertFilepath(file))
	HandleErr(err)
	defer f.Close()

	bucket := os.Getenv("bucket")
	key := keyPath.convertFilepath(file)
	input := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   f,
	}

	_, err = PutFile(context.TODO(), client, input)
	HandleErr(err)
}

type Test struct {
	Test string `yaml:"test" json:"test"`
}

var testMessage Test

func LambdaHandler() (Test, error) {
	DownloadDietFile(filename, AWSLambdaFilepath{}, AWSS3Filepath{})
	mt = InitMenu()
	mt.parseMenuFile(filename, AWSLambdaFilepath{})
	WriteFile(menuObjectName, AWSLambdaFilepath{})
	UploadFileToS3(menuObjectName, AWSLambdaFilepath{}, AWSS3Filepath{})
	testMessage.Test = "menu lambda test"
	return testMessage, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
