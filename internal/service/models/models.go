package models

type Deposit struct {
	From   string
	Amount int64
}

type Withdraw struct {
	To     string
	Amount int64
}
