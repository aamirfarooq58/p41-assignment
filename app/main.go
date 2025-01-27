package main

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"
	"time"
)

type Response struct {
	Timestamp string `json:"timestamp"`
	IP        string `json:"ip"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Extract IP address from RemoteAddr
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		// Handle IPv6 loopback
		if strings.HasPrefix(ip, "::1") {
			ip = "127.0.0.1"
		}

		now := time.Now().UTC()
		response := Response{
			Timestamp: now.Format(time.RFC3339),
			IP:        ip,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8123", nil)
}
