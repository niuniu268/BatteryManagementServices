package cmd

import (
	`encoding/json`
	`fmt`
)

type Charging struct {
	Status string `json:"charging"`
}

func ChargingSwitch(key, value string) Charging {

	url := "http://127.0.0.1:5000/charge"

	body := PostHttp(url, key, value)

	var data Charging
	err := json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf(data.Status + "\n")

	return data
}

func (c *Charging) ShowCharging(battery BatteryData) {
	if c.Status == "off" {
		fmt.Printf("charging: " + c.Status + "\n")
		return
	}

	battery.percentage()

}
