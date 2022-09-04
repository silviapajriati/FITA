package lib

type Item struct {
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Price float32 `json:"price"`
	Quantity   int `json:"quantity"`
}

type ListItem struct {
	Items []Item
}

type RequestCheckout []struct {
	SKU   string `json:"sku"`
	Quantity   int `json:"quantity"`
}

type ResponseCheckout struct {
	Status 			bool 	`json:"status"`
	TotalPrice   	float32 `json:"totalPrice"`
	Data 			[]Item
}