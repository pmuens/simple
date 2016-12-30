package function

import (
	"fmt"
	"strings"

	"github.com/pmuens/simple/plugins/package/sam/events/alexa"
	"github.com/pmuens/simple/plugins/package/sam/events/api"
	"github.com/pmuens/simple/plugins/package/sam/events/s3"
	"github.com/pmuens/simple/plugins/package/sam/events/timer"
	"github.com/pmuens/simple/util"
)

func NewFunction(name string) *function {
	event := strings.Split(name, "-")[0]

	f := new(function)

	p := []string{"AmazonDynamoDBFullAccess", "AmazonS3FullAccess"}

	f.Type = "AWS::Serverless::Function"
	f.Properties.Runtime = "nodejs4.3" // TODO auto detect
	f.Properties.Handler = fmt.Sprintf("%s/handler.handler", name)
	f.Properties.CodeUri = "s3://<some-bucket>/<some-file>.zip"
	f.Properties.Policies = p
	f.Properties.Events = make(map[string]interface{})

	f.name = name

	switch event {
	case "s3":
		f.addS3Event()
	case "api":
		f.addAPIEvents()
	case "timer":
		f.addTimerEvent()
	case "alexa-skill":
		f.addAlexaSkillEvent()
	}

	return f
}

func (f *function) addS3Event() {
	s3 := s3.NewS3Event()

	bucketName := fmt.Sprintf("%s %s", f.name, "Bucket")
	s3.Properties.Bucket.Ref = util.TitleizeAndConcat(bucketName)

	f.Properties.Events["S3"] = s3
}

func (f *function) addAPIEvents() {
	methods := []string{"post", "put", "get", "patch", "delete"}

	for _, method := range methods {
		api := api.NewAPIEvent()

		api.Properties.Path = f.name
		api.Properties.Method = method

		name := fmt.Sprintf("%s %s", "Api", method)

		f.Properties.Events[util.TitleizeAndConcat(name)] = api
	}
}

func (f *function) addTimerEvent() {
	timer := timer.NewTimerEvent()

	f.Properties.Events["Timer"] = timer
}

func (f *function) addAlexaSkillEvent() {
	alexaSkill := alexa.NewAlexaSkillEvent()

	f.Properties.Events["AlexaSkillEvent"] = alexaSkill
}

type function struct {
	Type       string     `json:"Type" yaml:"Type"`
	Properties properties `json:"Propeties" yaml:"Properties"`

	name string `json:"-" yaml:"-"` // TODO try to get rid of this...
}

type properties struct {
	Handler  string                 `json:"Handler" yaml:"Handler"`
	Runtime  string                 `json:"Runtime" yaml:"Runtime"`
	CodeUri  string                 `json:"CodeUri" yaml:"CodeUri"`
	Policies []string               `json:"Policies" yaml:"Policies"`
	Events   map[string]interface{} `json:"Events" yaml:"Events"`
}
