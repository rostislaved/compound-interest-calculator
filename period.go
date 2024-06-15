package main

type Period struct {
	startAmount       float64
	increaseByPercent float64
	percent           float64
	deposit           float64

	depositSum float64
	percentSum float64

	notYetReinvestedAmount float64
}

func (p *Period) EndAmount() float64 {
	return p.startAmount + p.increaseByPercent + p.deposit
}

func (p *Period) calculatePeriod(previousPeriodEndAmount float64) {
	p.startAmount = previousPeriodEndAmount

	p.increaseByPercent = p.startAmount * (p.percent / 100)
}
