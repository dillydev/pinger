syntax = "proto3";
option go_package = "protobuf";

import "google/api/annotations.proto";
import "protoc-gen-swagger/options/annotations.proto";
import "github.com/gogo/protobuf/gogoproto/gogo.proto";

// Enable registration with golang/protobuf for the grpc-gateway.
option (gogoproto.goproto_registration) = true;

service Pinger {
    rpc Ping(PingRequest) returns (PingResponse) {
        option (google.api.http) = {
            get: "/v1/ping"
        };
    }
}

message PingRequest {}

message PingResponse {
    string message = 1;
}

