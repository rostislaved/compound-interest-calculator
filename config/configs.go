package config

type duration int

const (
	Hour  duration = 1
	Day            = 24 * Hour
	Week           = 7 * Day
	Month          = 730 * Hour
	Year           = (365 * 24) * Hour
)

// Config replicating the fields of the calculator from the site calcus.ru.
type CommonConfig struct {
	StartAmount          float64
	DurationOfInvestment duration

	Percent float64 // per year

	Deposit       float64
	DepositEveryN duration

	ReinvestEveryN duration
}
type BaseConfig struct {
	StartAmount     float64
	NumberOfPeriods int

	Percent       float64
	PercentEveryN int //  // > 0 // Каждые сколько периодов начисляется процент

	Deposit       float64
	DepositEveryN int // > 0

	ReinvestEveryN int // > 0
}

func (c CommonConfig) GetBaseConfig() BaseConfig {
	periodDuration := min(c.DurationOfInvestment, c.ReinvestEveryN, c.DepositEveryN)

	numberOfPeriods := c.DurationOfInvestment / periodDuration

	depositEveryN := c.DepositEveryN / periodDuration

	reinvestEveryN := c.ReinvestEveryN / periodDuration

	cfg := BaseConfig{
		NumberOfPeriods: int(numberOfPeriods),
		StartAmount:     c.StartAmount,
		Percent:         c.Percent / float64(Year) * float64(Month),
		PercentEveryN:   1, // hardcoded
		Deposit:         c.Deposit,
		DepositEveryN:   int(depositEveryN),
		ReinvestEveryN:  int(reinvestEveryN),
	}

	return cfg
}

func (c BaseConfig) GetBaseConfig() BaseConfig {
	return c
}

type FundingConfig struct {
	FundingEveryNHours int
}
