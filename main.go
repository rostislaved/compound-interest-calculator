package main

import (
	"fmt"

	"github.com/rostislaved/compound-interest-calculator/config"
)

func main() {

	// Этот конфиг как на сайте
	//cfg := config.CommonConfig{
	//	StartAmount:          100.,
	//	DurationOfInvestment: 10 * config.Month,
	//	Percent:              12., // Годовых
	//	ReinvestmentPeriods:  config.Month,
	//	DepositPeriods:       config.Month,
	//	Deposit:              10,
	//}

	//x := 1 * 6 * 365.
	//nop := int(float64(1*6*365) / x)
	//fmt.Println(float64(1*6*365) / x)
	//fmt.Println(nop)
	//p := 0.01 * x
	//fmt.Println(p)
	//cfg := config.BaseConfig{
	//	NumberOfPeriods: 10 * nop,
	//	StartAmount:     9000.,
	//	Percent:         p,
	//	Deposit:         0.,
	//	DepositEveryN:          1, // Каждые сколько периодов делается deposit
	//}
	//
	cfg := config.BaseConfig{
		NumberOfPeriods: 10 * 12,
		StartAmount:     9000.,
		Percent:         1,
		PercentEveryN:   1,
		Deposit:         100.,
		DepositEveryN:   1, // Каждые сколько периодов делается deposit
	}

	//cfg := config.BaseConfig{
	//	NumberOfPeriods: 120,
	//	StartAmount:     100.,
	//	Percent:         1,
	//	Deposit:         10.,
	//	DepositEveryN:          1, // Каждые сколько периодов делается deposit
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

	depositVector := generateTopUpVector(cfg.NumberOfPeriods, cfg.DepositEveryN)
	percentVector := generateTopUpVector(cfg.NumberOfPeriods, cfg.PercentEveryN)

	for i := range periodsVector {
		periodsVector[i] = Period{
			startAmount: 0,
			//percent:     cfg.Percent,
		}

		if depositVector[i] == 1 {
			periodsVector[i].deposit = cfg.Deposit
		}

		if percentVector[i] == 1 {
			periodsVector[i].percent = cfg.Percent
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

	for i := range c.periods {
		var depositSum float64
		var percentSum float64

		if i == 0 {
			depositSum = c.startAmount
			percentSum = 0
		} else {
			depositSum = c.periods[i-1].depositSum
			percentSum = c.periods[i-1].percentSum
		}

		c.periods[i].depositSum = depositSum + c.periods[i].deposit
		c.periods[i].percentSum = percentSum + c.periods[i].increaseByPercent
	}

	return Result{c.periods}
}
