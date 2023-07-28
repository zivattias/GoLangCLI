package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Input retrieval helper function
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

// Instantiate a new bill
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("What's your name? ", reader)
	b := newBill(name)

	return b
}

// Display available options to manipulate the bill
func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - Add item, s - Save bill, t - Add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Invalid price (must be numeric)")
			promptOptions(b)
			return
		}

		b.addItem(name, p)
		fmt.Printf("Added item %v - %v \n", name, p)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Invalid tip (must be numeric)")
			promptOptions(b)
			return
		}

		b.updateTip(t)
		fmt.Printf("Added tip $%v \n", t)
		promptOptions(b)
	case "s":
		fmt.Println("Saving bill...")
		b.save()
	default:
		fmt.Println("Invalid choice, try again")
		promptOptions(b)
	}
}
