package webservice

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithError(t *testing.T) {
	t.Run("Respond with error", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/create", nil)
		response := httptest.NewRecorder()
		_, err := RespondWithError(response, request, http.StatusBadRequest, "something happened")
		require.Nil(t, err)
		require.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func TestRespondWithJSON(t *testing.T) {
	t.Run("Respond with json", func(t *testing.T) {
		response := httptest.NewRecorder()
		expected := map[string]int{"ola": 2, "per": 5}
		RespondWithJSON(response, http.StatusOK, expected)
		require.Equal(t, http.StatusOK, response.Code)
		var actual map[string]int
		json.Unmarshal(response.Body.Bytes(), &actual)
		require.Equal(t, expected, actual)
	})
}
