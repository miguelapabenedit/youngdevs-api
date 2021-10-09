package handlers

import (
	"encoding/json"
	"github/miguelapabenedit/youngdevs-api/app/repository"
	"net/http"
	"strconv"
)

var getLevelRepository repository.GetLevel

func NewGetLevel(repo repository.GetLevel) {
	getLevelRepository = repo
}

func GetLevel(w http.ResponseWriter, r *http.Request) {
	lvlHeader := r.URL.Query().Get("level")

	if len(lvlHeader) < 1 {
		http.Error(w, "Debe enviarse un level", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lvlInt, err := strconv.Atoi(lvlHeader)

	if err != nil {
		http.Error(w, "El nivel es incorrecto", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
	}

	levelObj := getLevelRepository.GetLevel(lvlInt)

	msg, err := json.Marshal(levelObj)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(msg)
}
