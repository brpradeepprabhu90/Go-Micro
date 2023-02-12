package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"log-service/data"
	"log-service/logs"
	"net"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{
			Result: "failed",
		}
		return res, err
	}
	res := &logs.LogResponse{
		Result: "logged",
	}

	return res, nil

}
func (app *Config) gRpcListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("failed to listen to gRpc %v", err)
	}
	s := grpc.NewServer()
	logs.RegisterLogServiceServer(s, &LogServer{
		Models: app.Models,
	})
	log.Println("grpc started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen to gRpc %v", err)
	}
}
