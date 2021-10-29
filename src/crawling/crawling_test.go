package main

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
)

type S3PutObjectMock struct{}

func (S3PutObjectMock) PutObject(ctx context.Context,
	params *s3.PutObjectInput,
	optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {

	output := &s3.PutObjectOutput{
		VersionId: aws.String("1.0"),
	}

	return output, nil
}

func TestPutObject(t *testing.T) {
	b := "testBucket"
	k := "file"
	input := s3.PutObjectInput{
		Bucket: &b,
		Key:    &k,
	}

	api := &S3PutObjectMock{}

	resp, err := PutFile(context.Background(), *api, &input)
	assert.Nilf(t, err, "error: %s", err)

	t.Log("Uploaded version " + *resp.VersionId + " of " + k + " to bucket " + b)
}

func TestUrl(t *testing.T) {
	d := DaejeonHRCUrl{
		ListURL:     "https://www.dju.ac.kr/hrc/na/ntt/selectNttList.do",
		InfoURL:     "https://www.dju.ac.kr/hrc/na/ntt/selectNttInfo.do",
		DownloadURL: "https://www.dju.ac.kr/common/nttFileDownload.do",
		BbsId:       "2126",
		Mi:          "4597",
	}

	dietListUrl := "https://www.dju.ac.kr/hrc/na/ntt/selectNttList.do?mi=4597&bbsId=2126"
	dietInfoUrl := "https://www.dju.ac.kr/hrc/na/ntt/selectNttInfo.do?nttSn=512606&bbsId=2126&mi=4597"
	downloadUrl := "https://www.dju.ac.kr/common/nttFileDownload.do?fileKey=9c53541b5a24c3d6442778acbfa71e09&nttSn=512606&bbsId=2126"
	d.SetUrl()

	t.Run("Test getDietListURL()", func(t *testing.T) {
		assert.Equal(t, dietListUrl, d.getDietListURL(), "invalid url")
	})
	t.Run("Test getDietInfoURL()", func(t *testing.T) {
		assert.Equal(t, dietInfoUrl, d.getDietInfoURL(), "invalid url")
	})
	t.Run("Test getDownloadUrl()", func(t *testing.T) {
		assert.Equal(t, downloadUrl, d.getDownloadURL(), "invalid url")
	})
}
