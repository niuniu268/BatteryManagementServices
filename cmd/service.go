package cmd

import (
	`fmt`
	`github.com/spf13/cobra`
)

var status bool

var showPriceCmd = &cobra.Command{
	Use:   "price",
	Short: "show SE 3 price",
	Long:  `show SE 3 price`,
	Run: func(cmd *cobra.Command, args []string) {
		price := GetPrice()

		price.ShowPrice()

	},
}

var showConsumptionCmd = &cobra.Command{
	Use:   "consumption",
	Short: "show consumption over 24 hours",
	Long:  `show consumption over 24 hours`,
	Run: func(cmd *cobra.Command, args []string) {
		consumption := GetConsumption()
		consumption.ShowConsumption()
	},
}

var chargeCmd = &cobra.Command{
	Use:   "charge",
	Short: "charging",
	Long:  `charging EV vehicle over 1 hour`,
	Run: func(cmd *cobra.Command, args []string) {

		info := GetInfo()

		var chargeStatus Charging

		if status {
			chargeStatus = ChargingSwitch("charging", "on")
			chargeStatus.ShowCharging(info)
		} else {
			chargeStatus = ChargingSwitch("charging", "off")

		}

	},
}

var showLowestConsumption = &cobra.Command{
	Use:   "showLowestConsumption",
	Short: "show the lowest consumption over 24 hours",
	Long:  `show the lowest household consumption over 24 hours`,
	Run: func(cmd *cobra.Command, args []string) {
		consumption := GetConsumption()
		index := consumption.ShowLowestConsumption()

		fmt.Printf("the lowest household consumption is %f, at clock %v \n", consumption.consumptionPerHour[index], index)
	},
}

var showLowestPrice = &cobra.Command{
	Use:   "showLowestPrice",
	Short: "show the lowest price over 24 hours",
	Long:  `show the lowest household cost over 24 hours`,
	Run: func(cmd *cobra.Command, args []string) {
		price := GetPrice()
		index := price.ShowLowestPrice()

		fmt.Printf("the lowest price  is %f, at clock %v \n", price.pricePerHour[index], index)
	},
}

var showTimeStamp = &cobra.Command{
	Use:   "showTimeStamp",
	Short: "show the timestamp over 24 hours",
	Long:  `show the timestamp over 24 hours`,
	Run: func(cmd *cobra.Command, args []string) {

		arr := make([]TimeStamp, 24)
		price := GetPrice()
		consumption := GetConsumption()

		arr = GetTimeStamp(consumption, price)

		for _, stamp := range arr {
			fmt.Printf("the clock %v, the value %v, the order %v, price %v, price order %v \n", stamp.time, stamp.ConsumptionPerHour, stamp.ConsumptionRecommend, stamp.PricePerHour, stamp.PriceRecommand)
		}

	},
}

func init() {
	rootCmd.AddCommand(showPriceCmd)
	rootCmd.AddCommand(showConsumptionCmd)
	rootCmd.AddCommand(chargeCmd)
	rootCmd.AddCommand(showLowestConsumption)
	rootCmd.AddCommand(showLowestPrice)
	rootCmd.AddCommand(showTimeStamp)

	chargeCmd.PersistentFlags().BoolVar(&status, "onoff", true, "turn on or off charging")

}
