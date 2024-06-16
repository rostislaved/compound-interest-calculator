package main

type Result struct {
	periods []Period
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
