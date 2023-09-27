GOPATH:=$(shell go env GOPATH)
# VERSION=$(shell git describe --tags --always)
APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)
API_PROTO_FILES=$(shell find . -name *.proto)
APP_NAME=$(shell echo $(APP_RELATIVE_PATH) | sed -En "s/\//-/p")


META_PROTO_FILES=$(shell find api/meta -name *.proto)
META_PB_FILES=$(shell find api/meta -name *.pb.go)

USER_PROTO_FILE=$(shell find api/user -name user.proto)
USER_PB_FILE=$(shell find api/user -name user.pb.go)
USER_SVC_FILE=$(shell find api/user -name service.proto)
USER_SVC_PB_FILE=$(shell find api/user -name service.pb.go)

TOKEN_PROTO_FILE=$(shell find api/token -name token.proto)
TOKEN_PB_FILE=$(shell find api/token -name token.pb.go)
TOKEN_SECTEST_PROTO_FILE=$(shell find api/token -name secrets.proto)
TOKEN_SECTEST_PB_FILE=$(shell find api/token -name secrets.pb.go)
TOKEN_SVC_FILE=$(shell find api/token -name service.proto)

CODE_PROTO_FILE=$(shell find api/token -name code.proto)
CODE_PB_FILE=$(shell find api/token -name code.pb.go)

APP_CONFIG_FILE=$(shell find app/configs -name conf.proto)


.PHONY: init
# init env
init:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/google/wire/cmd/wire
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install github.com/favadi/protoc-go-inject-tag@latest

.PHONY: meta
# generate internal proto struct
pb-meta:
	
	protoc --proto_path=. \
           --go_out=paths=source_relative:. \
           $(META_PROTO_FILES)

	protoc-go-inject-tag -input=$(META_PB_FILES)

	go fmt ./...

.PHONY: user
# generate user proto struct
pb-user:
	protoc --proto_path=. \
           --go_out=paths=source_relative:. \
           $(USER_PROTO_FILE)

	protoc-go-inject-tag -input=$(USER_PB_FILE)

	go fmt ./...

.PHONY: svc user grpc/http
# generate user proto struct
svc-user:
	protoc --proto_path=. \
		   --proto_path=./third_party \
		   --go_out=paths=source_relative:. \
           --go-grpc_out=paths=source_relative:. \
		   --grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
           $(USER_SVC_FILE)

	go fmt ./...

.PHONY: token
# generate user proto struct
pb-token:
	protoc --proto_path=. \
           --go_out=paths=source_relative:. \
           $(TOKEN_PROTO_FILE)

	protoc-go-inject-tag -input=$(TOKEN_PB_FILE)

	go fmt ./...

.PHONY: token secrets
# generate user proto struct
pb-tk-secrets:
	protoc --proto_path=. \
           --go_out=paths=source_relative:. \
           $(TOKEN_SECTEST_PROTO_FILE)

	protoc-go-inject-tag -input=$(TOKEN_SECTEST_PB_FILE)

	go fmt ./...

.PHONY: token code
# generate user proto struct
pb-tk-code:
	protoc --proto_path=. \
           --go_out=paths=source_relative:. \
           $(CODE_PROTO_FILE)

	protoc-go-inject-tag -input=$(CODE_PB_FILE)

	go fmt ./...

.PHONY: svc user grpc/http
# generate user proto struct
svc-token:
	protoc --proto_path=. \
		   --proto_path=./third_party \
		   --go_out=paths=source_relative:. \
           --go-grpc_out=paths=source_relative:. \
		   --grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
           $(TOKEN_SVC_FILE)

	go fmt ./...

.PHONY: config
# generate user proto struct
config:
	protoc --proto_path=. \
		   --proto_path=./third_party \
		   --go_out=paths=source_relative:. \
           $(APP_CONFIG_FILE)

	go fmt ./...

.PHONY: wire
# generate wire
wire:
	cd app/cmd && wire

.PHONY: all
# generate user proto struct
pb-all:pb-meta pb-user
