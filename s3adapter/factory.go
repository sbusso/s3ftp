package s3adapter

import (
	"github.com/koofr/graval"
)

type S3DriverFactory struct {
	AWSRegion      string
	AWSBucketName  string
	AWSAccessKeyID string
	AWSSecretKey   string
	Username       string
	Password       string
}

func (f *S3DriverFactory) NewDriver() (d graval.FTPDriver, err error) {
	return &S3Driver{
		AWSRegion:      f.AWSRegion,
		AWSBucketName:  f.AWSBucketName,
		AWSAccessKeyID: f.AWSAccessKeyID,
		AWSSecretKey:   f.AWSSecretKey,
		Username:       f.Username,
		Password:       f.Password,
	}, nil
}
