package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	require.NoError(t, err)
	require.NotEmpty(t, req)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Ping)
	handler.ServeHTTP(rr, req)

	require.Equal(t, rr.Code, http.StatusOK)
}
