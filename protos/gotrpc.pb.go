// Code generated by protoc-gen-go.
// source: protos/gotrpc.proto
// DO NOT EDIT!

/*
Package gotrpc is a generated protocol buffer package.

It is generated from these files:
	protos/gotrpc.proto

It has these top-level messages:
	OtrMessage
*/
package gotrpc

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// The otr message
type OtrMessage struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *OtrMessage) Reset()         { *m = OtrMessage{} }
func (m *OtrMessage) String() string { return proto.CompactTextString(m) }
func (*OtrMessage) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Conversation service

type ConversationClient interface {
	Receive(ctx context.Context, in *OtrMessage, opts ...grpc.CallOption) (*OtrMessage, error)
	Send(ctx context.Context, in *OtrMessage, opts ...grpc.CallOption) (*OtrMessage, error)
}

type conversationClient struct {
	cc *grpc.ClientConn
}

func NewConversationClient(cc *grpc.ClientConn) ConversationClient {
	return &conversationClient{cc}
}

func (c *conversationClient) Receive(ctx context.Context, in *OtrMessage, opts ...grpc.CallOption) (*OtrMessage, error) {
	out := new(OtrMessage)
	err := grpc.Invoke(ctx, "/gotrpc.Conversation/Receive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) Send(ctx context.Context, in *OtrMessage, opts ...grpc.CallOption) (*OtrMessage, error) {
	out := new(OtrMessage)
	err := grpc.Invoke(ctx, "/gotrpc.Conversation/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Conversation service

type ConversationServer interface {
	Receive(context.Context, *OtrMessage) (*OtrMessage, error)
	Send(context.Context, *OtrMessage) (*OtrMessage, error)
}

func RegisterConversationServer(s *grpc.Server, srv ConversationServer) {
	s.RegisterService(&_Conversation_serviceDesc, srv)
}

func _Conversation_Receive_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(OtrMessage)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ConversationServer).Receive(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Conversation_Send_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(OtrMessage)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ConversationServer).Send(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Conversation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gotrpc.Conversation",
	HandlerType: (*ConversationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Receive",
			Handler:    _Conversation_Receive_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Conversation_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
