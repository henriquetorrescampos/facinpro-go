package domain

// import "errors"

type Account struct {
	ID      string
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}
