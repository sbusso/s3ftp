package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sbusso/s3ftp/s3adapter"
	"github.com/yob/graval"
)

var (
	BUCKET_NAME       string
	ACCESS_KEY_ID     string
	SECRET_ACCESS_KEY string
	REGION            string
	USERNAME          string
	PASSWORD          string
	HOST              string
	PORT              string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	if REGION == "" {
		REGION = os.Getenv("REGION")
	}
	if BUCKET_NAME == "" {
		BUCKET_NAME = os.Getenv("BUCKET_NAME")
	}
	if ACCESS_KEY_ID == "" {
		ACCESS_KEY_ID = os.Getenv("ACCESS_KEY_ID")
	}
	if SECRET_ACCESS_KEY == "" {
		SECRET_ACCESS_KEY = os.Getenv("SECRET_ACCESS_KEY")
	}
	if USERNAME == "" {
		USERNAME = os.Getenv("USERNAME")
	}
	if PASSWORD == "" {
		PASSWORD = os.Getenv("PASSWORD")
	}
	if HOST == "" {
		HOST = os.Getenv("HOST")
	}
	if PORT == "" {
		PORT = os.Getenv("PORT")
	}
}

func main() {

	factory := &s3adapter.S3DriverFactory{
		AWSRegion:      REGION,
		AWSBucketName:  BUCKET_NAME,
		AWSAccessKeyID: ACCESS_KEY_ID,
		AWSSecretKey:   SECRET_ACCESS_KEY,
		Username:       USERNAME,
		Password:       PASSWORD,
	}

	p, _ := strconv.Atoi(PORT)

	server := graval.NewFTPServer(&graval.FTPServerOpts{
		ServerName:  "s3ftp",
		Factory:     factory,
		Hostname:    HOST,
		Port:        p,
		PasvMinPort: 60200,
		PasvMaxPort: 60300,
		// PassiveOpts: &graval.PassiveOpts{
		// 	ListenAddress: config.Host,
		// 	NatAddress:    config.Host,
		// 	PassivePorts: &graval.PassivePorts{
		// 		Low:  42000,
		// 		High: 45000,
		// 	},
		// },
	})

	log.Printf("S3FTP server listening on %s:%d", HOST, p)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
