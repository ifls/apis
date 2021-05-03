//go:generate protoc --go_out=./auth -I=./auth auth.proto

//go:generate protoc --go_out=./book -I=./book book.proto

//go:generate protoc --go_out=./login -I=./login login.proto

//go:generate protoc --go_out=./user -I=./user user.proto

//go:generate protoc -I=.;/Volumes/c/gitRepo/github/googleapis --go_out ./go/ --go_opt paths=source_relative --go-grpc_out ./go/ --go-grpc_opt paths=./book  ./book/book.proto
package pb
