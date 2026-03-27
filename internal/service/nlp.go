package service

import (
	"context"
	"time"

	"github.com/DennisMRitchie/go-nlp-service/internal/client"
	"github.com/DennisMRitchie/go-nlp-service/internal/model"
)

type NLPService struct {
	pythonClient *client.PythonNLPClient
}

func NewNLPService(pythonClient *client.PythonNLPClient) *NLPService {
	return &NLPService{pythonClient: pythonClient}
}

func (s *NLPService) Analyze(ctx context.Context, req model.AnalyzeRequest) (*model.AnalyzeResponse, error) {
	start := time.Now()

	resp, err := s.pythonClient.Analyze(ctx, req.Text, req.Task)
	if err != nil {
		return nil, err
	}

	duration := time.Since(start).Milliseconds()

	return &model.AnalyzeResponse{
		Result:           resp.Result,
		Confidence:       float64(resp.Confidence),
		Entities:         resp.Entities,
		ProcessingTimeMs: duration,
	}, nil
}
