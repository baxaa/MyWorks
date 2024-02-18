# Crypto Price API
## About Project
 This is a simple API written in Go by using Gin web framework. All of the data is fetched in CoinGecko API. 
## Endpoints
### Get all cryptocurrency prices
* Endpoint: /crypto <br>
* Method: GET <br>
* Description: Returns a list of all available cryptocurrencies along with their symbols and current prices in USD. <br>

### Get Cryptocurrency by Symbol
* Endpoint: /crypto/:symbol <br>
* Method: GET <br>
* Parameters: <br>
  * symbol: The symbol of the cryptocurrency. <br> 
* Description: Returns the current price of the specified cryptocurrency in USD. <br>

## Usage
### 1.Start the server
  ```
  go run main.go
  ```
### 2.Access the API 
  Example: Get all cryptocurrencies
  ```
  http://localhost:8080/crypto
  ```
  Example: Get cryptocurrency price by symbol (replace <symbol> with the desired cryptocurrency symbol)
  ```
  http://localhost:8080/crypto/<symbol>
  ```
## Dependencies
* Gin:  A web framework for Go.
