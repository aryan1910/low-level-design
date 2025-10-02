package main

import (
	"errors"
	"fmt"
)

type Card struct {
	CardNumber    string
	AccountNumber string
	Balance       float64
	PIN           string
}

type CardDb interface {
	GetCardDetails(cardNumber string) (*Card, error)
}

type MockCardStore struct {
	data map[string]Card
}

func (m *MockCardStore) GetCardDetails(cardNumber string) (*Card, error) {
	if card, exists := m.data[cardNumber]; exists {
		return &card, nil
	}
	return nil, errors.New("card not found")
}

type ATMReady struct {
	Store CardDb
}

func (r *ATMReady) InsertCard(cardNumber string) (*ATMCardInserted, *Card, error) {
	card, err := r.Store.GetCardDetails(cardNumber)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to insert card: %v", err)
	}
	fmt.Println("Card inserted. Please enter your PIN.")
	return &ATMCardInserted{}, card, nil
}

type ATMCardInserted struct{}

func (c *ATMCardInserted) EnterPIN(card *Card, enteredPIN string) (*ATMEnterPIN, error) {
	if enteredPIN == card.PIN {
		fmt.Println("PIN correct. Proceed to transaction.")
		return &ATMEnterPIN{}, nil
	}
	return nil, errors.New("incorrect PIN")
}

func (c *ATMCardInserted) EjectCard() *ATMReady {
	fmt.Println("Card ejected.")
	return &ATMReady{}
}

type ATMEnterPIN struct{}

func (e *ATMEnterPIN) WithdrawCash(card *Card, amount float64) (*ATMWithdrawCash, error) {
	if amount > card.Balance {
		return nil, errors.New("insufficient balance")
	}
	card.Balance -= amount
	fmt.Printf("Withdrawal successful. Amount: %.2f, Remaining Balance: %.2f\n", amount, card.Balance)
	return &ATMWithdrawCash{}, nil
}

func (e *ATMEnterPIN) EjectCard() *ATMReady {
	fmt.Println("Card ejected.")
	return &ATMReady{}
}

type ATMWithdrawCash struct{}

func (w *ATMWithdrawCash) EjectCard() *ATMReady {
	fmt.Println("Card ejected. Returning ATM to ready state.")
	return &ATMReady{}
}

func main() {
	store := &MockCardStore{
		data: map[string]Card{
			"1111-2222-3333-4444": {"1111-2222-3333-4444", "123456789", 1000.0, "1234"},
			"5555-6666-7777-8888": {"5555-6666-7777-8888", "987654321", 500.0, "5678"},
		},
	}

	var atmState interface{} = &ATMReady{Store: store}

	var card *Card
	if state, ok := atmState.(*ATMReady); ok {
		var err error
		atmState, card, err = state.InsertCard("1111-2222-3333-4444")
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if state, ok := atmState.(*ATMCardInserted); ok {
		var err error
		atmState, err = state.EnterPIN(card, "1234")
		if err != nil {
			fmt.Println(err)
			atmState = state.EjectCard()
		}
	}

	if state, ok := atmState.(*ATMEnterPIN); ok {
		var err error
		atmState, err = state.WithdrawCash(card, 200)
		if err != nil {
			fmt.Println(err)
			atmState = state.EjectCard()
		}
	}

	if state, ok := atmState.(*ATMWithdrawCash); ok {
		atmState = state.EjectCard()
	}

	if _, ok := atmState.(*ATMReady); ok {
		fmt.Println("ATM is back to ready state.")
	}
}
