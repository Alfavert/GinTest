package main

import (
	"GinServer/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	PointG := gin.Default()
	PointG.UseH2C = true
	internal.Config.ParseConfigs()
	internal.Init()
	internal.InitCart()
	PointG.GET("/api/get_info", internal.Dlist.GetInfoHandler)
	PointG.GET("/api/get_all_info", internal.Dlist.GetAllInfoHandler)
	PointG.POST("/api/add_info", internal.Dlist.AddInfoHandler)
	PointG.POST("/api/add_cart_info", internal.Clist.AddInfoCart)
	PointG.GET("/api/cart", internal.Clist.Compare)
	PointG.Run()
}
