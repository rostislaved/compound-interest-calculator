package main

import (
	"fmt"
	"strings"
)

func PrintStats(results ...Result) {

	var blocks1 []string
	var blocks2 []string
	var blocks3 []string
	var blocks4 []string

	for _, result := range results {
		startAmount := result.FirstPeriod().startAmount
		lastAmount := result.LastAmount()
		lastPeriod := result.LastPeriod()

		block1 := fmt.Sprintf("%.2f (%.1f%%)", startAmount, getRoundedPercent(startAmount, lastAmount))
		block2 := fmt.Sprintf("%.2f (%.1f%%)", lastPeriod.depositSum, getRoundedPercent(lastPeriod.depositSum, lastAmount))
		block3 := fmt.Sprintf("%.2f (%.1f%%)", lastPeriod.percentSum, getRoundedPercent(lastPeriod.percentSum, lastAmount))
		block4 := fmt.Sprintf("%.2f (100%%) ", lastAmount)
		length := max(len(block1), len(block2), len(block3), len(block4))

		block1 = complementWithSpaces(block1, length)
		block2 = complementWithSpaces(block2, length)
		block3 = complementWithSpaces(block3, length)
		block4 = complementWithSpaces(block4, length)

		blocks1 = append(blocks1, block1)
		blocks2 = append(blocks2, block2)
		blocks3 = append(blocks3, block3)
		blocks4 = append(blocks4, block4)
	}

	line1 := fmt.Sprintf("| Start Amount   | %s |", strings.Join(blocks1, " | "))
	line2 := fmt.Sprintf("| Total Deposits | %s |", strings.Join(blocks2, " | "))
	line3 := fmt.Sprintf("| Total Income   | %s |", strings.Join(blocks3, " | "))
	line4 := fmt.Sprintf("| Total Amount   | %s |", strings.Join(blocks4, " | "))

	fmt.Println(strings.Repeat("-", len(line1)))
	fmt.Println(line1)
	fmt.Println(line2)
	fmt.Println(line3)
	fmt.Println(line4)
	fmt.Println(strings.Repeat("-", len(line1)))
	fmt.Println()
}

func complementWithSpaces(s string, length int) string {
	diff := length - len(s)

	if diff > 0 {
		s = strings.Repeat(" ", diff) + s
	}

	return s
}

func PrintStatsByPeriod(r Result) {
	fmt.Println(" №\t\tstartAmount\tincreaseByPercent\tDeposit\tendAmount\tdepositCumSum\t\tpercentCumSum")

	format := "[%d]:\t\t%.2f\t\t%.2f\t\t\t%.2f\t%.2f\t\t%.2f(%.d%%)\t\t%.2f(%.d%%)\n"

	for i, period := range r.periods {
		fmt.Printf(format,
			i+1,
			period.startAmount,
			period.increaseByPercent,
			period.deposit,
			period.EndAmount(),
			period.depositSum, getRoundedPercent(period.depositSum, period.EndAmount()),
			period.percentSum, getRoundedPercent(period.percentSum, period.EndAmount()),
		)
	}
}

func getRoundedPercent(a, b float64) float64 {
	return a / b * 100
}

func Print(r Result) {
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
