package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"

)

type Stock struct {
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
	Open   float32 `json:"open"`
	Close  float32 `json:"close"`
}
type Stocks []Stock

func StockAPI(c *gin.Context) {
	c.JSON(http.StatusOK, getStocksData())
}

func GetStockBySymbol(c *gin.Context) {
	symbol := c.Param("symbol")
	if symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Symbol is required",
		})
		return
	}
	stocks := getStocksData()
	var foundStock *Stock
	for _, stock := range stocks {
		if stock.Symbol == symbol {
			foundStock = &stock
			break
		}
	}
	if foundStock == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, foundStock)
}

func getStocksData() Stocks {
	return Stocks{
		Stock{Name: "Apple Inc", Symbol: "AAPL", Open: 189.49, Close: 188.59},
		Stock{Name: "Alphabet Inc", Symbol: "GOOGL", Open: 1100.41, Close: 1103.38},
	}
}
