// Code generated by protoc-gen-go.
// source: protos/gotrpc.proto
// DO NOT EDIT!

/*
Package gotrpc is a generated protocol buffer package.

It is generated from these files:
	protos/gotrpc.proto

It has these top-level messages:
	OtrMsgResponse
	OtrMsgRequest
	OtrConvRequest
	OtrConvResponse
*/
package gotrpc

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// The otr Response message
type OtrMsgResponse struct {
	Plain  string `protobuf:"bytes,1,opt,name=plain" json:"plain,omitempty"`
	ToSend string `protobuf:"bytes,2,opt,name=toSend" json:"toSend,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *OtrMsgResponse) Reset()         { *m = OtrMsgResponse{} }
func (m *OtrMsgResponse) String() string { return proto.CompactTextString(m) }
func (*OtrMsgResponse) ProtoMessage()    {}

// The otr Request message
type OtrMsgRequest struct {
	Uuid    string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *OtrMsgRequest) Reset()         { *m = OtrMsgRequest{} }
func (m *OtrMsgRequest) String() string { return proto.CompactTextString(m) }
func (*OtrMsgRequest) ProtoMessage()    {}

type OtrConvRequest struct {
	Uuid string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
}

func (m *OtrConvRequest) Reset()         { *m = OtrConvRequest{} }
func (m *OtrConvRequest) String() string { return proto.CompactTextString(m) }
func (*OtrConvRequest) ProtoMessage()    {}

type OtrConvResponse struct {
	Uuid  string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *OtrConvResponse) Reset()         { *m = OtrConvResponse{} }
func (m *OtrConvResponse) String() string { return proto.CompactTextString(m) }
func (*OtrConvResponse) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for OTRService service

type OTRServiceClient interface {
	NewConv(ctx context.Context, in *OtrConvRequest, opts ...grpc.CallOption) (*OtrConvResponse, error)
	Receive(ctx context.Context, in *OtrMsgRequest, opts ...grpc.CallOption) (*OtrMsgResponse, error)
	Send(ctx context.Context, in *OtrMsgRequest, opts ...grpc.CallOption) (*OtrMsgResponse, error)
}

type oTRServiceClient struct {
	cc *grpc.ClientConn
}

func NewOTRServiceClient(cc *grpc.ClientConn) OTRServiceClient {
	return &oTRServiceClient{cc}
}

func (c *oTRServiceClient) NewConv(ctx context.Context, in *OtrConvRequest, opts ...grpc.CallOption) (*OtrConvResponse, error) {
	out := new(OtrConvResponse)
	err := grpc.Invoke(ctx, "/gotrpc.OTRService/NewConv", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oTRServiceClient) Receive(ctx context.Context, in *OtrMsgRequest, opts ...grpc.CallOption) (*OtrMsgResponse, error) {
	out := new(OtrMsgResponse)
	err := grpc.Invoke(ctx, "/gotrpc.OTRService/Receive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oTRServiceClient) Send(ctx context.Context, in *OtrMsgRequest, opts ...grpc.CallOption) (*OtrMsgResponse, error) {
	out := new(OtrMsgResponse)
	err := grpc.Invoke(ctx, "/gotrpc.OTRService/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OTRService service

type OTRServiceServer interface {
	NewConv(context.Context, *OtrConvRequest) (*OtrConvResponse, error)
	Receive(context.Context, *OtrMsgRequest) (*OtrMsgResponse, error)
	Send(context.Context, *OtrMsgRequest) (*OtrMsgResponse, error)
}

func RegisterOTRServiceServer(s *grpc.Server, srv OTRServiceServer) {
	s.RegisterService(&_OTRService_serviceDesc, srv)
}

func _OTRService_NewConv_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(OtrConvRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(OTRServiceServer).NewConv(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _OTRService_Receive_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(OtrMsgRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(OTRServiceServer).Receive(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _OTRService_Send_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(OtrMsgRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(OTRServiceServer).Send(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _OTRService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gotrpc.OTRService",
	HandlerType: (*OTRServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewConv",
			Handler:    _OTRService_NewConv_Handler,
		},
		{
			MethodName: "Receive",
			Handler:    _OTRService_Receive_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _OTRService_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
