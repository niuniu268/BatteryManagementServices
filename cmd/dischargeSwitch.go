package cmd

import (
	`encoding/json`
	`fmt`
)

type Discharging struct {
	Status string `json:"discharging"`
}

func DischargingSwitch(key, value string) Discharging {

	url := "http://127.0.0.1:5000/discharge"

	body := PostHttp(url, key, value)

	var data Discharging
	err := json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	fmt.Printf(data.Status + "\n")

	return data
}
