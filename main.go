package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gosuri/uilive"
)

type gasPrices interface {
	Slow() float64
	Medium() float64
	Fast() float64
	Prices()
}

const ethGasStationURL = "https://ethgasstation.info/json/ethgasAPI.json"

type ethGasStation struct {
	SafeLowPrice float64 `json:"safeLow"`
	AveragePrice float64 `json:"average"`
	FastPrice    float64 `json:"fast"`
}

func (e *ethGasStation) Slow() float64 {
	return e.SafeLowPrice / 10
}

func (e *ethGasStation) Medium() float64 {
	return e.AveragePrice / 10
}

func (e *ethGasStation) Fast() float64 {
	return e.FastPrice / 10
}

func (e *ethGasStation) Prices() {
	resp, err := http.Get(ethGasStationURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, e)
	return
}

func printGasPrices(w io.Writer, gp gasPrices) {
	fmt.Fprintf(w, "slow   : %5.f\n", gp.Slow())
	fmt.Fprintf(w, "medium : %5.f\n", gp.Medium())
	fmt.Fprintf(w, "fast   : %5.f\n", gp.Fast())
}

func main() {
	slow := flag.Bool("slow", false, "")
	medium := flag.Bool("medium", false, "")
	fast := flag.Bool("fast", false, "")
	provider := flag.String("provider", "ethGasStation", "")
	interval := flag.Int("interval", 1, "")
	flag.Parse()

	var prices gasPrices
	switch *provider {
	case "ethGasStation":
		prices = &ethGasStation{}
	}
	prices.Prices()

	switch {
	case *slow:
		fmt.Println(prices.Slow())
		return
	case *medium:
		fmt.Println(prices.Medium())
		return
	case *fast:
		fmt.Println(prices.Fast())
		return
	}

	writer := uilive.New()
	writer.Start()
	defer writer.Stop()
	intervalDuration := time.Duration(*interval)
	for {
		printGasPrices(writer, prices)
		<-time.After(intervalDuration * time.Second)
		prices.Prices()
	}
}
