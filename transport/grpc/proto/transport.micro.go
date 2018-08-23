// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/jinbanglin/go-plugins/transport/grpc/proto/transport.proto

/*
Package go_micro_grpc_transport is a generated protocol buffer package.

It is generated from these files:
	github.com/jinbanglin/go-plugins/transport/grpc/proto/transport.proto

It has these top-level messages:
	Message
*/
package go_micro_grpc_transport

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/jinbanglin/go-micro/client"
	server "github.com/jinbanglin/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Transport service

type TransportService interface {
	Stream(ctx context.Context, opts ...client.CallOption) (Transport_StreamService, error)
}

type transportService struct {
	c    client.Client
	name string
}

func NewTransportService(name string, c client.Client) TransportService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.grpc.transport"
	}
	return &transportService{
		c:    c,
		name: name,
	}
}

func (c *transportService) Stream(ctx context.Context, opts ...client.CallOption) (Transport_StreamService, error) {
	req := c.c.NewRequest(c.name, "Transport.Stream", &Message{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &transportServiceStream{stream}, nil
}

type Transport_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Message) error
	Recv() (*Message, error)
}

type transportServiceStream struct {
	stream client.Stream
}

func (x *transportServiceStream) Close() error {
	return x.stream.Close()
}

func (x *transportServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *transportServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *transportServiceStream) Send(m *Message) error {
	return x.stream.Send(m)
}

func (x *transportServiceStream) Recv() (*Message, error) {
	m := new(Message)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Transport service

type TransportHandler interface {
	Stream(context.Context, Transport_StreamStream) error
}

func RegisterTransportHandler(s server.Server, hdlr TransportHandler, opts ...server.HandlerOption) {
	type transport interface {
		Stream(ctx context.Context, stream server.Stream) error
	}
	type Transport struct {
		transport
	}
	h := &transportHandler{hdlr}
	s.Handle(s.NewHandler(&Transport{h}, opts...))
}

type transportHandler struct {
	TransportHandler
}

func (h *transportHandler) Stream(ctx context.Context, stream server.Stream) error {
	return h.TransportHandler.Stream(ctx, &transportStreamStream{stream})
}

type Transport_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Message) error
	Recv() (*Message, error)
}

type transportStreamStream struct {
	stream server.Stream
}

func (x *transportStreamStream) Close() error {
	return x.stream.Close()
}

func (x *transportStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *transportStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *transportStreamStream) Send(m *Message) error {
	return x.stream.Send(m)
}

func (x *transportStreamStream) Recv() (*Message, error) {
	m := new(Message)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
