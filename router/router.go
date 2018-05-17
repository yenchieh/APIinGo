package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yenchieh/APIinGo/controller"
)

func New() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(controller.Auth)

	api := r.Group("/api", controller.Auth)
	api.GET("/stock", controller.StockAPI)
	api.GET("/stock/:symbol", controller.GetStockBySymbol)
	api.PUT("/stock", controller.AddStockData)
	api.DELETE("/stock/:symbol", controller.DeleteStockDataBySymbol)

	return r
}