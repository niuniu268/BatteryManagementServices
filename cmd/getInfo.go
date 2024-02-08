package cmd

import (
	`encoding/json`
	`fmt`
)

type BatteryData struct {
	SimTimeHour        int     `json:"sim_time_hour"`
	SimTimeMin         int     `json:"sim_time_min"`
	BaseCurrentLoad    float32 `json:"base_current_load"`
	BatteryCapacityKWh float32 `json:"battery_capacity_kWh"`
}

func GetInfo() BatteryData {

	url := "http://127.0.0.1:5000/info"

	body := GetHttp(url)
	var batterydata BatteryData

	err := json.Unmarshal(body, &batterydata)
	if err != nil {
		panic(err)
	}

	fmt.Println("Hour:", batterydata.SimTimeHour)
	fmt.Println("Minute:", batterydata.SimTimeMin)
	fmt.Println("Current Load:", batterydata.BaseCurrentLoad)
	fmt.Println("Battery Capacity (kWh):", batterydata.BatteryCapacityKWh)

	return batterydata
}

func (b *BatteryData) percentage() {

	fmt.Printf("now the battery condition: %v \n", b.BaseCurrentLoad/b.BatteryCapacityKWh*100)

}
