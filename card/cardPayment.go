package card

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Card struct {
	cardNO  string
	balance float64
}

func NewCard(bal float64) *Card {
	return &Card{
		cardNO:  generateCardNo(),
		balance: bal,
	}
}

func (c *Card) Pay(a float64) {
	fmt.Printf("Paying %.2f Naira with Card !\n", a)
	c.balance = c.balance - a
}

func (c *Card) ReplaceCard(deposit float64) *Card {
	newCard := NewCard(deposit)
	newCard.balance += c.balance
	return newCard
}

func (c *Card) GetInfo() {
	line := "*"
	fmt.Printf("%v\n", strings.Repeat(line, 34))
	fmt.Printf("* Card Number: %v *\n", c.cardNO)
	fmt.Printf("* Card Balance: %v             *\n", c.balance)
	fmt.Printf("%v\n", strings.Repeat(line, 34))
}

func generateCardNo() string {
	rand.Seed(time.Now().UnixNano())
	length := 16
	cardNo := ""
	dictionary := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < length; i++ {
		//get a random number between 0 and 10
		pos := rand.Intn(10)
		cardNo += strconv.Itoa(dictionary[pos])
	}
	return cardNo
}
