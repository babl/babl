brew update && brew install --devel protobuf
_make sure you install v3_

go get -a github.com/golang/protobuf/protoc-gen-go

protoc -I ./protobuf/ ./protobuf/babl.proto --go_out=plugins=grpc:protobuf

cd shared && go-bindata -pkg shared data/...

go-bindata data/...

gox -osarch="linux/amd64"
