package main

import (
	"os"
	"strconv"
)

const DefaultConfigPath = "./ftp.toml"

func ParseConfig() (*Config, error) {
	var globalConfig Config
	globalConfig.Host = os.Getenv("HOST")
	globalConfig.Port, _ = strconv.Atoi(os.Getenv("PORT"))

	globalConfig.AWSRegion = os.Getenv("REGION")
	globalConfig.AWSBucketName = os.Getenv("BUCKET_NAME")
	globalConfig.AWSAccessKeyID = os.Getenv("ACCESS_KEY_ID")
	globalConfig.AWSSecretKey = os.Getenv("SECRET_ACCESS_KEY")
	globalConfig.Username = os.Getenv("USERNAME")
	globalConfig.Password = os.Getenv("PASSWORD")

	return &globalConfig, nil
}

type Config struct {
	Host string
	Port int

	AWSRegion      string
	AWSBucketName  string
	AWSAccessKeyID string
	AWSSecretKey   string
	Username       string
	Password       string
}
