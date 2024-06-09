package main

import (
	"fmt"
	"math"
)

type Result struct {
	periods []Period
}

func (r Result) PrintStats() {
	var totalDeposits float64
	var totalIncome float64

	for _, period := range r.periods {
		totalDeposits += period.deposit
		totalIncome += period.increaseByPercent
	}

	startAmount := r.FirstPeriod().startAmount
	lastAmount := r.LastAmount()

	fmt.Printf("Start Amount: %.2f (%.d%%)\n", startAmount, getRoundedPercent(startAmount, lastAmount))
	fmt.Printf("Total Deposits: %.2f (%.d%%)\n", totalDeposits, getRoundedPercent(totalDeposits, lastAmount))
	fmt.Printf("Total Income: %.2f (%.d%%)\n", totalIncome, getRoundedPercent(totalIncome, lastAmount))
	fmt.Printf("Total Amount: %.2f (100%%)\n", lastAmount)
}

func (r Result) PrintStatsByPeriod() {
	fmt.Println(" №\t\tstartAmount\tincreaseByPercent\tDeposit\tendAmount")
	for i, period := range r.periods {
		fmt.Printf("[%d]:\t\t%.2f\t\t%.2f\t\t\t%.2f\t%.2f\n", i, period.startAmount, period.increaseByPercent, period.deposit, period.EndAmount())
	}
}

func getRoundedPercent(a, b float64) int {
	return int(math.Round(a / b * 100))
}

func (r Result) Print() {
	a := r.periods[0].startAmount
	b := r.LastAmount()
	c := b - a

	const n = 160 // TODO

	diff := c / n

	for _, period := range r.periods {
		r := (period.EndAmount() - a) / diff
		for i := 0; i < int(r); i++ {
			fmt.Printf("|")
		}

		fmt.Println()
	}
}

// Last отдает последнее значение amount. После всех расчетов.
func (r Result) LastAmount() float64 {
	lastPeriod := r.LastPeriod()

	return lastPeriod.EndAmount()
}

func (r Result) LastPeriod() Period {
	return r.periods[len(r.periods)-1]
}

func (r Result) FirstPeriod() Period {
	return r.periods[0]
}
