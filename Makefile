gen-proto:
	protoc \
		--go_out=server/grpc --go_opt=paths=source_relative \
		--go-grpc_out=server/grpc --go-grpc_opt=paths=source_relative \
		--proto_path=server/grpc \
		server/grpc/*.proto
