package s3adapter

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Driver struct {
	AWSRegion        string
	AWSBucketName    string
	AWSEndpoint      string
	AWSAccessKeyID   string
	AWSSecretKey     string
	WorkingDirectory string
	Username         string
	Password         string
}

func (d *S3Driver) s3service() *s3.Client {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(d.AWSAccessKeyID, d.AWSSecretKey, "")),
	)

	if err != nil {
		return nil
	}

	// Create an Amazon S3 service client
	s3client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = d.AWSRegion
	})

	return s3client
}

func pathToS3PathPrefix(path string) *string {
	path = strings.TrimPrefix(path, "/")

	if path == "" || strings.HasSuffix(path, "/") {
		return aws.String(path)
	}

	p := string(path) + "/"
	return aws.String(p)
}

// func (d *S3Driver) s3DirContents(path string, maxKeys int64, marker string) (*s3.ListObjectsOutput, error) {
// 	svc := d.s3service()

// 	prefix := pathToS3PathPrefix(path)

// 	params := &s3.ListObjectsInput{
// 		Bucket:    aws.String(d.AWSBucketName), // Required
// 		Delimiter: aws.String(d.WorkingDirectory),
// 		// EncodingType: aws.String("EncodingType"),
// 		// Marker:       aws.String("Marker"),
// 		MaxKeys: aws.Int64(maxKeys),
// 		Prefix:  prefix,
// 	}

// 	if marker != "" {
// 		params.Marker = aws.String(marker)
// 	}

// 	resp, err := svc.ListObjects(params)

// 	if err != nil {
// 		// A service error occurred.
// 		fmt.Println("Error: ", err)
// 	} else if err != nil {
// 		// A non-service error occurred.
// 		panic(err)
// 	}

// 	return resp, err
// }

// Authenticate checks that the FTP username and password match
func (d *S3Driver) Authenticate(username string, password string) bool {
	return username == d.Username && password == d.Password
}

// Bytes returns the ContentLength for the path if the key exists
func (d *S3Driver) Bytes(path string) int64 {
	// svc := d.s3service()

	// path = strings.TrimPrefix(path, "/")

	// params := &s3.HeadObjectInput{
	// 	Bucket: aws.String(d.AWSBucketName), // Required
	// 	Key:    aws.String(path),            // Required
	// }
	// resp, err := svc.HeadObject(params)

	// if err != nil {
	// 	// A service error occurred.
	// 	fmt.Println("Error: ", err)
	// 	return -1
	// }

	// return *resp.ContentLength

	return -1
}

// ModifiedTime returns the LastModifiedTime for the path if the key exists
func (d *S3Driver) ModifiedTime(path string) (time.Time, bool) {
	// svc := d.s3service()

	// path = strings.TrimPrefix(path, "/")

	// params := &s3.HeadObjectInput{
	// 	Bucket: aws.String(d.AWSBucketName), // Required
	// 	Key:    aws.String(path),            // Required
	// }
	// resp, err := svc.HeadObject(params)

	// if err != nil {
	// 	// A service error occurred.
	// 	fmt.Println("Error: ", err)
	// 	return time.Now(), false
	// }

	// return *resp.LastModified, true

	return time.Now(), false
}

// ChangeDir “changes directories” on S3 if there are files under the given path
func (d *S3Driver) ChangeDir(path string) bool {
	// resp, err := d.s3DirContents(path, 1, "")

	if strings.HasPrefix(path, "/") {
		d.WorkingDirectory = strings.TrimPrefix(path, "/")
	} else {
		if strings.HasSuffix(d.WorkingDirectory, "/") {
			d.WorkingDirectory = d.WorkingDirectory + path
		} else {
			d.WorkingDirectory = d.WorkingDirectory + "/" + path
		}
	}

	fmt.Println("PWD:", d.WorkingDirectory)
	return true

	//
	// if err == nil && len(resp.Contents) > 0 {
	// 	d.WorkingDirectory = strings.TrimPrefix(path, "/")
	// 	return true
	// } else {
	// 	return false
	// }
}

