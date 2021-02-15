package provider

import "encoding/json"

const gasWatchURL = "https://ethgas.watch/api/gas"

type GasWatch struct {
	SlowPrice struct {
		Gwei float64 `json:"gwei"`
	} `json:"slow"`
	NormalPrice struct {
		Gwei float64 `json:"gwei"`
	} `json:"normal"`
	FastPrice struct {
		Gwei float64 `json:"gwei"`
	} `json:"fast"`
}

func (p *GasWatch) Slow() float64 {
	return p.SlowPrice.Gwei
}

func (p *GasWatch) Medium() float64 {
	return p.NormalPrice.Gwei
}

func (p *GasWatch) Fast() float64 {
	return p.FastPrice.Gwei
}

func (p *GasWatch) Prices() error {
	body, err := callAPI(gasWatchURL)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, p)
}
