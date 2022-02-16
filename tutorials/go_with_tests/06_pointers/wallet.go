package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(funds Bitcoin) {
	w.balance += funds
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//Do we realy need this? :-/
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(funds Bitcoin) error {
	if funds > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= funds
	return nil
}
