// Code generated by bin/module-mapping
// source: config.yml
// DO NOT EDIT!

package babl

import (
	grpc "google.golang.org/grpc"

	larskluge "github.com/larskluge/babl/protobuf/modules/larskluge"
	mondoreale "github.com/larskluge/babl/protobuf/modules/mondoreale"
	omnisyle "github.com/larskluge/babl/protobuf/modules/omnisyle"
)


var Modules = map[string]Component{
	"larskluge/babl-build": Component{
		Client: func(cc *grpc.ClientConn) BinaryClient {
			return BinaryClient(larskluge.NewBablBuildClient(cc))
		},
		Server: func(s *grpc.Server, srv BinaryServer) {
			larskluge.RegisterBablBuildServer(s, srv)
		},
  },
	"larskluge/babl-deploy": Component{
		Client: func(cc *grpc.ClientConn) BinaryClient {
			return BinaryClient(larskluge.NewBablDeployClient(cc))
		},
		Server: func(s *grpc.Server, srv BinaryServer) {
			larskluge.RegisterBablDeployServer(s, srv)
		},
  },
	"larskluge/babl-init-module": Component{
		Client: func(cc *grpc.ClientConn) BinaryClient {
			return BinaryClient(larskluge.NewBablInitModuleClient(cc))
		},
		Server: func(s *grpc.Server, srv BinaryServer) {
			larskluge.RegisterBablInitModuleServer(s, srv)
		},
  },
	"larskluge/bar": Component{
		Client: func(cc *grpc.ClientConn) BinaryClient {
			return BinaryClient(larskluge.NewBarClient(cc))
		},
		Server: func(s *grpc.Server, srv BinaryServer) {
			larskluge.RegisterBarServer(s, srv)
		},
  },
	"larskluge/baz": Component{
		Client: func(cc *grpc.ClientConn) BinaryClient {
			return BinaryClient(larskluge.NewBazClient(cc))
		},
		Server: func(s *grpc.Server, srv BinaryServer) {
			larskluge.RegisterBazServer(s, srv)
		},
  },
	"larskluge/download": Component{
		Client: func(cc *grpc.ClientConn) BinaryClient {
			return BinaryClient(larskluge.NewDownloadClient(cc))
		},
		Server: func(s *grpc.Server, srv BinaryServer) {
			larskluge.RegisterDownloadServer(s, srv)
		},
  },
	"larskluge/foo": Component{
		Client: func(cc *grpc.ClientConn) BinaryClient {
			return BinaryClient(larskluge.NewFooClient(cc))
		},
		Server: func(s *grpc.Server, srv BinaryServer) {
			larskluge.RegisterFooServer(s, srv)
		},
  },
	"larskluge/image-resize": Component{
		Client: func(cc *grpc.ClientConn) BinaryClient {
			return BinaryClient(larskluge.NewImageResizeClient(cc))
		},
		Server: func(s *grpc.Server, srv BinaryServer) {
			larskluge.RegisterImageResizeServer(s, srv)
		},
  },
	"larskluge/render-website": Component{
		Client: func(cc *grpc.ClientConn) BinaryClient {
			return BinaryClient(larskluge.NewRenderWebsiteClient(cc))
		},
		Server: func(s *grpc.Server, srv BinaryServer) {
			larskluge.RegisterRenderWebsiteServer(s, srv)
		},
  },
	"larskluge/s3": Component{
		Client: func(cc *grpc.ClientConn) BinaryClient {
			return BinaryClient(larskluge.NewS3Client(cc))
		},
		Server: func(s *grpc.Server, srv BinaryServer) {
			larskluge.RegisterS3Server(s, srv)
		},
  },
	"larskluge/string-append": Component{
		Client: func(cc *grpc.ClientConn) BinaryClient {
			return BinaryClient(larskluge.NewStringAppendClient(cc))
		},
		Server: func(s *grpc.Server, srv BinaryServer) {
			larskluge.RegisterStringAppendServer(s, srv)
		},
  },
	"larskluge/string-upcase": Component{
		Client: func(cc *grpc.ClientConn) BinaryClient {
			return BinaryClient(larskluge.NewStringUpcaseClient(cc))
		},
		Server: func(s *grpc.Server, srv BinaryServer) {
			larskluge.RegisterStringUpcaseServer(s, srv)
		},
  },
	"larskluge/test-fail": Component{
		Client: func(cc *grpc.ClientConn) BinaryClient {
			return BinaryClient(larskluge.NewTestFailClient(cc))
		},
		Server: func(s *grpc.Server, srv BinaryServer) {
			larskluge.RegisterTestFailServer(s, srv)
		},
  },
	"mondoreale/foobar": Component{
		Client: func(cc *grpc.ClientConn) BinaryClient {
			return BinaryClient(mondoreale.NewFoobarClient(cc))
		},
		Server: func(s *grpc.Server, srv BinaryServer) {
			mondoreale.RegisterFoobarServer(s, srv)
		},
  },
	"omnisyle/official-module": Component{
		Client: func(cc *grpc.ClientConn) BinaryClient {
			return BinaryClient(omnisyle.NewOfficialModuleClient(cc))
		},
		Server: func(s *grpc.Server, srv BinaryServer) {
			omnisyle.RegisterOfficialModuleServer(s, srv)
		},
  },
	"omnisyle/test": Component{
		Client: func(cc *grpc.ClientConn) BinaryClient {
			return BinaryClient(omnisyle.NewTestClient(cc))
		},
		Server: func(s *grpc.Server, srv BinaryServer) {
			omnisyle.RegisterTestServer(s, srv)
		},
  },
	"omnisyle/test2private": Component{
		Client: func(cc *grpc.ClientConn) BinaryClient {
			return BinaryClient(omnisyle.NewTest2PrivateClient(cc))
		},
		Server: func(s *grpc.Server, srv BinaryServer) {
			omnisyle.RegisterTest2PrivateServer(s, srv)
		},
  },
}
