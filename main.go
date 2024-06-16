package main

import (
	"fmt"

	"github.com/rostislaved/compound-interest-calculator/config"
)

func main() {
	c := 3

	var baseConfig1 config.BaseConfig

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

		baseConfig1 = cfg.GetBaseConfig()
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

		baseConfig1 = cfg.GetBaseConfig()
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

		baseConfig1 = cfg.GetBaseConfig()
	default:
		return
	}

	//baseConfig1 := config.BaseConfig{
	//	StartAmount:     9000.,
	//	NumberOfPeriods: 20 * 365 * 6,
	//	Percent:         0.005,
	//	PercentEveryN:   1,
	//	Deposit:         0.,
	//	DepositEveryN:   1, // Каждые сколько периодов делается deposit
	//	ReinvestEveryN:  1,
	//}
	//
	//baseConfig2 := config.BaseConfig{
	//	StartAmount:     9000.,
	//	NumberOfPeriods: 20 * 365 * 6,
	//	Percent:         0.005,
	//	PercentEveryN:   1,
	//	Deposit:         0.,
	//	DepositEveryN:   1, // Каждые сколько периодов делается deposit
	//	ReinvestEveryN:  6 * 365,
	//}

	calculation1 := New(baseConfig1)
	//calculation2 := New(baseConfig2)

	result1 := calculation1.Calc()
	//result2 := calculation2.Calc()

	//PrintStatsByPeriod(result)
	fmt.Println()
	PrintStatsByPeriod(result1)
	PrintStats(result1)
}
