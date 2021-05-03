//go:generate protoc --proto_path=. --go_out ../go/ --go_opt paths=source_relative ./book/book.proto
//go:generate protoc --proto_path=.  --go-grpc_out ../go/ --go-grpc_opt paths=source_relative ./book/book.proto
//go:generate protoc --proto_path=.  --grpc-gateway_out ../go/ --grpc-gateway_opt paths=source_relative --grpc-gateway_opt logtostderr=true  ./book/book.proto
package pb
