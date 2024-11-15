
package main

import "fmt"

func GetName() string {

	name := ""

	fmt.Println("Welcome to Joshua Kings Go_Lab Casino...")

	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("an error occured")
		return ""
	}
	fmt.Printf("Welcome %s, Let's Play!\n", name)
	return name
}

func GetBet(balance uint) uint {
	var bet uint

	for {
		fmt.Printf("Enter your bet, or 0 to quit (balance = $%d): ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Bet cannot be larger than balance.")
		} else {
			break
		}
	}
	return bet
}
