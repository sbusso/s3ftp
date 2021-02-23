package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sbusso/s3ftp/s3adapter"
	"github.com/yob/graval"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config, _ := ParseConfig()

	factory := &s3adapter.S3DriverFactory{
		AWSRegion:      config.AWSRegion,
		AWSBucketName:  config.AWSBucketName,
		AWSAccessKeyID: config.AWSAccessKeyID,
		AWSSecretKey:   config.AWSSecretKey,
		Username:       config.Username,
		Password:       config.Password,
	}

	server := graval.NewFTPServer(&graval.FTPServerOpts{
		ServerName: "s3ftp",
		Factory:    factory,
		Hostname:   config.Host,
		Port:       config.Port,
		// PassiveOpts: &graval.PassiveOpts{
		// 	ListenAddress: config.Host,
		// 	NatAddress:    config.Host,
		// 	PassivePorts: &graval.PassivePorts{
		// 		Low:  42000,
		// 		High: 45000,
		// 	},
		// },
	})

	log.Printf("S3FTP server listening on %s:%d", config.Host, config.Port)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
