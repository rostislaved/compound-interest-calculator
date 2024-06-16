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

		block1 := fmt.Sprintf("%.2f (%.1f%%)", startAmount, getPercent(startAmount, lastAmount))
		block2 := fmt.Sprintf("%.2f (%.1f%%)", lastPeriod.depositSum, getPercent(lastPeriod.depositSum, lastAmount))
		block3 := fmt.Sprintf("%.2f (%.1f%%)", lastPeriod.percentSum, getPercent(lastPeriod.percentSum, lastAmount))
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
	diff := length - len([]rune(s))

	if diff > 0 {
		s = strings.Repeat(" ", diff) + s
	}

	return s
}

func PrintStatsByPeriod(r Result) {
	header := []string{"№", "startAmount", "increaseByPercent", "Deposit", "endAmount", "depositCumSum", "percentCumSum"}

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
