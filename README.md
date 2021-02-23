# s3ftp

FTP to S3 upload only.

## Compile

```shell
  CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o dist/arm64/s3ftp *.go

  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/amd64/s3ftp *.go
```

```
  docker run -t -p 21:21 -p 60200-60300:60200-60300 s3ftp
```
