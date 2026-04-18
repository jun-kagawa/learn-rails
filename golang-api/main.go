package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	// "os"
	"runtime"
	"time"
)

type Hello struct {
	Hello string    `json:"hello"`
	T     time.Time `json:"time"`
}

func main() {
	logger := slog.New(slog.NewJSONHandler(io.Discard, nil))
	logger.Info("GOMAXPROCS", "count", runtime.GOMAXPROCS(0))
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("start")
		// time.Sleep(1 * time.Second)
		// logger.Info("wait done")
		res := Hello{
			Hello: "world",
			T:     time.Now(),
		}
		logger.Info("marshal json")
		buf, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		logger.Info("write body")
		if _, err := w.Write(buf); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logger.Info("done")
	})
	server := &http.Server{
		Addr:        ":3000",
		Handler:     mux,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 5 * time.Second,
	}
	server.ListenAndServe()
}
