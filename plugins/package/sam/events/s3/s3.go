package s3

func NewS3Event() *s3Event {
	s3 := new(s3Event)

	s3.Type = "S3"
	s3.Properties.Events = append(s3.Properties.Events, "s3:ObjectCreated:*", "s3:ObjectRemoved:*")

	return s3
}

type s3Event struct {
	Type       string     `json:"Type" yaml:"Type"`
	Properties properties `json:"Properties" yaml:"Properties"`
}

type properties struct {
	Bucket bucket   `json:"Bucket" yaml:"Bucket"`
	Events []string `json:"Events" yaml:"Events"`
}

type bucket struct {
	Ref string `json:"Ref" yaml:"Ref"`
}
