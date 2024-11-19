package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(11) //TODO: revert to expected value of 10 after inspecting failure message

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
