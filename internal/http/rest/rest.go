package rest

import (
	"log"
	"net/http"
	"ryoeuyo/testgo/internal/domain/models"
	"ryoeuyo/testgo/internal/http/httputils"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	resp := models.Response{
		Data: map[string]interface{}{
			"ping": "pong xd",
		},
		StatusCode: http.StatusOK,
	}

	log.Printf("StatusCode: %d, Success", resp.StatusCode)
	httputils.WriteResponse(w, &resp, map[string]string{
		"Content-Type": "application/json",
	})
}

func PingV2(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	resp := models.Response{
		Data: map[string]interface{}{
			"ping": "pong pong ping pong pong xd xd",
		},
		StatusCode: http.StatusOK,
	}

	log.Printf("StatusCode: %d, Success", resp.StatusCode)
	httputils.WriteResponse(w, &resp, map[string]string{
		"Content-Type": "application/json",
	})
}
