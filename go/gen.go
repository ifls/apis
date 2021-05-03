//go:generate mkdir -p auth
//go:generate protoc --go_out=./auth -I=../pb/auth auth.proto

//go:generate mkdir -p book
//go:generate protoc --go_out=./book -I=../pb/book book.proto

//go:generate mkdir -p login
//go:generate protoc --go_out=./login -I=../pb/login login.proto

//go:generate mkdir -p user
//go:generate protoc --go_out=./user -I=../pb/user user.proto
package _go