// DirContents lists “directory” contents on S3
func (d *S3Driver) DirContents(path string) ([]os.FileInfo, bool) {
	// moreObjects := true
	// var objects []*s3.Object

	// var resp *s3.ListObjectsOutput
	// var err error
	// marker := ""

	// for moreObjects {
	// 	resp, err = d.s3DirContents(path, 1000, marker)

	// 	if err == nil {
	// 		for _, obj := range resp.Contents {
	// 			objects = append(objects, obj)
	// 		}

	// 		moreObjects = *resp.IsTruncated

	// 		if moreObjects {
	// 			last := objects[len(objects)-1]
	// 			marker = *last.Key
	// 		}
	// 	}
	// }

	// prefix := pathToS3PathPrefix(path)
	// var files []os.FileInfo
	// var dirs []string

	// for _, object := range objects {
	// 	p := *object.Key

	// 	p = strings.TrimPrefix(p, *prefix)
	// 	var fi os.FileInfo

	// 	if strings.Contains(p, "/") || p == "" {

	// 		parts := strings.Split(p, "/")
	// 		dirPart := parts[0]

	// 		if dirPart != d.WorkingDirectory && dirPart != "" && dirPart != "/" && !stringInSlice(dirPart, dirs) {
	// 			fi = graval.NewDirItem(dirPart)
	// 			files = append(files, fi)

	// 			dirs = append(dirs, dirPart)
	// 		}
	// 	} else {
	// 		fi = graval.NewFileItem(p, *object.Size, *object.LastModified)
	// 		files = append(files, fi)
	// 	}
	// }

	// return files, true

	return []os.FileInfo{}, true
}

// DeleteDir would delete a directory, but isn't currently implemented
func (d *S3Driver) DeleteDir(path string) bool {
	return false
}

// DeleteFile deletes the files from the given path
func (d *S3Driver) DeleteFile(path string) bool {
	// svc := d.s3service()
	// path = strings.TrimPrefix(path, "/")

	// params := &s3.DeleteObjectInput{
	// 	Bucket: aws.String(d.AWSBucketName), // Required
	// 	Key:    aws.String(path),            // Required
	// }
	// _, err := svc.DeleteObject(params)

	// if err != nil {
	// 	// A service error occurred.
	// 	fmt.Println("Error: ", err)
	// 	return false
	// }

	// return true

	return false
}

// Rename isn't supported directly on S3
func (d *S3Driver) Rename(fromPath string, toPath string) bool {
	return false
}

// MakeDir would normally make a directory but this isn't supported on S3
func (d *S3Driver) MakeDir(path string) bool {
	d.PutFile(path+"/", nil)

	return false
}

// GetFile returns a reader for the given path on S3
func (d *S3Driver) GetFile(path string, position int64) (io.ReadCloser, bool) {
	// svc := d.s3service()

	// path = strings.TrimPrefix(path, "/")

	// params := &s3.GetObjectInput{
	// 	Bucket: aws.String(d.AWSBucketName), // Required
	// 	Key:    aws.String(path),            // Required
	// }
	// resp, err := svc.GetObject(params)
	// if err != nil {
	// 	// A service error occurred.
	// 	fmt.Println("Error: ", err)
	// 	return nil, false
	// }

	// return resp.Body, true

	return nil, false
}

// PutFile uploads a file to S3
func (d *S3Driver) PutFile(path string, reader io.Reader) bool {
	svc := d.s3service()

	fmt.Println("put path: ", path)
	fmt.Println("wd: ", d.WorkingDirectory)

	if strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	} else {
		path = d.WorkingDirectory + path
	}

	fileExt := filepath.Ext(path)

	contentType := mime.TypeByExtension(fileExt)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	var body io.ReadSeeker
	if reader != nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(reader)
		body = bytes.NewReader(buf.Bytes())
	}

	uploader := manager.NewUploader(svc)

	params := &s3.PutObjectInput{
		Bucket:      aws.String(d.AWSBucketName), // Required
		Key:         aws.String(path),            // Required
		Body:        body,
		ContentType: aws.String(contentType),
	}

	resp, err := uploader.Upload(context.TODO(), params)

	// resp, err := svc.PutObject(params)
	if err != nil {
		// A service error occurred.
		fmt.Println("Error: ", err)
		return false
	}

	// Pretty-print the response data.
	fmt.Println(resp)

	return true
}

// func stringInSlice(a string, list []string) bool {
// 	for _, b := range list {
// 		if b == a {
// 			return true
// 		}
// 	}
// 	return false
// }
