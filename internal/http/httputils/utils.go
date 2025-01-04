package httputils

import (
	"encoding/json"
	"net/http"
	"ryoeuyo/testgo/internal/domain/models"
)

func WriteResponse(w http.ResponseWriter, resp *models.Response, headers ...map[string]string) (int, error) {
	w.WriteHeader(resp.StatusCode)

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header().Add(k, v)
		}
	}

	r, err := json.Marshal(resp)
	if err != nil {
		return 0, err
	}

	return w.Write(r)
}
