package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type ServerStats struct {
	ServerHitCount int `json:"server_hit_count"`
}

type APIConfig struct {
	ServerStats ServerStats
	DB          *sqlx.DB
}

func NewAPIConfig(db *sqlx.DB) *APIConfig {
	return &APIConfig{
		ServerStats: ServerStats{
			ServerHitCount: 0,
		},
		DB: db,
	}
}
func (a *APIConfig) MetricsHandler(w http.ResponseWriter, r *http.Request) {
	type returnVals struct {
		ServerMetrics string `json:"server_metrics"`
	}

	RespondWithJSON(w, http.StatusOK, returnVals{
		ServerMetrics: a.GetServerMetrics(),
	})
}

func (a *APIConfig) GetServerMetrics() string {
	return a.ServerHitCountString()
}

func (a *APIConfig) ServerHitCountString() string {
	return strconv.Itoa(a.ServerStats.ServerHitCount)
}

func (a *APIConfig) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func (a *APIConfig) ServiceNowWebHookHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}
