package bank

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BankClient interface {
	// Deposit deposits given amount to clients account
	Deposit(amount int)

	// Withdrawal withdraws given amount from clients account.
	// return error if clients balance less the withdrawal amount
	Withdrawal(amount int) error

	// Balance returns clients balance
	Balance() int
}

// in-memory account db for 1 clientttt
type Client struct {
	m       sync.RWMutex
	account int
}

func (c *Client) Deposit(amount int) {
	c.m.Lock()
	defer c.m.Unlock()
	c.account += amount
}

func (c *Client) Withdrawal(amount int) error {
	c.m.Lock()
	defer c.m.Unlock()

	if c.account < amount {
		return fmt.Errorf("Withdrawal declined: not enough funds")
	}
	c.account -= amount
	return nil
}

func (c *Client) Balance() int {
	c.m.RLock()
	defer c.m.RUnlock()

	return c.account
}

func RandWithdrawal(c BankClient) {

	for {
		time.Sleep(time.Millisecond * time.Duration(1000-rand.Intn(500)))
		amount := rand.Intn(5) + 1

		err := c.Withdrawal(amount)

		if err != nil {
			fmt.Println("Withdrawal attempt failed: ", err)
		}
	}
}

func RandDeposit(c BankClient) {

	for {
		time.Sleep(time.Millisecond * time.Duration(1000-rand.Intn(500)))
		amount := rand.Intn(10) + 1

		c.Deposit(amount)
	}
}
