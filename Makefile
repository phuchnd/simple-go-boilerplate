gen-proto:
	echo "Generate proto"
	protoc \
		--go_out=server/grpc/pb --go_opt=paths=source_relative \
		--go-grpc_out=server/grpc/pb --go-grpc_opt=paths=source_relative \
		--proto_path=server/grpc/pb \
		server/grpc/pb/*.proto

go-gen:
	echo "Generate mocks"
	go generate ./...
	echo "Format code"
	go fmt ./...

test-all:
	echo "Test all"
	ginkgo -r ./...

test-unit:
	echo "Test unit tests"
	ginkgo -r --focus unit

test-integration:
	echo "Test integration tests"
	ginkgo -r --focus integration