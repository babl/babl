
brew update && brew install --devel protobuf
_make sure you install v3_

go get -a github.com/golang/protobuf/protoc-gen-go

protoc -I ./protobuf/ ./protobuf/helloworld.proto --go_out=plugins=grpc:protobuf

go-bindata data/...
