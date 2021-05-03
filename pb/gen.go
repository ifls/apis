//go:generate protoc --go_out=./auth -I=./auth auth.proto

//go:generate protoc --go_out=./book -I=./book book.proto

//go:generate protoc --go_out=./login -I=./login login.proto

//go:generate protoc --go_out=./user -I=./user user.proto

//go:generate protoc --proto_path=. --go_out ../go/  --go-grpc_out ../go/ --grpc-gateway_opt paths=source_relative ./book/book.proto
package pb
