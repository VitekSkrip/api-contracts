GEN_SRC=./gen
PROTO_DIR=proto/*/v1/*.proto

GO_OUT_DIR=${GEN_SRC}/go/
OPENAPI_OUT_DIR=${GEN_SRC}/openapiv2

generate:
	make -j 3 gen-grpc gen-gw gen-openapi

gen-grpc:
	echo "Generating gRPC stubs and clients"
	protoc -I proto \
		--go_out=$(GO_OUT_DIR) \
		--go_opt=paths=source_relative \
	  	--go-grpc_out=$(GO_OUT_DIR) \
	  	--go-grpc_opt=paths=source_relative \
	   	$(PROTO_DIR)

gen-gw:
	echo "Generating gRPC Gateway bindings"
	protoc -I proto \
		--grpc-gateway_out $(GO_OUT_DIR) \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
		$(PROTO_DIR)

gen-openapi:
	echo "Generating OpenAPI specs"
	protoc -I proto \
		--openapiv2_out $(OPENAPI_OUT_DIR) \
		--openapiv2_opt logtostderr=true \
		--openapiv2_opt generate_unbound_methods=true \
		--openapiv2_opt allow_merge=true \
		--openapiv2_opt merge_file_name=api \
		$(PROTO_DIR)

.PHONY: generate gen-grpc gen-gw gen-openapi