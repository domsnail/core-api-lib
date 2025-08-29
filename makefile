PROTO_DIR=api/proto
GO_OUT=gen/go
TS_OUT=gen/ts

# autodetect всех proto-файлов
PROTO_FILES=$(shell powershell -Command "Get-ChildItem -Recurse -Filter *.proto $(PROTO_DIR) | ForEach-Object { $$_.FullName }")

# путь к стандартным proto (замени под свою установку protoc)
PROTOC_INC=C:\Users\ydlysak\protoc-32.0\include
PROTO_GRPC_GW_INC=C:\Users\ydlysak\go\pkg\mod\github.com\grpc-ecosystem\grpc-gateway\v2@v2.27.2

all: go gateway ts

go:
	protoc -I . -I $(PROTO_DIR) -I $(PROTOC_INC) -I $(PROTO_GRPC_GW_INC) \
		--go_out=$(GO_OUT) --go_opt=paths=source_relative \
		--go-grpc_out=$(GO_OUT) --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
		api/proto/coordinator/v1/*.proto \
		api/proto/probe/v1/*.proto \
		api/proto/host/v1/*.proto \
		api/proto/util/v1/*.proto \
		api/proto/bot/v1/*.proto \
		api/proto/whois/v1/*.proto

# https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_openapi_output/

gateway:
	protoc -I . -I $(PROTO_DIR) -I $(PROTOC_INC) -I $(PROTO_GRPC_GW_INC) \
		--grpc-gateway_out=$(GO_OUT) --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=$(GO_OUT) --openapiv2_opt=allow_merge=true \
		api/proto/coordinator/v1/*.proto \
		api/proto/probe/v1/*.proto \
		api/proto/host/v1/*.proto \
		api/proto/util/v1/*.proto \
		api/proto/bot/v1/*.proto \
		api/proto/whois/v1/*.proto

ts:
	protoc -I . -I $(PROTO_DIR) -I $(PROTOC_INC) -I $(PROTO_GRPC_GW_INC) \
		--ts_proto_out=$(TS_OUT) \
		--ts_proto_opt=esModuleInterop=true,forceLong=string,outputServices=grpc-js \
		api/proto/coordinator/v1/*.proto \
		api/proto/probe/v1/*.proto \
		api/proto/host/v1/*.proto \
		api/proto/util/v1/*.proto \
		api/proto/bot/v1/*.proto \
		api/proto/whois/v1/*.proto

clean:
	del /S /Q $(GO_OUT)\* || echo ""
	del /S /Q $(TS_OUT)\* || echo ""

.PHONY: all go gateway ts clean
