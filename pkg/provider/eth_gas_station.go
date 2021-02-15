package provider

import "encoding/json"

const ethGasStationURL = "https://ethgasstation.info/json/ethgasAPI.json"

type EthGasStation struct {
	SafeLowPrice float64 `json:"safeLow"`
	AveragePrice float64 `json:"average"`
	FastPrice    float64 `json:"fast"`
}

func (p *EthGasStation) Slow() float64 {
	return p.SafeLowPrice / 10
}

func (p *EthGasStation) Medium() float64 {
	return p.AveragePrice / 10
}

func (p *EthGasStation) Fast() float64 {
	return p.FastPrice / 10
}

func (p *EthGasStation) Prices() error {
	body, err := callAPI(ethGasStationURL)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, p)
}
