server-gen:
	protoc --proto_path=./proto --go_out=./server --go-grpc_out=./server proto/movieService.proto
	protoc --proto_path=./proto --go_out=./server --go-grpc_out=./server proto/hallService.proto
	echo "server generated"