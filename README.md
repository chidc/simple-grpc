# simple-grpc
run :   go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway                                                                                         
        go get google.golang.org/protobuf/cmd/protoc-gen-go                                                                                                              
        go get google.golang.org/grpc/cmd/protoc-gen-go-grpc                                                                                                             
download : https://github.com/protocolbuffers/protobuf/releases                                                                                                         

run main.go and client/main.go

to generate proto file : protoc -I ./proto  --go_out ./proto --go_opt paths=source_relative --go-grpc_out ./proto --go-grpc_opt paths=source_relative  ./proto/demo.proto

