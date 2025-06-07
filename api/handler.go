package api

import (
	"encoding/json"
	"irrig8r/internal/controller"
	"net/http"
)

type ScheduleRequest struct {
	Zone     string `json:"zone"`
	Start    string `json:"start"`
	Duration int    `json:"duration"`
}

func ScheduleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ScheduleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := controller.AddSchedule(req.Zone, req.Start, req.Duration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Schedule added"))
}

func PauseHandler(w http.ResponseWriter, r *http.Request) {
	minutes := r.URL.Query().Get("minutes")
	controller.Pause(minutes)
	w.Write([]byte("Pause set"))
}
