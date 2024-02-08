package cmd

import (
	`encoding/json`
	`fmt`
)

type Price struct {
	pricePerHour []float32
}

func GetPrice() Price {
	url := "http://127.0.0.1:5000/priceperhour"

	body := GetHttp(url)

	var price Price

	err := json.Unmarshal(body, &price.pricePerHour)
	if err != nil {
		panic(err)
	}

	return price
}

func (p *Price) ShowPrice() {

	for index, elem := range p.pricePerHour {
		elemS := fmt.Sprintf("%f", elem)
		fmt.Printf("clock %v hour, price is %s ore per kwh \n", index, elemS)
	}

}

func (p *Price) ShowLowestPrice() int {

	var m float32 = 1000
	index := 0

	for i, f := range p.pricePerHour {
		if m > f {
			m = f
			index = i
		}
	}

	return index
}
