gin:
	go run main.go -server=gin
grpc:
	go run main.go -server=grpc
protobuf:
	protoc -I=sdk/proto --go_out=sdk/pb  --go-grpc_out=sdk/pb sdk/proto/*.proto
