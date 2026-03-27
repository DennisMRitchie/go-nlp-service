package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DennisMRitchie/go-nlp-service/internal/client"
	"github.com/DennisMRitchie/go-nlp-service/internal/handler"
	"github.com/DennisMRitchie/go-nlp-service/internal/service"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func initTracer() func(context.Context) error {
	exporter, err := otlptracegrpc.New(context.Background(),
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint("otel-collector:4317"),
	)
	if err != nil {
		log.Printf("Warning: failed to create trace exporter: %v (tracing disabled)", err)
		return func(context.Context) error { return nil }
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("go-nlp-service"),
		)),
	)
	otel.SetTracerProvider(tp)
	return tp.Shutdown
}

func main() {
	// Initialize Tracer
	shutdown := initTracer()
	defer shutdown(context.Background())

	// Python NLP Client
	pythonAddr := os.Getenv("PYTHON_NLP_ADDR")
	if pythonAddr == "" {
		pythonAddr = "python-nlp:50051"
	}

	pythonClient, err := client.NewPythonNLPClient(pythonAddr)
	if err != nil {
		log.Fatalf("Failed to connect to Python NLP service: %v", err)
	}
	defer pythonClient.Close()

	nlpService := service.NewNLPService(pythonClient)
	h := handler.NewHandler(nlpService)

	r := gin.Default()
	r.Use(otelgin.Middleware("go-nlp-service"))

	api := r.Group("/api/v1")
	{
		api.POST("/analyze", h.Analyze)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Graceful shutdown
	go func() {
		log.Printf("Server starting on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
