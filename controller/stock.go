package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"strings"
	"fmt"
)

type Stock struct {
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
	Open   float32 `json:"open"`
	Close  float32 `json:"close"`
}

var savedStocks []Stock

func init() {
	stocks := []Stock{
		{Name: "Apple Inc", Symbol: "AAPL", Open: 189.49, Close: 188.59},
		{Name: "Alphabet Inc", Symbol: "GOOGL", Open: 1100.41, Close: 1103.38},
	}
	savedStocks = append(savedStocks, stocks...)
}

func AllStocks(c *gin.Context) {
	c.JSON(http.StatusOK, savedStocks)
}

func GetStockBySymbol(c *gin.Context) {
	symbol := c.Param("symbol")
	if symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Symbol is required",
		})
		return
	}
	var foundStock *Stock
	for _, stock := range savedStocks {
		if strings.ToLower(stock.Symbol) == strings.ToLower(symbol) {
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

func AddStockData(c *gin.Context) {
	var newStockData Stock
	if err := c.BindJSON(&newStockData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Please enter valid data",
		})
		return
	}
	savedStocks = append(savedStocks, newStockData)

	c.JSON(http.StatusOK, savedStocks)
}

func DeleteStockDataBySymbol(c *gin.Context) {
	symbol := c.Param("symbol")
	if symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Symbol is required",
		})
		return
	}
	found := false
	for index, stock := range savedStocks {
		if strings.ToLower(stock.Symbol) == strings.ToLower(symbol) {
			savedStocks = append(savedStocks[:index], savedStocks[index+1:]...)
			found = true
			break
		}
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Symbol %s not exist", symbol),
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}