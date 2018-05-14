package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yenchieh/APIinGo/controller"
)

func New() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(controller.Auth)
	r.GET("/stock", controller.StockAPI)

	r.GET("/stock/:symbol", controller.GetStockBySymbol)

	return r
}