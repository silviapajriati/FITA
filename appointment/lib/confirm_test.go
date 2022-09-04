package lib

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConform(t *testing.T) {
	r := SetUpRouter()
	api := r.Group("data")
	{
		api.POST("/confirmation", Booking)
	}
	request := RequestApprove{
		Approve:    true,
		Reschedule: true,
		ID:         25,
		NewSchedule: RequestBooking{
			Name:           "Christy Schumm",
			Day:            "Wednesday",
			AvailableAt:    "7:00AM",
			AvailableUntil: "10:30AM",
		},
	}
	jsonValue, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/data/confirmation", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserConform(t *testing.T) {
	r := SetUpRouter()
	api := r.Group("data")
	{
		api.POST("/user-confirmation", Booking)
	}
	request := RequestUserApproved{
		ID: 25,
		Approved: true,
	}
	jsonValue, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/data/user-confirmation", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
