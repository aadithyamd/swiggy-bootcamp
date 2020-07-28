package main

import "fmt"

type Trade struct {
	CompName string
	Volume int
	Price float64
	Buy bool
}

func NewTrade(companyName string, volume int, price float64, buy bool) (*Trade, error){
	if companyName == "" {
		return nil, fmt.Errorf("Sorry")
	}
	if (price < 0){
		return nil, fmt.Errorf("Sorry")
	}
	if (volume <= 0){
		return nil, fmt.Errorf("Sorry")
	}

	trade := &Trade{
		CompName: companyName,
		Price: price,
		Volume: volume,
		Buy: buy,
	}

	return trade, nil
}

func main() {

}
