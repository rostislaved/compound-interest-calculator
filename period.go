package main

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
