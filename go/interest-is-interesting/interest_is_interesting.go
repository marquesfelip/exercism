package interest

func InterestRate(balance float64) float32 {
  switch  {
    case balance >= 5000:
      return 2.475
    case balance >= 1000:
      return 1.621
    case balance >= 0:
      return 0.5
    default:
      return 3.213
  }
}

func Interest(balance float64) float64 {
  return balance * float64(InterestRate(balance) / 100)
}

func AnnualBalanceUpdate(balance float64) float64 {
  return Interest(balance) + balance
}

func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
  var year int = 0

  for year = 0; balance < targetBalance; year++ {
    balance = AnnualBalanceUpdate(balance)
  }

  return year
}
