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
			NumberOfPeriods: 10 * 12,
			Percent:         1,
			PercentEveryN:   1,
			Deposit:         100.,
			DepositEveryN:   1, // Каждые сколько периодов делается deposit
			ReinvestEveryN:  12,
		}

		baseConfig = cfg.GetBaseConfig()
	case 2:
		// Этот конфиг как на сайте
		cfg := config.CommonConfig{
			StartAmount:          9000.,
			DurationOfInvestment: 10 * config.Year,
			Percent:              12., // Годовых
			Deposit:              100,
			DepositEveryN:        config.Month,
			ReinvestEveryN:       config.Month,
		}

		baseConfig = cfg.GetBaseConfig()
	default:
		return
	}

	calculation := New(baseConfig)

	result := calculation.Calc()

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

// Генерирует вектор длиной n, в котором каждый everyN'й элемент имеет значение value.
// generateVector(6, 3, 9) = [0 0 9 0 0 9].
func generateVector[T any](n, everyN int, value T) []T {
	res := make([]T, n)

	for i := 1; i <= n; i++ {
		if i%everyN == 0 {
			res[i-1] = value
		}
	}

	return res
}

func (c Calculation) Calc() Result {
	for i := range c.periods {
		var previousPeriod Period

		if i == 0 {
			previousPeriod = Period{
				startAmount: c.startAmount,
				endAmount:   c.startAmount,
			}
		} else {
			previousPeriod = c.periods[i-1]
		}

		c.periods[i].calculatePeriod(previousPeriod)
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
	endAmount         float64
	increaseByPercent float64
	percent           float64
	deposit           float64

	depositSum float64
	percentSum float64

	notYetReinvestedAmount float64
	reinvestInThisPeriod   bool
}

func (p *Period) EndAmount() float64 {
	return p.endAmount
}

func (p *Period) calculatePeriod(prev Period) {
	p.startAmount = prev.endAmount
	p.notYetReinvestedAmount += prev.notYetReinvestedAmount

	p.increaseByPercent = p.startAmount * (p.percent / 100)

	p.notYetReinvestedAmount += p.increaseByPercent

	totalIncrease := p.deposit

	if p.reinvestInThisPeriod {
		totalIncrease += p.notYetReinvestedAmount
		p.notYetReinvestedAmount = 0
	}

	p.endAmount = p.startAmount + totalIncrease
}
