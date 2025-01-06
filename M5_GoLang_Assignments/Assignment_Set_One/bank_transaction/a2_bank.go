package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	DepositOption            = "1"
	WithdrawOption           = "2"
	ViewBalanceOption        = "3"
	TransactionHistoryOption = "4"
	ExitOption               = "5"
)

type Account struct {
	ID      int
	Name    string
	Balance float64
	History []string
}

var accounts = map[int]*Account{
	1: {ID: 1, Name: "Ashwat", Balance: 1000.0, History: []string{"Account created with rs1000"}},
	2: {ID: 2, Name: "Shaurya", Balance: 500.0, History: []string{"Account created with rs500"}},
}

func main() {
	fmt.Println("Welcome to the Bank Transaction System!")
	reader := bufio.NewReader(os.Stdin)

	for {
		displayMenu()
		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case DepositOption:
			handleDeposit(reader)
		case WithdrawOption:
			handleWithdraw(reader)
		case ViewBalanceOption:
			handleViewBalance(reader)
		case TransactionHistoryOption:
			handleTransactionHistory(reader)
		case ExitOption:
			fmt.Println("Exiting system. Thank you!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func displayMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. View Balance")
	fmt.Println("4. View Transaction History")
	fmt.Println("5. Exit")
}

func handleDeposit(reader *bufio.Reader) {
	account, amount, err := getAccountAndAmount(reader, "deposit")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	account.Balance += amount
	account.History = append(account.History, fmt.Sprintf("Deposited $%.2f", amount))
	fmt.Printf("Deposit successful. New balance: rs%.2f\n", account.Balance)
}

func handleWithdraw(reader *bufio.Reader) {
	account, amount, err := getAccountAndAmount(reader, "withdraw")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if amount > account.Balance {
		fmt.Println("Error: Insufficient balance.")
		return
	}

	account.Balance -= amount
	account.History = append(account.History, fmt.Sprintf("Withdrew rs%.2f", amount))
	fmt.Printf("Withdrawal successful. New balance: $%.2f\n", account.Balance)
}

func handleViewBalance(reader *bufio.Reader) {
	account, err := getAccountByID(reader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Account ID: %d\nName: %s\nBalance: rs%.2f\n", account.ID, account.Name, account.Balance)
}

func handleTransactionHistory(reader *bufio.Reader) {
	account, err := getAccountByID(reader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Transaction History for Account ID: %d\n", account.ID)
	for _, record := range account.History {
		fmt.Println("-", record)
	}
}

func getAccountAndAmount(reader *bufio.Reader, action string) (*Account, float64, error) {
	account, err := getAccountByID(reader)
	if err != nil {
		return nil, 0, err
	}

	fmt.Printf("Enter amount to %s: ", action)
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		return nil, 0, errors.New("invalid amount")
	}

	return account, amount, nil
}

func getAccountByID(reader *bufio.Reader) (*Account, error) {
	fmt.Print("Enter Account ID: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return nil, errors.New("invalid account ID")
	}

	account, exists := accounts[id]
	if !exists {
		return nil, errors.New("account not found")
	}

	return account, nil
}
