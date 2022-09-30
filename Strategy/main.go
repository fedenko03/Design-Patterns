package main

import "fmt"

type Payer interface {
	Pay(int) error
}

type Card struct {
	Balance int
	CVV     string
	Number  string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("You don't have enough money!")
	}
	c.Balance -= amount
	return nil
}

type KaspiQR struct {
	Money       int
	PhoneNumber string
}

func (a *KaspiQR) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("You don't have enough money!")
	}
	a.Money -= amount
	return nil
}

func Buy(p Payer) {
	switch p.(type) {
	case *Card:
		_, ok := p.(*Card)
		if !ok {
			fmt.Println("It looks like it's not a payment card")
			return
		}
		fmt.Println("Wait... The order is being processed")
	case *KaspiQR:
		_, ok := p.(*KaspiQR)
		if !ok {
			fmt.Println("It looks like it's not a Kaspi QR")
			return
		}
		fmt.Println("Wait... The order is being processed")
	default:
		fmt.Println("Payment is possible ONLY by payment card or via Kaspi QR!")
		return
	}

	err := p.Pay(1000000)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Thank you for buying iPhone 14 via %T\n\n", p)
}

func main() {
	myCard := &Card{
		Balance: 3400000,
		CVV:     "840",
		Number:  "483922315434",
	}

	myKaspiQR := &KaspiQR{
		Money: 1222333,
	}

	fmt.Println("--------------------------------------------------")
	Buy(myCard)
	fmt.Println("--------------------------------------------------")
	Buy(myKaspiQR)
	fmt.Println("--------------------------------------------------")
	Buy(myKaspiQR)
}
