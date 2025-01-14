package main

import (
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/goccy/go-json"
)

const gasNowUrl string = "https://beaconcha.in/api/v1/execution/gasnow"

const (
	WeiPerEth  float64 = 1e18
	WeiPerGwei float64 = 1e9
)

func WeiToGwei(wei *big.Int) float64 {
	var weiFloat big.Float
	var gwei big.Float
	weiFloat.SetInt(wei)
	gwei.Quo(&weiFloat, big.NewFloat(WeiPerGwei))
	gwei64, _ := gwei.Float64()
	return gwei64
}

// Standard response
type gasNowResponse struct {
	Data struct {
		Rapid    uint64  `json:"rapid"`
		Fast     uint64  `json:"fast"`
		Standard uint64  `json:"standard"`
		Slow     uint64  `json:"slow"`
		PriceUSD float64 `json:"priceUSD"`
	} `json:"data"`
}

type GasFeeSuggestion struct {
	RapidWei  float64
	RapidTime string

	FastWei  float64
	FastTime string

	StandardWei  float64
	StandardTime string

	SlowWei  float64
	SlowTime string

	EthUsd float64
}

func (gfs GasFeeSuggestion) String() string {
	return fmt.Sprintf("%s:\t%.1f\n%s:\t%.1f\nETH/USD:\t%.2f\n",
		gfs.RapidTime, gfs.RapidWei,
		gfs.FastTime, gfs.FastWei,
		gfs.EthUsd,
	)
}

// Get gas prices
func GetGasPrices() (GasFeeSuggestion, error) {

	// Send request
	response, err := http.Get(gasNowUrl)
	if err != nil {
		return GasFeeSuggestion{}, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	// Check the response code
	if response.StatusCode != http.StatusOK {
		return GasFeeSuggestion{}, fmt.Errorf("request failed with code %d", response.StatusCode)
	}

	// Get response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return GasFeeSuggestion{}, err
	}

	// Deserialize response
	var gnResponse gasNowResponse
	if err := json.Unmarshal(body, &gnResponse); err != nil {
		return GasFeeSuggestion{}, fmt.Errorf("Could not decode Gas Now response: %w", err)
	}

	suggestion := GasFeeSuggestion{
		RapidWei:  WeiToGwei(big.NewInt(0).SetUint64(gnResponse.Data.Rapid)),
		RapidTime: "15 Seconds",

		FastWei:  WeiToGwei(big.NewInt(0).SetUint64(gnResponse.Data.Fast)),
		FastTime: "1 Minute",

		StandardWei:  WeiToGwei(big.NewInt(0).SetUint64(gnResponse.Data.Standard)),
		StandardTime: "3 Minutes",

		SlowWei:  WeiToGwei(big.NewInt(0).SetUint64(gnResponse.Data.Slow)),
		SlowTime: ">10 Minutes",

		EthUsd: gnResponse.Data.PriceUSD,
	}

	// Return
	return suggestion, nil

}
