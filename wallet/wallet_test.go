package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("test deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("test withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		wallet.Withdraw(Bitcoin(3))

		got := wallet.Balance()
		want := Bitcoin(7)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

}
