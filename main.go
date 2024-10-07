package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	bt := BudgetTracker{}

	for {
		fmt.Println("\n--- Personal Budget Tracker ---")
		fmt.Println("1. Add Transaction")
		fmt.Println("2. View Transactions")
		fmt.Println("3. Show Total income")
		fmt.Println("4. Show Total expenses")
		fmt.Println("5 Save Transactions To CSV")
		fmt.Println("6. Exit")
		fmt.Println("Choose an option: ")

		var choice int

		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter Amount. ")
			var amount float64
			fmt.Scanln(&amount)

			fmt.Println("Enter Category. ")
			var category string
			fmt.Scanln(&category)

			fmt.Println("Enter Type Income/Expense. ")
			var typeTransaction string
			fmt.Scanln(&typeTransaction)

			//Add Transaction
			bt.AddTransaction(amount, category, typeTransaction)

			fmt.Println("Trasaction Added Successfully" + time.Now().String())

		case 2:
			//View Transactions
			bt.DisplayTransactions()

		case 3:
			//Show Total Income
			fmt.Println("Total Income: ", bt.CalculateTotal("Income"))
		case 4:
			//Show Total Expenses
			fmt.Println("Total Expenses: ", bt.CalculateTotal("Expense"))

		case 5:
			//Save Transactions To CSV
			fmt.Println("Enter The File Name: e.g(transaction.csv) ")
			var fileName string
			fmt.Scanln(&fileName)
			if err := bt.SaveToCSV(fileName); err != nil {
				fmt.Println("Error Saving Transactions To CSV: ", err)
			}
		case 6:
			//Exit
			fmt.Println("Exiting...")
			os.Exit(1)

		default:
			fmt.Println("Invalid Choice Try Again ...")

		}

	}
}

// Transaction is a struct that holds the transaction information
type Transaction struct {
	ID       int
	Amount   float64
	Cateogry string
	Date     time.Time
	Type     string
}

// BudgetTracker is a struct that holds the transaction information
type BudgetTracker struct {
	Transactions []Transaction
	NextID       int
}

// interface for common behaviors
type FinancialRecored interface {
	GetAmmount() float64
	GetType() string
}

// implemntation of FinancialRecored interface

func (t Transaction) GetAmmount() float64 {
	return t.Amount
}

func (t Transaction) GetType() string {
	return t.Type
}

// add a new transaction
func (bt *BudgetTracker) AddTransaction(amount float64, category string, tType string) {
	newTrasaction := Transaction{
		ID:       bt.NextID,
		Amount:   amount,
		Cateogry: category,
		Date:     time.Now(),
		Type:     tType,
	}
	bt.Transactions = append(bt.Transactions, newTrasaction)
	bt.NextID++
}

// DisplayTansaction
func (bt BudgetTracker) DisplayTransactions() {
	fmt.Println("ID\tAmount\tCategory\tDate\tType")
	for _, t := range bt.Transactions {
		fmt.Printf("%d\t%.2f\t%s\t%s\t%s\n", t.ID, t.Amount, t.Cateogry, t.Date.Format("2006-01-02"), t.Type)
	}
}

// CalculateTotal income and expense
func (bt BudgetTracker) CalculateTotal(tType string) float64 {

	var total float64
	for _, t := range bt.Transactions {
		if t.Type == tType {
			total += t.Amount
		}
	}
	return total

}

// SaveToCSV
func (bt BudgetTracker) SaveToCSV(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	write := csv.NewWriter(file)
	defer write.Flush()

	if err := write.Write([]string{"ID", "Amount", "Category", "Date", "Type"}); err != nil {
		return err
	}

	for _, t := range bt.Transactions {
		record := []string{
			strconv.Itoa(t.ID),
			fmt.Sprintf("%.2f", t.Amount),
			t.Cateogry,
			t.Date.Format("2006-01-02"),
			t.Type,
		}
		if err := write.Write(record); err != nil {
			return err
		}
	}

	fmt.Println("Transactions saved to CSV file. " + filename)

	return nil
}
