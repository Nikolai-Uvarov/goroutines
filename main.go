package main

import (
	"account/bank"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	var C bank.BankClient = &bank.Client{}

	rand.Seed(time.Now().Unix())

	for i := 0; i < 10; i++ {

		fmt.Println("Starting random depositing № ", i+1)
		go bank.RandDeposit(C)

		if odd := i % 2; odd == 0 {
			fmt.Println("Starting random withdrawal № ", i/2+1)
			go bank.RandWithdrawal(C)
		}
	}

	var amount int
	var menu string

	for {
		_, err := fmt.Scanln(&menu)
		if err != nil {
			fmt.Println("Fatal error - incorrect input: ", err)
			os.Exit(3)
		}

		switch menu {
		case "balance":
			fmt.Println("Your current balance is: ", C.Balance())

		case "deposit":
			fmt.Println("Enter deposit amount: ")

			_, err := fmt.Scanln(&amount)
			if err != nil {
				fmt.Println("Fatal error - incorrect input: ", err)
				os.Exit(1)
			}

			C.Deposit(amount)

			fmt.Println("Your balance after deposit is: ", C.Balance())

		case "withdrawal":
			fmt.Println("Enter withdrawal amount: ")

			_, err := fmt.Scanln(&amount)
			if err != nil {
				fmt.Println("Fatal error - incorrect input: ", err)
				os.Exit(1)
			}

			err = C.Withdrawal(amount)
			if err != nil {
				fmt.Println("Withdrawal on demand failed: ", err)
				break
			}

			fmt.Println("Your balance after withdrawal is: ", C.Balance())

		case "exit":
			os.Exit(0)

		default:
			fmt.Println("Unsupported command. You can use commands: balance, deposit, withdrawal, exit")
		}
	}

}
