#!/bin/bash
protoc -I/usr/local/include -Iapi/protobuf \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
  --gogoslick_out=plugins=grpc:api/protobuf \
  --swagger_out=logtostderr=true,allow_delete_body=true:api/openapi/ \
  --grpc-gateway_out=logtostderr=true,allow_delete_body=true:api/protobuf \
  api/protobuf/*.pb
  