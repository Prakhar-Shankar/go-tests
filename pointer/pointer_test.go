package pointer

import "testing"


//so when using the method the test is sending the copy of the balance and hence the pointer (or memory address) is different so we can use the pointer to solve this issue.

func TestWallet(t *testing.T){

	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10)) // to use the new type "Bitcoin" we have to write Bitcoin(10)

	got := wallet.Balance()

	want := Bitcoin(10) 
	
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}