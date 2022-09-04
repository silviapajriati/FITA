package main

import (
	"github.com/gin-gonic/gin"
	"appointment/lib"
)

func main(){
	r := gin.Default()
    api := r.Group("data")
	{
		 api.GET("/all", lib.ReadData)
		 api.POST("/appointment", lib.Booking)
		 api.POST("/confirmation", lib.Conform)
		 api.POST("/user-confirmation", lib.UserConform)
	}

	r.Run(":8090")
}