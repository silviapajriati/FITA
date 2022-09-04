package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func Test_listProduct(t *testing.T) {
	tests := []struct {
		name string
		want []Item
	}{
		{
			name: "Should Success",
			want: []Item{
				{
					SKU:      "120P90",
					Name:     "Google Pro",
					Price:    49.99,
					Quantity: 10,
				},
				{
					SKU:      "43N23P",
					Name:     "Mackbook Pro",
					Price:    5399.99,
					Quantity: 5,
				},
				{
					SKU:      "A304SD",
					Name:     "Alexa Speaker",
					Price:    109.50,
					Quantity: 10,
				},
				{
					SKU:      "234234",
					Name:     "Raspberry Pi B",
					Price:    30.00,
					Quantity: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListProduct(t *testing.T) {
	mockResponse := `[{"sku":"120P90","name":"Google Pro","price":49.99,"quantity":10},{"sku":"43N23P","name":"Mackbook Pro","price":5399.99,"quantity":5},{"sku":"A304SD","name":"Alexa Speaker","price":109.5,"quantity":10},{"sku":"234234","name":"Raspberry Pi B","price":30,"quantity":2}]`
	r := SetUpRouter()
	api := r.Group("product")
	{
		api.GET("/data", ListProduct)
	}

	req, _ := http.NewRequest("GET", "/product/data", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	fmt.Println(responseData)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, 200)
}

func TestCheckout(t *testing.T) {
	r := SetUpRouter()
    api := r.Group("product")
	{
		api.POST("/checkout", Checkout)
	}
    request := RequestCheckout{
        struct{SKU string "json:\"sku\""; Quantity int "json:\"quantity\""}{
			SKU: "A304SD",
			Quantity: 3,
		},
    }
    jsonValue, _ := json.Marshal(request)
    req, _ := http.NewRequest("POST", "/product/checkout", bytes.NewBuffer(jsonValue))

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
}
