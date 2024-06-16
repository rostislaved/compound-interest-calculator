package main

import (
	"fmt"
	"math"
	"strings"
)

func PrintStats(results ...Result) {
	var (
		blocks1 []string
		blocks2 []string
		blocks3 []string
		blocks4 []string
		blocks5 []string
		blocks6 []string
	)

	maxPercentSum := math.MaxFloat64
	maxLastAmount := math.MaxFloat64

	for _, result := range results {
		lastAmount := result.LastAmount()
		lastPeriod := result.LastPeriod()

		maxPercentSum = min(maxPercentSum, lastPeriod.percentSum)
		maxLastAmount = min(maxLastAmount, lastAmount)
	}

	for _, result := range results {
		startAmount := result.FirstPeriod().startAmount
		lastAmount := result.LastAmount()
		lastPeriod := result.LastPeriod()
		allDeposits := startAmount + lastPeriod.depositSum

		block1 := fmt.Sprintf("%.2f (%.1f%%) (%.1f%%)", startAmount, getPercent(startAmount, lastAmount), getPercent(startAmount, allDeposits))
		block2 := fmt.Sprintf("%.2f (%.1f%%) (%.1f%%)", lastPeriod.depositSum, getPercent(lastPeriod.depositSum, lastAmount), getPercent(lastPeriod.depositSum, allDeposits))
		block3 := fmt.Sprintf("%.2f (%.1f%%) (%.1f%%)", lastPeriod.percentSum, getPercent(lastPeriod.percentSum, lastAmount), getPercent(lastPeriod.percentSum, allDeposits))
		block4 := fmt.Sprintf("%.2f (100%%) (%.1f%%)", lastAmount, getPercent(lastAmount, allDeposits))
		block5 := fmt.Sprintf("%.2f (%.1f%%)", lastPeriod.percentSum-maxPercentSum, getPercent(lastPeriod.percentSum, maxPercentSum))
		block6 := fmt.Sprintf("%.2f (%.1f%%)", lastAmount-maxLastAmount, getPercent(lastAmount, maxLastAmount))
		length := max(len(block1), len(block2), len(block3), len(block4))

		block1 = complementWithSpaces(block1, length)
		block2 = complementWithSpaces(block2, length)
		block3 = complementWithSpaces(block3, length)
		block4 = complementWithSpaces(block4, length)
		block5 = complementWithSpaces(block5, length)
		block6 = complementWithSpaces(block6, length)

		blocks1 = append(blocks1, block1)
		blocks2 = append(blocks2, block2)
		blocks3 = append(blocks3, block3)
		blocks4 = append(blocks4, block4)
		blocks5 = append(blocks5, block5)
		blocks6 = append(blocks6, block6)
	}

	line1 := fmt.Sprintf("| Start Amount   | %s |", strings.Join(blocks1, " | "))
	line2 := fmt.Sprintf("| Total Deposits | %s |", strings.Join(blocks2, " | "))
	line3 := fmt.Sprintf("| Total Income   | %s |", strings.Join(blocks3, " | "))
	line4 := fmt.Sprintf("| Total Amount   | %s |", strings.Join(blocks4, " | "))
	line5 := fmt.Sprintf("| Diff Income    | %s |", strings.Join(blocks5, " | "))
	line6 := fmt.Sprintf("| Diff Amount    | %s |", strings.Join(blocks6, " | "))

	fmt.Println(strings.Repeat("-", len(line1)))
	fmt.Println(line1)
	fmt.Println(line2)
	fmt.Println(line3)
	fmt.Println(line4)
	fmt.Println(strings.Repeat("-", len(line1)))
	fmt.Println(line5)
	fmt.Println(line6)
	fmt.Println(strings.Repeat("-", len(line1)))
	fmt.Println()
}

func complementWithSpaces(s string, length int) string {
	diff := length - len([]rune(s))

	if diff > 0 {
		s = strings.Repeat(" ", diff) + s
	}

	return s
}

func PrintStatsByPeriod(r Result) {
	header := []string{"№", "Start Amount", "Percent", "Deposit", "End Amount", "depositCumSum", "percentCumSum"}

	var arrayOfBlocks [][]string

	for i, period := range r.periods {
		var blocks []string

		blocks = append(blocks,
			fmt.Sprintf("%d", i+1),
			fmt.Sprintf("%.2f", period.startAmount),
			fmt.Sprintf("%.2f", period.increaseByPercent),
			fmt.Sprintf("%.2f", period.deposit),
			fmt.Sprintf("%.2f", period.EndAmount()),
			fmt.Sprintf("%.2f (%4.1f%%)", period.depositSum, getPercent(period.depositSum, period.EndAmount())),
			fmt.Sprintf("%.2f (%4.1f%%)", period.percentSum, getPercent(period.percentSum, period.EndAmount())),
		)

		arrayOfBlocks = append(arrayOfBlocks, blocks)
	}

	arrayOfBlocks = append([][]string{header}, arrayOfBlocks...)

	blocksMaxLength := make([]int, len(arrayOfBlocks[0]))

	for _, blocks := range arrayOfBlocks {
		for i, block := range blocks {
			blocksMaxLength[i] = max(blocksMaxLength[i], len(block))
		}
	}

	for i := range arrayOfBlocks {
		for j := range arrayOfBlocks[i] {
			arrayOfBlocks[i][j] = complementWithSpaces(arrayOfBlocks[i][j], blocksMaxLength[j])
		}
	}

	var lines []string

	for _, blocks := range arrayOfBlocks {

		line := fmt.Sprintf("| %s |\n", strings.Join(blocks, " | "))
		lines = append(lines, line)
	}

	for i, line := range lines {
		l := len([]rune(line)) - 1
		dashLine := strings.Repeat("-", l)

		if i == 0 {
			fmt.Println(dashLine)
		}

		fmt.Print(line)

		switch i {
		case 0, len(arrayOfBlocks) - 1:
			fmt.Println(dashLine)
		}

	}

}

func getPercent(a, b float64) float64 {
	return a / b * 100
}
