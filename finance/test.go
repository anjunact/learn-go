package main

import (
	"fmt"

	"github.com/FlashBoys/go-finance"
)

func main() {
	// 15-min delayed full quote for Apple.
	symbols := []string{"AAPL", "TWTR", "FB"}
	quotes, err := finance.GetQuotes(symbols)
	if err == nil {
		fmt.Println(quotes)
	}
}