package lib

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadData(t *testing.T) {
	r := SetUpRouter()
	api := r.Group("data")
	{
		api.GET("/all", Booking)
	}
	request := RequestUserApproved{
		ID: 25,
		Approved: true,
	}
	jsonValue, _ := json.Marshal(request)
	req, _ := http.NewRequest("GET", "/data/all", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
