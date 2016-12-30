package sam

import (
	"fmt"
	"strings"

	"github.com/go-yaml/yaml"

	"github.com/pmuens/simple/plugins/package/sam/bucket"
	"github.com/pmuens/simple/plugins/package/sam/function"
	"github.com/pmuens/simple/util"
)

var serviceBucketName = util.GetServiceBucketName()
var zipName = util.GetZipName()

func NewSAM(functions []string) *sam {
	s := new(sam)
	s.AWSTemplateFormatVersion = "2010-09-09"
	s.Transform = "AWS::Serverless-2016-10-31"
	s.Resources = make(map[string]interface{})

	s.functions = functions

	return s
}

func (s *sam) GetCreateStackYAML() []byte {
	s.addServiceBucket()

	res, err := yaml.Marshal(s)
	if err != nil {
		util.LogPanic(err)
	}

	return res
}

func (s *sam) GetUpdateStackYAML() []byte {
	s.addServiceBucket()
	s.addFunctions()

	res, err := yaml.Marshal(s)
	if err != nil {
		util.LogPanic(err)
	}

	return res
}

func (s *sam) addFunctions() {
	for _, f := range s.functions {
		fun := function.NewFunction(f)

		// special case because the event needs an own S3 bucket
		event := strings.Split(f, "-")[0]
		if event == "s3" {
			bucketResourceLogicalId := fmt.Sprintf("%s %s", f, "Bucket")
			bucketName := fmt.Sprintf("%s-%s", f, "bucket")
			s.Resources[util.TitleizeAndConcat(bucketResourceLogicalId)] = bucket.NewBucket(bucketName)
		}

		codeUri := fmt.Sprintf("s3://%s/%s", serviceBucketName, zipName)
		fun.Properties.CodeUri = codeUri

		s.Resources[util.TitleizeAndConcat(f)] = fun
	}
}

func (s *sam) addServiceBucket() {
	s.Resources["SimpleServiceBucket"] = bucket.NewBucket(serviceBucketName)
}

type sam struct {
	AWSTemplateFormatVersion string                 `json:"AWSTemplateFormatVersion" yaml:"AWSTemplateFormatVersion"`
	Transform                string                 `json:"Transform" yaml:"Transform"`
	Resources                map[string]interface{} `json:"Resources" yaml:"Resources"`

	functions []string `json:"-" yaml:"-"` // TODO try to get rid of this...
}
