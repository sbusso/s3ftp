include .env
GO ?= go
DIST=dist
PLATFORM=linux
ARCH=amd64
GOFLAGS :=
LDFLAGS := "-X main.BUCKET_NAME=$(BUCKET_NAME) \
						-X main.ACCESS_KEY_ID=$(ACCESS_KEY_ID) \
						-X main.SECRET_ACCESS_KEY=$(SECRET_ACCESS_KEY) \
						-X main.REGION=$(REGION) \
						-X main.USERNAME=$(USERNAME) \
						-X main.PASSWORD=$(PASSWORD) \
						-X main.HOST=$(HOST) \
						-X main.PORT=$(PORT)"

.PHONY: all

all:
	@echo " make <cmd>"
	@echo ""
	@echo "commands:"
	@echo " build          - runs go build"
	@echo ""

build:
	@echo "Building..."
	@CGO_ENABLED=0 GOOS=$(PLATFORM) GOARCH=$(ARCH) $(GO) build -ldflags $(LDFLAGS) -o $(DIST)/$(PLATFORM)/$(ARCH)/s3ftp *.go

clean:
	@rm -f $(exec)
