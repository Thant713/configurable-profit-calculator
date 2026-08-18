package main

import "fmt"

func main() {
	for {
		fmt.Println("Choose mode: 1) Adjust items  2) View items")
		choice := getInput("> ")

		switch choice {
		case 1:
			adjustMenu()
		case 2:
			viewMenu()
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func adjustMenu() {
	fmt.Println("1) Add  2) Edit  3) Remove")
	choice := getInput("> ")

	switch choice {
	case 1:
		item := Item{}
		item.AddItem()
	case 2:
		// edit - later
	case 3:
		// remove - later
	default:
		fmt.Println("Invalid choice")
	}
}

func viewMenu() {
	fmt.Println("1) View existing items  2) Update existing items")
	choice := getInput("> ")

	switch choice {
	case 1:
		// view - later
	case 2:
		// update - later
	default:
		fmt.Println("Invalid choice")
	}
}
