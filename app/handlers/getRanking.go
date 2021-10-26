package handlers

import (
	"encoding/json"
	"github/miguelapabenedit/youngdevs-api/app/data"
	"github/miguelapabenedit/youngdevs-api/app/repository"
	"net/http"
	"strings"
)

var getRankingRepository repository.GetRanking

func NewGetRanking(repo repository.GetRanking) {
	getRankingRepository = repo
}

func GetRanking(w http.ResponseWriter, r *http.Request) {
	u := []data.User{}
	getRankingRepository.GetRanking(&u)

	ur := []data.Ranking{}
	for _, u := range u {
		ur = append(ur, data.Ranking{Name: strings.TrimRight(u.Email, "@"), Score: u.Score})
	}

	msg, err := json.Marshal(&ur)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(msg)
}
