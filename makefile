create:
	protoc --go_out=createproto/. --go_opt=module=separategrpc/createproto --go-grpc_out=createproto/. --go-grpc_opt=module=separategrpc/createproto createproto/create.proto
read:
	protoc --go_out=readproto/. --go_opt=module=separategrpc/readproto --go-grpc_out=readproto/. --go-grpc_opt=module=separategrpc/readproto readproto/read.proto