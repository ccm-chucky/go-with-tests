package pointers

import (
	"errors"
	"fmt"
)

type Wallet struct{
	balance Bitcoin
}

type Bitcoin int

type Striger interface{
	String() string
}

func (b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC",b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
    w.balance += amount
}


func (w *Wallet) Balance() Bitcoin{
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin)error{
	if amount > w.balance {
        return errors.New("oh no")
    }

    w.balance -= amount
    return nil
}