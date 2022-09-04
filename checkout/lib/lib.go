package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func listProduct() []Item {
	item := []Item{
		{
			SKU:   		"120P90",
			Name:  		"Google Pro",
			Price: 		49.99,
			Quantity:   10,
		},
		{
			SKU:   "43N23P",
			Name:  "Mackbook Pro",
			Price: 5399.99,
			Quantity:   5,
		},
		{
			SKU:   		"A304SD",
			Name:  		"Alexa Speaker",
			Price: 		109.50,
			Quantity:   10,
		},
		{
			SKU:   		"234234",
			Name:  		"Raspberry Pi B",
			Price: 		30.00,
			Quantity:   2,
		},
	}
	listItem := ListItem{}
	listItem.Items = item

	return listItem.Items
}

func ListProduct(c *gin.Context){
		
	arr := listProduct()

	c.JSON(200, arr)
}

func Checkout(c *gin.Context){
	var ReqStruct RequestCheckout
	var Response ResponseCheckout
	var arr_item []Item
	var item Item
	total := 0.0

	err := c.BindJSON(&ReqStruct)
	if err != nil {
		c.JSON(200, gin.H{"errcode": http.StatusBadRequest, "description": "Post Data Err"})
		return
	}

	//get all product
	arr := listProduct()
	for _, v1 := range arr {
		for _, v2 := range ReqStruct{
			if v1.SKU == v2.SKU{
				if v1.SKU == "120P90"{
					qty := v2.Quantity / 3
					item = Item{
						SKU: v1.SKU,
						Quantity: v2.Quantity-qty,
						Price: v1.Price,
						Name: v1.Name,
					}
				} else if v1.SKU == "A304SD"{
					price := v1.Price
					if v2.Quantity >= 3{
						price = price * 0.9
					}
					item = Item{
						SKU: v1.SKU,
						Quantity: v2.Quantity,
						Price: price,
						Name: v1.Name,
					}
				} else{
					item = Item{
						SKU: v1.SKU,
						Quantity: v2.Quantity,
						Price: v1.Price,
						Name: v1.Name,
					}
				}

				total += float64(item.Quantity) * float64(item.Price)
				arr_item = append(arr_item, item)

				if v1.SKU == "43N23P" {
					item = getOne("234234", arr)
					item.Quantity = v2.Quantity 
					arr_item = append(arr_item, item)
				}	
			}
		}
	}	

	Response.Status = true
	Response.TotalPrice = float32(total)
	Response.Data = arr_item

	c.JSON(200, Response)
	
}

func getOne(sku string, list []Item) Item{
	for _, v := range list{
		if v.SKU == sku {
			return v
		}
	}
	return Item{}
}