GO_WORKSPACE := $(GOPATH)/src

protoc:
	protoc *.proto --go_out=$(GO_WORKSPACE) --go-grpc_out=$(GO_WORKSPACE)
	@echo "Protoc compile success"