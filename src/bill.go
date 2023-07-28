package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Generates new empty bill with given name
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// Formats the bill for prettier prints
func (b *bill) format() string {
	formattedString := fmt.Sprintf("%v's bill breakdown: \n", b.name)
	total := 0.0

	// List items (e.g. cake: ... $5.5)
	for item, price := range b.items {
		formattedString += fmt.Sprintf("%-8v ... $%0.2f \n", item+":", price)
		total += price
	}

	// Tip formatting
	formattedString += fmt.Sprintf("%-8v ... $%0.2f \n", "tip:", b.tip)

	// Total formatting
	formattedString += fmt.Sprintf("%-8v ... $%0.2f", "total:", total+b.tip)
	return formattedString

}

// Update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// Add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// Save bill to a file
func (b *bill) save() {
	data := []byte(b.format())
	fileName := fmt.Sprintf("%v-%v.txt", strings.ReplaceAll(b.name, " ", "-"), time.Now().Unix())
	filePath := fmt.Sprintf("bills/%v.txt", fileName)

	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Bill %v was saved in /bills!", fileName)
}
