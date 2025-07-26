package pointer

import "testing"


//so when using the method the test is sending the copy of the balance and hence the pointer (or memory address) is different so we can use the pointer to solve this issue.

func TestWallet(t *testing.T){

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin){
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

 t.Run("deposit", func(t *testing.T){
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10)) // to use the new type "Bitcoin" we have to write Bitcoin(10)
	assertBalance(t, wallet, Bitcoin(10))
})

t.Run("withdraw", func(t *testing.T){
	wallet := Wallet{balance: Bitcoin(20)}
	wallet.Withdraw(Bitcoin(10))
	assertBalance(t, wallet, Bitcoin(10))
})

t.Run("withdraw insufficient funds", func(t *testing.T){
	startingBalance := Bitcoin(20)
	wallet := Wallet{startingBalance}
	err := wallet.Withdraw(Bitcoin(100))

	assertBalance(t, wallet, startingBalance)

	if err == nil {
		t.Errorf("wanted an error but didn't get one")
	}
})
}