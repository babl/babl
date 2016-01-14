// Code generated by bin/module-mapping
// source: babl.proto
// DO NOT EDIT!


func NewBablBuildClient2(cc *grpc.ClientConn) BinaryClient {
	return &bablBuildClient{cc}
}
func RegisterBablBuildServer2(s *grpc.Server, srv BinaryServer) {
	s.RegisterService(&_BablBuild_serviceDesc, srv)
}

func NewStringUpcaseClient2(cc *grpc.ClientConn) BinaryClient {
	return &stringUpcaseClient{cc}
}
func RegisterStringUpcaseServer2(s *grpc.Server, srv BinaryServer) {
	s.RegisterService(&_StringUpcase_serviceDesc, srv)
}

func NewStringAppendClient2(cc *grpc.ClientConn) BinaryClient {
	return &stringAppendClient{cc}
}
func RegisterStringAppendServer2(s *grpc.Server, srv BinaryServer) {
	s.RegisterService(&_StringAppend_serviceDesc, srv)
}

func NewDownloadClient2(cc *grpc.ClientConn) BinaryClient {
	return &downloadClient{cc}
}
func RegisterDownloadServer2(s *grpc.Server, srv BinaryServer) {
	s.RegisterService(&_Download_serviceDesc, srv)
}

func NewS3Client2(cc *grpc.ClientConn) BinaryClient {
	return &s3Client{cc}
}
func RegisterS3Server2(s *grpc.Server, srv BinaryServer) {
	s.RegisterService(&_S3_serviceDesc, srv)
}

func NewImageResizeClient2(cc *grpc.ClientConn) BinaryClient {
	return &imageResizeClient{cc}
}
func RegisterImageResizeServer2(s *grpc.Server, srv BinaryServer) {
	s.RegisterService(&_ImageResize_serviceDesc, srv)
}

func NewRenderWebsiteClient2(cc *grpc.ClientConn) BinaryClient {
	return &renderWebsiteClient{cc}
}
func RegisterRenderWebsiteServer2(s *grpc.Server, srv BinaryServer) {
	s.RegisterService(&_RenderWebsite_serviceDesc, srv)
}

func NewTestFailClient2(cc *grpc.ClientConn) BinaryClient {
	return &testFailClient{cc}
}
func RegisterTestFailServer2(s *grpc.Server, srv BinaryServer) {
	s.RegisterService(&_TestFail_serviceDesc, srv)
}

var Modules = map[string]Component{
	"babl-build": Component{Client: NewBablBuildClient2, Server: RegisterBablBuildServer2},
	"string-upcase": Component{Client: NewStringUpcaseClient2, Server: RegisterStringUpcaseServer2},
	"string-append": Component{Client: NewStringAppendClient2, Server: RegisterStringAppendServer2},
	"download": Component{Client: NewDownloadClient2, Server: RegisterDownloadServer2},
	"s3": Component{Client: NewS3Client2, Server: RegisterS3Server2},
	"image-resize": Component{Client: NewImageResizeClient2, Server: RegisterImageResizeServer2},
	"render-website": Component{Client: NewRenderWebsiteClient2, Server: RegisterRenderWebsiteServer2},
	"test-fail": Component{Client: NewTestFailClient2, Server: RegisterTestFailServer2},
}
