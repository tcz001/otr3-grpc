/*
 *
 * Copyright 2015, Google Inc.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *     * Redistributions of source code must retain the above copyright
 * notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above
 * copyright notice, this list of conditions and the following disclaimer
 * in the documentation and/or other materials provided with the
 * distribution.
 *     * Neither the name of Google Inc. nor the names of its
 * contributors may be used to endorse or promote products derived from
 * this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

package main

import (
	"crypto/rand"
	"log"
	"net"

	pb "github.com/tcz001/otr3-grpc/protos"
	"github.com/twstrike/otr3"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement server.OTRService.
type server struct {
	c *otr3.Conversation
}

// Receive implements server.Receive
func (s *server) Receive(ctx context.Context, in *pb.OtrRequest) (*pb.OtrResponse, error) {
	s.ensureConv()
	plain, toSend, err := s.c.Receive(otr3.ValidMessage(in.Message))
	if err != nil {
		return &pb.OtrResponse{Error: err.Error()}, nil
	}
	if toSend == nil {
		return &pb.OtrResponse{Plain: string(plain)}, nil
	}
	return &pb.OtrResponse{Plain: string(plain), ToSend: string(toSend[0])}, nil
}

// Send implements server.Send
func (s *server) Send(ctx context.Context, in *pb.OtrRequest) (*pb.OtrResponse, error) {
	s.ensureConv()
	toSend, err := s.c.Send(otr3.ValidMessage(in.Message))
	if err != nil {
		return &pb.OtrResponse{Error: err.Error()}, nil
	}
	return &pb.OtrResponse{ToSend: string(toSend[0])}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOTRServiceServer(s, &server{})
	s.Serve(lis)
}

func (s *server) ensureConv() {
	if s.c == nil {
		c := otr3.Conversation{}

		// You will need to prepare a long-term PrivateKey for otr conversation handshakes.
		priv := &otr3.PrivateKey{}
		priv.Generate(rand.Reader)
		c.SetKeys(priv, nil)

		// set the Policies.
		c.Policies.AllowV2()
		c.Policies.AllowV3()
		c.Policies.RequireEncryption()

		s.c = &c
	}
}
