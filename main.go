package main

import (
	"fmt"

	"github.com/rostislaved/compound-interest-calculator/config"
)

func main() {
	cfg := config.CommonConfig{
		StartAmount:          100.,
		DurationOfInvestment: 10 * config.Yearly,
		Percent:              12., // Годовых
		ReinvestmentPeriods:  config.Monthly,
		DepositPeriods:       config.Monthly,
		Deposit:              10,
	}

	//cfg := config.BaseConfig{
	//	NumberOfPeriods: 60,
	//	StartAmount:     9000.,
	//	Percent:         0.005,
	//	Deposit:         100.,
	//	EveryN:          1, // Каждые сколько периодов делается deposit
	//}

	//cfg := config.BaseConfig{
	//	NumberOfPeriods: 120,
	//	StartAmount:     100.,
	//	Percent:         1,
	//	Deposit:         10.,
	//	EveryN:          1, // Каждые сколько периодов делается deposit
	//}

	baseConfig := cfg.GetBaseConfig()

	periods := New(baseConfig)

	result := periods.Calc()
	result.PrintStatsByPeriod()
	fmt.Println()
	result.PrintStats()
}

type Calculation struct {
	startAmount float64
	periods     []Period
}

func New(cfg config.BaseConfig) Calculation {
	periodsVector := make([]Period, cfg.NumberOfPeriods)

	topUpVector := generateTopUpVector(cfg.NumberOfPeriods, cfg.EveryN)

	for i := range periodsVector {
		periodsVector[i] = Period{
			startAmount: 0,
			percent:     cfg.Percent,
		}

		if topUpVector[i] == 1 {
			periodsVector[i].deposit = cfg.Deposit
		}
	}

	return Calculation{
		startAmount: cfg.StartAmount,
		periods:     periodsVector,
	}
}

// generateTopUpVector создает вектор пополнений.
// numberOfPeriods = 6
// everyN = 3
// [1, 0, 0, 1, 0, 0]. То есть в 1й и 4й период будет пополнение.
func generateTopUpVector(numberOfPeriods, everyN int) []int {
	res := make([]int, numberOfPeriods)

	for i := 0; i < numberOfPeriods; i++ {
		if i%everyN == 0 {
			res[i] = 1
		}
	}

	return res
}

func (c Calculation) Calc() Result {
	for i := range c.periods {
		var previousPeriodEndAmount float64

		if i == 0 {
			previousPeriodEndAmount = c.startAmount
		} else {
			previousPeriodEndAmount = c.periods[i-1].EndAmount()
		}

		c.periods[i].calculatePeriod(previousPeriodEndAmount)
	}

	return Result{c.periods}
}
