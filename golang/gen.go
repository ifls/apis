//go:generate protoc --proto_path=../pb --go_out . --go_opt paths=source_relative book/book.proto
//go:generate protoc --proto_path=../pb  --go-grpc_out . --go-grpc_opt paths=source_relative book/book.proto
//go:generate protoc --proto_path=../pb  --grpc-gateway_out . --grpc-gateway_opt paths=source_relative --grpc-gateway_opt logtostderr=true  book/book.proto
package main

func main() {

}
