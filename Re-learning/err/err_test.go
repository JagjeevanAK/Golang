package err

import "testing"

func TestWallet(t *testing.T){
	wallet := Wallet{}
	wallet.Deposit(20);

	got := wallet.Balance()
	var want Bitcoin = 20.0

	if got != want {
		t.Errorf("Got %f instead of %f", got, want)
	}
}