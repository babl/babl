# Babl

Client app to access the Babl Network.


### Collection of old commands

```
# make sure you install v3
brew update && brew install --devel protobuf

go get -a github.com/golang/protobuf/protoc-gen-go

protoc -I ./protobuf/ ./protobuf/babl.proto --go_out=plugins=grpc:protobuf
```
