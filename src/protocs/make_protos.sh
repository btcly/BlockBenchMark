#!/bin/bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc -I=./ --go_out=./ --plugin=`which protoc-gen-go` --plugin=`go env |grep GOPATH|awk -F "=" '{print$2}'|awk -F "'" '{print $2}'`/bin/protoc-gen-go-grpc --go-grpc_out=./ *.proto
# protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./*.proto

