package err

import "testing"

func TestWallet(t *testing.T){
	wallet := Wallet{}
	wallet.Deposit(20);

	got := wallet.Balance()
	want := 20.0

	if got != want {
		t.Errorf("Got %f instead of %f", got, want)
	}
}