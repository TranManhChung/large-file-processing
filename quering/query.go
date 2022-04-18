package quering

import (
	"context"
	"encoding/json"
	"github.com/TranManhChung/large-file-processing/parser"
	"net/http"
	"strconv"
)

type GetResponse struct {
	Status  string         `json:"status"`
	Message string         `json:"message,omitempty"`
	Prices  []parser.Price `json:"prices,omitempty"`
}

func (s Service) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		json.NewEncoder(w).Encode(GetResponse{
			Status:  "failed",
			Message: "offset is invalid",
		})
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		json.NewEncoder(w).Encode(GetResponse{
			Status:  "failed",
			Message: "limit is invalid",
		})
		return
	}
	symbol := r.URL.Query().Get("symbol")
	if symbol == "" {
		json.NewEncoder(w).Encode(GetResponse{
			Status:  "failed",
			Message: "symbol is invalid",
		})
		return
	}

	prices, err := s.PriceRepo.GetByOffsetLimit(context.Background(), symbol, offset, limit)
	if err != nil {
		json.NewEncoder(w).Encode(GetResponse{
			Status:  "failed",
			Message: "internal error",
		})
		return
	}

	json.NewEncoder(w).Encode(GetResponse{
		Status: "success",
		Prices: prices,
	})
}
