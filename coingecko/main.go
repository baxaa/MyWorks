package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

const (
	apiURL   = "https://api.coingecko.com/api/v3/coins/markets"
	interval = 10 * time.Minute
)

type CoinMarketData struct {
	ID           string  `json:"id"`
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	CurrentPrice float64 `json:"current_price"`
}

func GetMarketData() ([]CoinMarketData, error) {
	url := fmt.Sprintf("%s?vs_currency=usd&order=market_cap_desc&per_page=250&page=1", apiURL)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("failed to close %v", err)
		}
	}(response.Body)

	var marketData []CoinMarketData
	err = json.NewDecoder(response.Body).Decode(&marketData)
	if err != nil {
		return nil, err
	}

	return marketData, nil
}

func GetCurrencyPrice(symbol string) (float64, error) {
	marketData, err := GetMarketData()
	if err != nil {
		return 0, err
	}

	for _, data := range marketData {
		if data.Symbol == symbol {
			return data.CurrentPrice, nil
		}
	}

	return 0, fmt.Errorf("krypto with value %s not found", symbol)
}

func GetAllCurrencies(c *gin.Context) {
	marketData, err := GetMarketData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in getting data"})
		return
	}

	var response []gin.H
	for _, data := range marketData {
		response = append(response, gin.H{
			"name":         data.Name,
			"symbol":       data.Symbol,
			"currentPrice": data.CurrentPrice,
		})
	}

	c.JSON(http.StatusOK, response)
}

func GetCurrencyBySymbol(c *gin.Context) {
	symbol := c.Param("symbol")
	price, err := GetCurrencyPrice(symbol)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("rate for %s not found", symbol)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"symbol": symbol, "currentPrice": price})
}

func main() {
	r := gin.Default()

	r.GET("/crypto", GetAllCurrencies)
	r.GET("/crypto/:symbol", GetCurrencyBySymbol)

	go func() {
		err := r.Run(":8080")
		if err != nil {
			fmt.Println("error to start server:", err)
		}
	}()

	fmt.Println("server started")

	for {
		time.Sleep(interval)
	}
}
