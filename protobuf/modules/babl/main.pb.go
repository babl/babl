// Code generated by protoc-gen-go.
// source: github.com/larskluge/babl/protobuf/modules/babl/main.proto
// DO NOT EDIT!

/*
Package babl_babl is a generated protocol buffer package.

It is generated from these files:
	github.com/larskluge/babl/protobuf/modules/babl/main.proto

It has these top-level messages:
*/
package babl_babl

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import babl "github.com/larskluge/babl/protobuf/messages"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Job service

type JobClient interface {
	IO(ctx context.Context, in *babl.BinRequest, opts ...grpc.CallOption) (*babl.BinReply, error)
	Ping(ctx context.Context, in *babl.Empty, opts ...grpc.CallOption) (*babl.Pong, error)
}

type jobClient struct {
	cc *grpc.ClientConn
}

func NewJobClient(cc *grpc.ClientConn) JobClient {
	return &jobClient{cc}
}

func (c *jobClient) IO(ctx context.Context, in *babl.BinRequest, opts ...grpc.CallOption) (*babl.BinReply, error) {
	out := new(babl.BinReply)
	err := grpc.Invoke(ctx, "/babl.babl.Job/IO", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) Ping(ctx context.Context, in *babl.Empty, opts ...grpc.CallOption) (*babl.Pong, error) {
	out := new(babl.Pong)
	err := grpc.Invoke(ctx, "/babl.babl.Job/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Job service

type JobServer interface {
	IO(context.Context, *babl.BinRequest) (*babl.BinReply, error)
	Ping(context.Context, *babl.Empty) (*babl.Pong, error)
}

func RegisterJobServer(s *grpc.Server, srv JobServer) {
	s.RegisterService(&_Job_serviceDesc, srv)
}

func _Job_IO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(babl.BinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(JobServer).IO(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Job_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(babl.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(JobServer).Ping(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Job_serviceDesc = grpc.ServiceDesc{
	ServiceName: "babl.babl.Job",
	HandlerType: (*JobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IO",
			Handler:    _Job_IO_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Job_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for NsqJob service

type NsqJobClient interface {
	IO(ctx context.Context, in *babl.BinRequest, opts ...grpc.CallOption) (*babl.BinReply, error)
	Ping(ctx context.Context, in *babl.Empty, opts ...grpc.CallOption) (*babl.Pong, error)
}

type nsqJobClient struct {
	cc *grpc.ClientConn
}

func NewNsqJobClient(cc *grpc.ClientConn) NsqJobClient {
	return &nsqJobClient{cc}
}

func (c *nsqJobClient) IO(ctx context.Context, in *babl.BinRequest, opts ...grpc.CallOption) (*babl.BinReply, error) {
	out := new(babl.BinReply)
	err := grpc.Invoke(ctx, "/babl.babl.NsqJob/IO", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsqJobClient) Ping(ctx context.Context, in *babl.Empty, opts ...grpc.CallOption) (*babl.Pong, error) {
	out := new(babl.Pong)
	err := grpc.Invoke(ctx, "/babl.babl.NsqJob/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NsqJob service

type NsqJobServer interface {
	IO(context.Context, *babl.BinRequest) (*babl.BinReply, error)
	Ping(context.Context, *babl.Empty) (*babl.Pong, error)
}

func RegisterNsqJobServer(s *grpc.Server, srv NsqJobServer) {
	s.RegisterService(&_NsqJob_serviceDesc, srv)
}

func _NsqJob_IO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(babl.BinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NsqJobServer).IO(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _NsqJob_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(babl.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NsqJobServer).Ping(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _NsqJob_serviceDesc = grpc.ServiceDesc{
	ServiceName: "babl.babl.NsqJob",
	HandlerType: (*NsqJobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IO",
			Handler:    _NsqJob_IO_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _NsqJob_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
