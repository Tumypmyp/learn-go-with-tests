package arrays

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{from.Name, to.Name, sum}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransaction, account)
}

func applyTransaction(a Account, t Transaction) Account {
	if t.From == a.Name {
		a.Balance -= t.Sum
	}
	if t.To == a.Name {
		a.Balance += t.Sum
	}
	return a
}

type Account struct {
	Name    string
	Balance float64
}
