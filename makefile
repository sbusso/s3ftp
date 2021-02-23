-include .env
GO ?= go
DIST=dist
PLATFORM=linux
ARCH=arm64
GOFLAGS :=
LDFLAGS := "-X main.BUCKET_NAME=$(BUCKET_NAME) \
						-X main.ACCESS_KEY_ID=$(ACCESS_KEY_ID) \
						-X main.SECRET_ACCESS_KEY=$(SECRET_ACCESS_KEY) \
						-X main.REGION=$(REGION) \
						-X main.USERNAME=$(USERNAME) \
						-X main.PASSWORD=$(PASSWORD) \
						-X main.HOST=$(HOST) \
						-X main.PORT=$(PORT)"
SSH_SERVER=$(DEPLOY_SERVER)
SSH_USER=$(DEPLOY_USER)
SSH_KEY=$(SSH_KEY_PATH)
BIN=$(DIST)/$(PLATFORM)/$(ARCH)
SSH=ssh -i $(SSH_KEY) $(SSH_USER)@$(SSH_SERVER)

.PHONY: all

all:
	@echo " make <cmd>"
	@echo ""
	@echo "commands:"
	@echo " build          - runs go build"
	@echo ""

build::
	@echo "≫ Building..."
	@CGO_ENABLED=0 GOOS=$(PLATFORM) GOARCH=$(ARCH) $(GO) build -ldflags $(LDFLAGS) -o $(BIN)/$(SERVICE_NAME) *.go

# clean::
# 	@rm -f $(exec)

# test::
#   go test ./...

deploy:: build
	@echo "≫ Deploying..."
	@tar czf - $(BIN)/$(SERVICE_NAME) | $(SSH) 'echo ≫ Backing up old executable...\
				&& test -f /srv/$(SERVICE_NAME)/$(SERVICE_NAME)\
				&& mv /srv/$(SERVICE_NAME)/$(SERVICE_NAME){,.old}\
				&& echo ≫ Extracting into /srv/$(SERVICE_NAME)/...\
				;  tar xzf - -C /srv/$(SERVICE_NAME)/\
				&& echo ≫ Restarting service...\
				&& sudo systemctl daemon-reload\
				&& sudo service $(SERVICE_NAME) restart\
				&& echo ≫ Checking status...\
				&& sudo service $(SERVICE_NAME) status\
				&& echo ≫ Done'
setup::
		echo "Running command on $(SSH_SERVER): $$SETUP_CMD"
		echo "$$SERVICE_FILE" | $(SSH) "$$SETUP_CMD"

define SETUP_CMD
 echo Making directory \
 && sudo mkdir -p /srv/$(SERVICE_NAME)/ \
 && echo Touching service file \
 && sudo touch /srv/$(SERVICE_NAME)/$(SERVICE_NAME).service \
 && echo Setting folder permissions \
 && sudo chown $$USER -R /srv/$(SERVICE_NAME)/ \
 && echo Creating service file \
 && sudo cat > /srv/$(SERVICE_NAME)/$(SERVICE_NAME).service \
 && echo Enabling service \
 && sudo sudo systemctl enable /srv/$(SERVICE_NAME)/$(SERVICE_NAME).service \
 && sudo systemctl daemon-reload
endef
export SETUP_CMD

define SERVICE_FILE
[Unit]
Description=$(SERVICE_NAME)

[Service]
Type=simple
WorkingDirectory=/srv/$(SERVICE_NAME)/
ExecStart=/srv/$(SERVICE_NAME)/$(BIN)/$(SERVICE_NAME)
Restart=always
RestartSec=90
StartLimitInterval=400
StartLimitBurst=3

[Install]
WantedBy=multi-user.target

endef
export SERVICE_FILE