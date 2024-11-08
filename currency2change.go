package main

import (
	"os"
	"strconv"
	"strings"
)

// I know it's ugly. I'm sorry.
func printHelp() {
	println(
`currency2change is a quick program for running the cashier's algorithm on arbitrary amounts of money.


Usage:

        currency2change <dollar/cent amount>

        i.e.
        currency2change $123.45
        currency2change $123.4
        currency2change $123.
        currency2change $123
        currency2change 123
        currency2change .2
        curremcy2change $999,999,999.99


Example output:

        > ./currency2change $25,709.82
        257 - $100 bill
        1 - $5 bill
        4 - $1 bill
        3 - Quarter
        1 - Nickle
        2 - Penny


NOTE:
        In most shells, dollar signs must be escaped, or they will cause
        problems.


TODO:
        Pad the amounts nicely -- the current output formatting irritates me.
`,
	)
}

func printCashiersBills(b int) {
	// We assume we're using the most common bill types here
	// Yes, the federal reserve issues $2 bills still; no, I do not care
	// Output format is as such, where n is some nonzero int
	// n - $100 bill
	// n - $50 bill
	// n - $20 bill
	// 0 - $10 bill (This line would be omitted because n == 0)
	// n - $5 bill
	// n - $1 bill

	// Array makes it easy to iterate, and we will never have a final remainder
	// because we have a one dollar bill, so we're safe for now.
	// Also makes it easier to add new bills if needed, but some bills will
	// make it so that our greedy approach is no longer optimal in the general
	// case.
	billTypes := []int{100, 50, 20, 10, 5, 1}

	for _, billType := range(billTypes) {
		bills := b / billType

		if bills == 0 {
			continue
		}

		b %= billType

		println(bills, "-", "$" + strconv.Itoa(billType), "bill")
	}
}

func printCashiersCoins(c int) {
	// We assume we're using the most common coin types here
	// Yes, the federal reserve issues 50 cent coins still; no, I do not care
	// Output format is like above, where n is some nonzero int
	// n - Quarter
	// n - Dime
	// n - Nickle
	// 0 - Penny (omitted because n = 0)

	// Like with before, Array makes it easy to iterate, and we will never have
	// a final remainder because of the pennies, so we're safe unless we invent
	// a seven cent coin or something. You know, out of spite for me
	// specifically.
	// Unfortunately, because the example output uses coin names here, so will
	// we. We achieve this with a parallel array, because I'm too lazy to make
	// a struct, and they would make our code less pretty.
	// Also note that we're actually grabbing things by index here, not just
	// abusing for-each. Que un lastima

	coinTypes := []int{25, 10, 5, 1}
	coinTypeNames := []string{"Quarter", "Dime", "Nickle", "Penny"}


	for i := range(coinTypes) {
		coins := c / coinTypes[i]

		if coins == 0 {
			continue
		}

		c %= coinTypes[i]

		println(coins, "-", coinTypeNames[i])
	}
}

func main() {
	if len(os.Args) != 2 {
		printHelp()
		return
	}

	currencyString := os.Args[1]

	// Trim everything prior to and including the dollar sign for ease of use later
	dollarSignLocation := strings.Index(currencyString, "$")
	if dollarSignLocation != -1 {
		currencyString = currencyString[dollarSignLocation+1:]
	}

	// Rip commas out
	currencyString = strings.ReplaceAll(currencyString, ",", "")

	// Split the string into two potential-integers
	// We do not use floats here because we are dealing with money, and floating
	// point imprecision would be a very bad thing in this case.
	currencyStringSliced := strings.Split(currencyString, ".")

	// Late input validation
	if 2 < len(currencyStringSliced) {
		printHelp()
		return
	}

	// Finally make our values into integers
	// Would have liked to use a sentry here, but I'd rather not over-abstract
	// just to be able to do so
	inputHasDollars := 0 < len(currencyStringSliced[0])
	if inputHasDollars {
		bills, err := strconv.Atoi(currencyStringSliced[0])
		if err != nil {
			printHelp()
			return
		}

		// Handles the cashier's algorithm for dollar bills, prints results so that
		// we needn't worry about it in main.
		printCashiersBills(bills)
	}

	inputHasCents := len(currencyStringSliced) < 2 || len(currencyStringSliced[1]) < 1
	if  inputHasCents {
		return
	}

	cents, err := strconv.Atoi(currencyStringSliced[1])
	if err != nil {
		printHelp()
		return
	}

	// In the case where we get something like $0.2, our string parse will
	// short change us by an order of magnitude, so we correct that here if need
	// be.
	if len(currencyStringSliced[1]) == 1 {
		cents *= 10
	}

	// Handles the Cashier's algorithm for cents, and is nice enough to print
	// the results for us so we don't have to do it in main
	printCashiersCoins(cents)
}
