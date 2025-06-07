package main

import (
	"irrig8r/api"
	"irrig8r/internal/mqtt"
	"irrig8r/internal/scheduler"
	"log"
	"net/http"
)

func main() {
	mqtt.Init()
	scheduler.Init()

	http.HandleFunc("/schedule", api.ScheduleHandler)
	http.HandleFunc("/pause", api.PauseHandler)

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
