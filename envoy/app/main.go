/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a server for Greeter service.
package main

import (
    "math/rand"
    "time"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "helloworld"
)

const (
	port = ":50051"
    letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

func RandString(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(in *pb.HelloRequest, stream pb.Greeter_SayHelloServer) error {
	log.Printf("Received: %v bytes' string will send %v times.", in.GetTimes())

    var details []*pb.HelloDetail
    for i := 0; i < 10; i++ {
        if rand.Intn(2) == 1 {
            details = append(details, &pb.HelloDetail{
                Hoge: RandString(10),
                Fuga: RandString(10),
                Piyo: RandString(10),
            })
        }
    }

    res := &pb.HelloReply{
        Id: RandString(10),
        From: RandString(10),
        Body: RandString(10),
        Details: details,
    }
	times := int(in.GetTimes())

	for i := 0; i < times; i++ {
		if err := stream.Send(res); err != nil {
			return err
		}
        time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
