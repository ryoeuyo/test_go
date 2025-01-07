package rest

import (
	"log"
	"net/http"
	"ryoeuyo/testgo/internal/domain/models"
	"ryoeuyo/testgo/internal/http/httputils"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{
		Data: map[string]interface{}{
			"ping": "pong xdrweqreqwrqwe",
		},
		StatusCode: http.StatusOK,
	}

	if b, err := httputils.WriteResponse(w, &resp, map[string]string{
		"Content-Type": "application/json",
	}); err != nil {
		log.Printf("failed write response: %s", err.Error())
	} else {
		log.Printf("successfully writed %d bytes", b)
	}
}

func PingV2(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{
		Data: map[string]interface{}{
			"ping": "pong pong ping pong pong xd xd",
		},
		StatusCode: http.StatusOK,
	}

	if b, err := httputils.WriteResponse(w, &resp, map[string]string{
		"Content-Type": "application/json",
	}); err != nil {
		log.Printf("failed write response: %s", err.Error())
	} else {
		log.Printf("successfully writed %d bytes", b)
	}
}
