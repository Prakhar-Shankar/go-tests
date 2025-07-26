package pointer



type Bitcoin int //go lrts you create new type from the existing one

type Wallet struct {
	balance Bitcoin
}

//Here we have added the pointer so instead of pasting the whole Wallet we are directing the address of the wallet. Also we can use the referencing in the reciever type 

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcon {
	return w.balance
}