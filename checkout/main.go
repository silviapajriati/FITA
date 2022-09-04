package main

import (
	"github.com/gin-gonic/gin"
	"checkout/lib"
)

func main(){
	r := gin.Default()
    api := r.Group("product")
	{
		 api.GET("/data", lib.ListProduct)
		 api.POST("/checkout", lib.Checkout)
	}

	r.Run(":8090")
}

