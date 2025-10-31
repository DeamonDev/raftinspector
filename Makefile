.PHONY: proto build

PROTO_DIR=./proto
PROTO_FILES=$(PROTO_DIR)/raft_inspector.proto

proto:
	protoc --proto_path=$(PROTO_DIR) \
		--go_out=$(PROTO_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)