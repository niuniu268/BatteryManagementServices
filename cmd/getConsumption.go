package cmd

import (
	`encoding/json`
	`fmt`
)

type Consumption struct {
	consumptionPerHour []float32
}

func GetConsumption() Consumption {
	url := "http://127.0.0.1:5000/baseload"

	body := GetHttp(url)

	var consumption Consumption

	err := json.Unmarshal(body, &consumption.consumptionPerHour)
	if err != nil {
		panic(err)
	}

	return consumption
}

func (c *Consumption) ShowConsumption() {

	for index, elem := range c.consumptionPerHour {
		elemS := fmt.Sprintf("%f", elem)
		fmt.Printf("clock %v hour, consumption is %s kwh \n", index, elemS)
	}

}

func (c *Consumption) ShowLowestConsumption() int {

	var m float32 = 1000
	index := 0

	for i, f := range c.consumptionPerHour {
		if m > f {
			m = f
			index = i
		}
	}

	return index
}
