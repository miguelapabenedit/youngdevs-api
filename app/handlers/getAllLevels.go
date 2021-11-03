package handlers

import (
	"encoding/json"
	"net/http"
)

func GetAllLevels(w http.ResponseWriter, r *http.Request) {
	levels := getLevelRepository.GetAllLevels()

	msg, err := json.Marshal(levels)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(msg)
}
