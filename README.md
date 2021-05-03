# apis
api


## grpc-gateway 安装步骤

go get -u -v github.com/grpc-ecosystem/grpc-gateway
go get -u -v google.golang.org/protobuf
go get -u -v google.golang.org/grpc


go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

查看 GOBIN 目录看是否安装了