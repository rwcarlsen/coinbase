package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	log.SetFlags(0)

	url := "https://coinbase.com/api/v1/currencies/exchange_rates"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	rates := map[string]string{}
	if err := json.Unmarshal(data, &rates); err != nil {
		log.Fatal(err)
	}

	usdRate, err := strconv.ParseFloat(rates["btc_to_usd"], 64)

	fmt.Printf("1 BTC = $%.2f\n", usdRate)
}
