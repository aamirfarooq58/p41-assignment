package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Timestamp string `json:"timestamp"`
	IP        string `json:"ip"`
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Extract IP address from request context
	ip := req.RequestContext.Identity.SourceIP
	if ip == "" {
		if xff, ok := req.Headers["X-Forwarded-For"]; ok {
			ip = xff
		} else {
			ip = "Unknown"
		}
	}

	// Handle IPv6 loopback
	if strings.HasPrefix(ip, "::1") {
		ip = "127.0.0.1"
	}

	now := time.Now().UTC()
	response := Response{
		Timestamp: now.Format(time.RFC3339),
		IP:        ip,
	}

	// Create JSON response
	responseBody, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"error": "failed to create response"}`,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(responseBody),
	}, nil
}

func main() {
	// Start Lambda function
	lambda.Start(handler)
}
