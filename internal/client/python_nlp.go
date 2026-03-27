package client

import (
	"context"
	"time"

	pb "github.com/DennisMRitchie/go-nlp-service/proto/nlp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PythonNLPClient struct {
	conn   *grpc.ClientConn
	client pb.NLPServerClient
}

func NewPythonNLPClient(addr string) (*PythonNLPClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &PythonNLPClient{
		conn:   conn,
		client: pb.NewNLPServerClient(conn),
	}, nil
}

func (c *PythonNLPClient) Analyze(ctx context.Context, text, task string) (*pb.AnalyzeResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return c.client.AnalyzeText(ctx, &pb.AnalyzeRequest{
		Text: text,
		Task: task,
	})
}

func (c *PythonNLPClient) Close() error {
	return c.conn.Close()
}
