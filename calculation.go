package main

import "github.com/rostislaved/compound-interest-calculator/config"

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

		// Add not yet reinvested sums
		if i == len(c.periods)-1 {
			c.periods[i].endAmount += c.periods[i].notYetReinvestedAmount //
		}
	}

	return Result{c.periods}
}
