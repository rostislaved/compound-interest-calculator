package config

type duration int

const (
	Hourly  duration = 1
	Daily            = 24 * Hourly
	Weekly           = 7 * Daily
	Monthly          = 730 * Hourly          // ceil(365*24 + 6)/12
	Yearly           = (365*24 + 6) * Hourly // 6 - учет високосных лет
)

type CommonConfig struct {
	StartAmount          float64
	DurationOfInvestment duration
	Percent              float64 // в год
	ReinvestmentPeriods  duration
	DepositPeriods       duration
	Deposit              float64
}

func (c CommonConfig) GetBaseConfig() BaseConfig {
	b := float64(c.DurationOfInvestment) / float64(c.ReinvestmentPeriods)
	a := int(b)
	d := float64(c.DepositPeriods) / float64(Monthly)
	cfg := BaseConfig{
		NumberOfPeriods: a,
		StartAmount:     c.StartAmount,
		Percent:         c.Percent / float64(Yearly) * float64(Monthly),
		Deposit:         c.Deposit,
		EveryN:          int(d),
	}

	return cfg
}

type BaseConfig struct {
	NumberOfPeriods int
	StartAmount     float64
	Percent         float64
	Deposit         float64
	EveryN          int
}

func (c BaseConfig) GetBaseConfig() BaseConfig {
	return c
}

type FundingConfig struct {
	FundingEveryNHours int
}
