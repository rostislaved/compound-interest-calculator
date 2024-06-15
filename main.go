package main

import (
	"fmt"

	"github.com/rostislaved/compound-interest-calculator/config"
)

func main() {
	c := 1

	var baseConfig config.BaseConfig

	switch c {
	case 1:
		cfg := config.BaseConfig{
			StartAmount:     9000.,
			NumberOfPeriods: 10 * 365 * 6,
			Percent:         0.005,
			PercentEveryN:   1,
			Deposit:         0.,
			DepositEveryN:   1, // Каждые сколько периодов делается deposit
			ReinvestEveryN:  1,
		}

		baseConfig = cfg.GetBaseConfig()
	case 2:
		cfg := config.BaseConfig{
			StartAmount:     9000.,
			NumberOfPeriods: 10 * 365 * 6,
			Percent:         0.005,
			PercentEveryN:   1,
			Deposit:         0.,
			DepositEveryN:   1, // Каждые сколько периодов делается deposit
			ReinvestEveryN:  6 * 60,
		}

		baseConfig = cfg.GetBaseConfig()
	case 3:
		// Этот конфиг как на сайте
		cfg := config.CommonConfig{
			StartAmount:          9000.,
			DurationOfInvestment: 10 * config.Year,
			Percent:              12., // Годовых

			Deposit:       100,
			DepositEveryN: config.Month,

			ReinvestEveryN: 6 * config.Month,
		}

		baseConfig = cfg.GetBaseConfig()
	default:
		return
	}

	calculation := New(baseConfig)

	result := calculation.Calc()

	//result.PrintStatsByPeriod()
	fmt.Println()
	result.PrintStats()
}
