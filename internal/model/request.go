package model

type AnalyzeRequest struct {
    Text string `json:"text" binding:"required"`
    Task string `json:"task" binding:"required,oneof=sentiment classification entities"`
}

type AnalyzeResponse struct {
    Result           string             `json:"result"`
    Confidence       float64            `json:"confidence"`
    Entities         map[string]float64 `json:"entities,omitempty"`
    ProcessingTimeMs int64              `json:"processing_time_ms"`
}
