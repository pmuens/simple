package aws

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/pmuens/simple/util"
)

func PutObject(bucket string, key string, body []byte) *s3.PutObjectOutput {
	s := newS3(region)

	params := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(body),
	}

	resp, err := s.PutObject(params)
	if err != nil {
		util.LogPanic(err)
	}

	return resp
}

func ListObjects(bucket string) *s3.ListObjectsOutput {
	s := newS3(region)

	params := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	}

	resp, err := s.ListObjects(params)
	if err != nil {
		util.LogPanic(err)
	}

	return resp
}

func DeleteObjects(bucket string, objects []*s3.ObjectIdentifier) *s3.DeleteObjectsOutput {
	s := newS3(region)

	params := &s3.DeleteObjectsInput{
		Bucket: aws.String(bucket),
		Delete: &s3.Delete{
			Objects: objects,
		},
	}

	resp, err := s.DeleteObjects(params)
	if err != nil {
		util.LogPanic(err)
	}

	return resp
}

func EmptyBucket(bucket string) {
	resp := ListObjects(bucket)

	objects := resp.Contents
	if len(objects) == 0 {
		return
	}

	objectIdentifiers := []*s3.ObjectIdentifier{}
	for _, s3Object := range objects {
		object := &s3.ObjectIdentifier{
			Key: s3Object.Key,
		}
		objectIdentifiers = append(objectIdentifiers, object)
	}

	DeleteObjects(bucket, objectIdentifiers)
}

func newS3(region string) *s3.S3 {
	sess := CreateSession(region)
	return s3.New(sess)
}
