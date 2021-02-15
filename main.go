package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gosuri/uilive"
	"github.com/harryscholes/gp/pkg/provider"
)

func main() {
	slow := flag.Bool("slow", false, "")
	medium := flag.Bool("medium", false, "")
	fast := flag.Bool("fast", false, "")
	providerName := flag.String("provider", "EthGasStation", "")
	interval := flag.Int("interval", 1, "")
	flag.Parse()

	var p provider.Provider
	switch *providerName {
	case "EthGasStation":
		p = &provider.EthGasStation{}
	case "GasWatch":
		p = &provider.GasWatch{}
	default:
		panic("Invalid gas price provider")
	}

	err := p.Prices()
	if err != nil {
		panic("Could not queury %s API")
	}

	switch {
	case *slow:
		fmt.Println(p.Slow())
		return
	case *medium:
		fmt.Println(p.Medium())
		return
	case *fast:
		fmt.Println(p.Fast())
		return
	}

	writer := uilive.New()
	writer.Start()
	defer writer.Stop()
	intervalDuration := time.Duration(*interval)
	for {
		provider.Print(writer, p)
		<-time.After(intervalDuration * time.Second)
		p.Prices()
	}
}
