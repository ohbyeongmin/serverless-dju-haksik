package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// type S3PutObjectMock struct{}

// func (S3PutObjectMock) PutObject(ctx context.Context,
// 	params *s3.PutObjectInput,
// 	optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {

// 	output := &s3.PutObjectOutput{
// 		VersionId: aws.String("1.0"),
// 	}

// 	return output, nil
// }

// func TestPutObject(t *testing.T) {
// 	b := "testBucket"
// 	k := "file"
// 	input := s3.PutObjectInput{
// 		Bucket: &b,
// 		Key:    &k,
// 	}

// 	api := &S3PutObjectMock{}

// 	resp, err := PutFile(context.Background(), *api, &input)
// 	assert.Nilf(t, err, "error: %s", err)

// 	t.Log("Uploaded version " + *resp.VersionId + " of " + k + " to bucket " + b)
// }

func TestUrl(t *testing.T) {
	assert := assert.New(t)
	d := DaejeonHRCUrl{
		ListURL:     "https://www.dju.ac.kr/hrc/na/ntt/selectNttList.do",
		InfoURL:     "https://www.dju.ac.kr/hrc/na/ntt/selectNttInfo.do",
		DownloadURL: "https://www.dju.ac.kr/common/nttFileDownload.do",
		BbsId:       "2126",
		Mi:          "4597",
	}

	d.SetUrl()
	t.Run("Test getDietListURL()", func(t *testing.T) {
		res, _ := http.Get(d.getDietListURL())
		assert.Equal(res.StatusCode, http.StatusOK)
	})
	t.Run("Test getDietInfoURL()", func(t *testing.T) {
		res, _ := http.Get(d.getDietInfoURL())
		assert.Equal(res.StatusCode, http.StatusOK)
	})
	t.Run("Test getDownloadUrl()", func(t *testing.T) {
		res, _ := http.Get(d.getDownloadURL())
		assert.Equal(res.StatusCode, http.StatusOK)
	})
}
