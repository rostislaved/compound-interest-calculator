package main

import (
	"fmt"

	"github.com/rostislaved/compound-interest-calculator/config"
)

func main() {
	fmt.Println(generateVector(5, 3, 9))
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
		ReinvestEveryN:  1,
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

	depositVector := generateVector(cfg.NumberOfPeriods, cfg.DepositEveryN, cfg.Deposit)
	percentVector := generateVector(cfg.NumberOfPeriods, cfg.PercentEveryN, cfg.Percent)
	reinvestVector := generateVector(cfg.NumberOfPeriods, cfg.ReinvestEveryN, true)

	for i := range periodsVector {
		periodsVector[i].deposit = depositVector[i]
		periodsVector[i].percent = percentVector[i]
		periodsVector[i].reinvestInThisPeriod = reinvestVector[i]
	}

	return Calculation{
		startAmount: cfg.StartAmount,
		periods:     periodsVector,
	}
}

// Генерирует вектор длиной n, в котором значение value каждые everyN элементов. Начинает с 1го.
// generateVector(5, 3, 9) = [9 0 0 9 0]
func generateVector[T any](n, everyN int, value T) []T {
	res := make([]T, n)

	for i := 0; i < n; i++ {
		if i%everyN == 0 {
			res[i] = value
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

type Period struct {
	startAmount       float64
	increaseByPercent float64
	percent           float64
	deposit           float64

	depositSum float64
	percentSum float64

	notYetReinvestedAmount float64
	reinvestInThisPeriod   bool
}

func (p *Period) EndAmount() float64 {
	return p.startAmount + p.increaseByPercent + p.deposit
}

func (p *Period) calculatePeriod(previousPeriodEndAmount float64) {
	p.startAmount = previousPeriodEndAmount

	p.increaseByPercent = p.startAmount * (p.percent / 100)
}
