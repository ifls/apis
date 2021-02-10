//go:generate protoc --go_out=./auth -I=./auth auth.proto

//go:generate protoc --go_out=./book -I=./book book.proto

//go:generate protoc --go_out=./login -I=./login login.proto

//go:generate protoc --go_out=./user -I=./user user.proto
package pb
