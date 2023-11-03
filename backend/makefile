server-gen:
	protoc --proto_path=./proto --go_out=./server --go-grpc_out=./server proto/movieService.proto
	protoc --proto_path=./proto --go_out=./server --go-grpc_out=./server proto/hallService.proto
	protoc --proto_path=./proto --go_out=./server --go-grpc_out=./server proto/ticketService.proto
	echo "server generated"

ticket-gen:
	protoc --proto_path=./proto --go_out=./services/ticket --go-grpc_out=./services/ticket proto/movieService.proto
	protoc --proto_path=./proto --go_out=./services/ticket --go-grpc_out=./services/ticket proto/hallService.proto
	protoc --proto_path=./proto --go_out=./services/ticket --go-grpc_out=./services/ticket proto/ticketService.proto
	echo "ticket generated"