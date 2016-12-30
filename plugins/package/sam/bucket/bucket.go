package bucket

func NewBucket(name string) *bucket {
	b := new(bucket)

	b.Type = "AWS::S3::Bucket"
	b.Properties.BucketName = name

	return b
}

type bucket struct {
	Type       string     `json:"Type" yaml:"Type"`
	Properties properties `json:"Properties" yaml:"Properties"`
}

type properties struct {
	BucketName string `json:"BucketName" yaml:"BucketName"`
}
