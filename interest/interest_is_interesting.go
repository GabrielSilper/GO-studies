package interest

const (
	lessThan0               = 3.213
	greaterThan0Less1000    = 0.5
	greaterThan1000Less5000 = 1.621
	greaterThan5000         = 2.475
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return lessThan0
	case balance >= 0 && balance < 1000:
		return greaterThan0Less1000
	case balance >= 1000 && balance < 5000:
		return greaterThan1000Less5000
	default:
		return greaterThan5000
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * (float64(InterestRate(balance)) / 100.0)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	countYears := 0
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		countYears++
	}
	return countYears
}
