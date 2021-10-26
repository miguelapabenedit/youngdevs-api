package handlers

import (
	"encoding/json"
	"github/miguelapabenedit/youngdevs-api/app/data"
	"github/miguelapabenedit/youngdevs-api/app/repository"
	"net/http"
	"strconv"
)

var getAllUsersRepository repository.GetAllUsers

func NewGetAllUsers(repo repository.GetAllUsers) {
	getAllUsersRepository = repo
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	pageIndex := r.URL.Query().Get("page_index")
	pageSize := r.URL.Query().Get("page_size")

	if len(pageIndex) < 1 && len(pageSize) < 1 {
		http.Error(w, "Se debe eviar pageIndex y pageSize", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pi, err := strconv.Atoi(pageIndex)

	if err != nil {
		http.Error(w, "Page index is incorrect", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
	}

	ps, err := strconv.Atoi(pageSize)

	if err != nil {
		http.Error(w, "Page size is incorrect", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
	}

	users := data.UsersPaginated{
		PageSize:  ps,
		PageIndex: pi,
	}

	getAllUsersRepository.GetAllWithPagination(&users)

	msg, err := json.Marshal(&users)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(msg)
}
