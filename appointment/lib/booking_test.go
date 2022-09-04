package lib

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestBooking(t *testing.T) {
	r := SetUpRouter()
	api := r.Group("data")
	{
		api.POST("/appointment", Booking)
	}
	request := RequestBooking{
		Name:           "Christy Schumm",
		Day:            "Wednesday",
		AvailableAt:    "7:00AM",
		AvailableUntil: "10:30AM",
	}
	jsonValue, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/data/appointment", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_saveData(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Should Success",
			wantErr: true,
		},
		{
			name:    "Failed",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := RequestBooking{
				Name:           "Christy Schumm",
				Day:            "Wednesday",
				AvailableAt:    "7:00AM",
				AvailableUntil: "10:30AM",
			}
			if err := saveData(request); (err != nil) != tt.wantErr {
				t.Errorf("saveData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_inTimeSpan(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "Should Success",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var start, end, at, until string
			start = "9:00AM"
			end = "5:30PM"
			at = "10:00AM"
			until = "11:00AM"

			if got := inTimeSpan(start, end, at, until); got != tt.want {
				t.Errorf("inTimeSpan() = %v, want %v", got, tt.want)
			}
		})
	}
}
