package config

type duration int

const (
	Hour  duration = 1
	Day            = 24 * Hour
	Week           = 7 * Day
	Month          = 730 * Hour
	Year           = (365 * 24) * Hour
)

// Config replicating the fields of the calculator from the site calcus.ru
type CommonConfig struct {
	StartAmount          float64
	DurationOfInvestment duration
	Percent              float64 // per year
	ReinvestmentPeriods  duration
	DepositPeriods       duration
	Deposit              float64
}

func (c CommonConfig) GetBaseConfig() BaseConfig {
	b := float64(c.DurationOfInvestment) / float64(c.ReinvestmentPeriods)
	a := int(b)
	d := float64(c.DepositPeriods) / float64(Month)
	cfg := BaseConfig{
		NumberOfPeriods: a,
		StartAmount:     c.StartAmount,
		Percent:         c.Percent / float64(Year) * float64(Month),
		Deposit:         c.Deposit,
		DepositEveryN:   int(d),
	}

	return cfg
}

type BaseConfig struct {
	NumberOfPeriods int
	StartAmount     float64
	Percent         float64
	PercentEveryN   int // Каждые сколько периодов начисляется процент
	Deposit         float64
	DepositEveryN   int
}

func (c BaseConfig) GetBaseConfig() BaseConfig {
	return c
}

type FundingConfig struct {
	FundingEveryNHours int
}
