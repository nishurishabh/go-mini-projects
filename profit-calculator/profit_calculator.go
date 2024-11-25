package main

import "fmt"

func main() {

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculeFinances(revenue, expenses, taxRate)

	outputResult(ebt, profit, ratio)
}

func getUserInput(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}

// returns Earning before tax, Profit and ratio to earning to profit
func calculeFinances(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func outputResult(ebt, profit, ratio float64) {
	fmt.Printf("Earning before tax: %.1f \n", ebt)
	fmt.Printf("Profit: %.1f \n", profit)
	fmt.Printf("Ratio: %.2f \n", ratio)
}
