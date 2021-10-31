package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
)

type DaejeonHRCUrl struct {
	ListURL     string `yaml:"listURL"`
	InfoURL     string `yaml:"infoURL"`
	DownloadURL string `yaml:"downloadURL"`
	BbsId       string `yaml:"bbsId"`
	Mi          string `yaml:"mi"`
	NttSn       string `yaml:"nttSn"`
	FileKey     string `yaml:"fileKey"`
}

var filename string = "diet.xlsx"

func HandleErr(err error) {
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
}

func GetConfiguration(cFileName, cFileType, cFilePath string) *DaejeonHRCUrl {
	d := &DaejeonHRCUrl{}
	urlViper := viper.New()
	urlViper.SetConfigName(cFileName)
	urlViper.SetConfigType(cFileType)
	urlViper.AddConfigPath(cFilePath)
	err := urlViper.ReadInConfig()
	if err != nil {
		fmt.Printf("fatal error config file: url, err: %s \n", err)
		os.Exit(1)
	}

	err = urlViper.Unmarshal(&d)
	if err != nil {
		fmt.Printf("fatal error config file: unmarshal url config, err: %s \n", err)
		os.Exit(1)
	}
	return d
}

func (d *DaejeonHRCUrl) SetUrl() {
	d.setNttSn()
	d.setFileKey()
}

func (d *DaejeonHRCUrl) setNttSn() {
	url := d.getDietListURL()

	res, err := http.Get(url)
	HandleErr(err)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	HandleErr(err)

	defer res.Body.Close()

	doc.Find(".nttInfoBtn").Each(func(i int, s *goquery.Selection) {
		_, ok := s.Attr("href")
		if ok {
			data, _ := s.Attr("data-id")
			d.NttSn = data
		}
	})
}

func (d *DaejeonHRCUrl) setFileKey() {
	var downloadLink string
	url := d.getDietInfoURL()
	res, err := http.Get(url)
	HandleErr(err)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	HandleErr(err)

	defer res.Body.Close()

	doc.Find(".file").Children().Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			downloadLink = href
		}
	})

	d.FileKey = strings.Split(strings.Split(strings.Split(downloadLink, "?")[1], "&")[0], "=")[1]
}

func (d *DaejeonHRCUrl) getDietListURL() string {
	return fmt.Sprintf("%s?mi=%s&bbsId=%s", d.ListURL, d.Mi, d.BbsId)
}

func (d *DaejeonHRCUrl) getDietInfoURL() string {
	return fmt.Sprintf("%s?nttSn=%s&bbsId=%s&mi=%s", d.InfoURL, d.NttSn, d.BbsId, d.Mi)
}

func (d *DaejeonHRCUrl) getDownloadURL() string {
	return fmt.Sprintf("%s?fileKey=%s&nttSn=%s&bbsId=%s", d.DownloadURL, d.FileKey, d.NttSn, d.BbsId)
}

func (d *DaejeonHRCUrl) DownloadDietFile() {
	log.Printf("Download file : %s", filename)
	downloadLink := d.getDownloadURL()
	downRes, err := http.Get(downloadLink)
	HandleErr(err)

	download, err := ioutil.ReadAll(downRes.Body)
	HandleErr(err)

	file, err := os.Create(fmt.Sprintf("/tmp/%s", filename))
	HandleErr(err)
	defer file.Close()

	file.Write(download)
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

	file, err := os.Open(fmt.Sprintf("/tmp/%s", filename))
	HandleErr(err)
	defer file.Close()

	bucket := os.Getenv("bucket")
	key := fmt.Sprintf("data/%s", filename)
	input := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
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
	url := GetConfiguration("url", "yaml", "config")
	url.SetUrl()
	url.DownloadDietFile()
	UploadFileToS3()
	testMessage.Test = "OK"
	return testMessage, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
