package provider

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Provider interface {
	Slow() float64
	Medium() float64
	Fast() float64
	Prices() error
}

func callAPI(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func Print(w io.Writer, p Provider) {
	fmt.Fprintf(w, "slow   : %5.f\n", p.Slow())
	fmt.Fprintf(w, "medium : %5.f\n", p.Medium())
	fmt.Fprintf(w, "fast   : %5.f\n", p.Fast())
}
