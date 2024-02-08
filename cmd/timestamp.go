package cmd

import (
	`sort`
)

type TimeStamp struct {
	time                 int     `json:"time,omitempty"`
	ConsumptionPerHour   float32 `json:"consumptionPerHour,omitempty"`
	ConsumptionRecommend int     `json:"consumptionRecommand,omitempty"`
	PricePerHour         float32 `json:"pricePerHour,omitempty"`
	PriceRecommand       int     `json:"priceRecommand,omitempty"`
}

func GetTimeStamp(consumption Consumption, price Price) []TimeStamp {

	arr := make([]TimeStamp, 0, 24)

	var orderPerhourC = make([]float32, 24)

	for i := range orderPerhourC {
		orderPerhourC[i] = consumption.consumptionPerHour[i]
	}

	sort.Slice(orderPerhourC, func(i, j int) bool {
		return orderPerhourC[i] < orderPerhourC[j]
	})

	var orderPerhourP = make([]float32, 24)

	for i := range orderPerhourP {
		orderPerhourP[i] = price.pricePerHour[i]
	}

	sort.Slice(orderPerhourP, func(i, j int) bool {
		return orderPerhourP[i] < orderPerhourP[j]
	})

	var findindex func(arr []float32, value float32) int

	findindex = func(arr []float32, value float32) int {
		for i, v := range arr {
			if v == value {
				return i
			}
		}
		return -1
	}

	for i := 0; i < len(consumption.consumptionPerHour); i++ {
		var timestamp TimeStamp
		timestamp.time = i
		timestamp.ConsumptionPerHour = consumption.consumptionPerHour[i]
		timestamp.ConsumptionRecommend = findindex(orderPerhourC, consumption.consumptionPerHour[i])
		timestamp.PricePerHour = price.pricePerHour[i]
		timestamp.PriceRecommand = findindex(orderPerhourP, price.pricePerHour[i])

		arr = append(arr, timestamp)
	}

	return arr
}
