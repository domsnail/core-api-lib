# Core API Library

protoc -I . -I api/proto -I C:\Users\ydlysak\protoc-32.0\include --go_out=gen/go --go_opt=paths=source_relative --go-grpc_out=$(GO_OUT) --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false api/proto/**/*.proto
