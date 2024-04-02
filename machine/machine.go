package machine

import (
	cashregister "coffee_machine/cash_register"
	"coffee_machine/inventary"
	"coffee_machine/inventary/cups"
	"coffee_machine/inventary/supplies"
	"coffee_machine/utils"
	"fmt"
	"os"
)

func Init() {
	mainMenu()
	selectOptionMainMenu()
}

func selectOptionMainMenu() {
	fmt.Printf("\nSelect an option: ")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		toInventary()
	case 2:
		toCashRegister()
	case 3:
		// utils.ClearScreen()
		is_brewed := startBrewCoffee()

		if is_brewed {
			var input string
			fmt.Scanln(&input)
		} else {
			os.Exit(0)
		}

	case 4:
		os.Exit(0)
	default:
		utils.ClearScreen()
		fmt.Printf("\nInvalid choice, Please select a valid option.\n")
		selectOptionMainMenu()
	}

	Init()
}

func toInventary() {
	inventary.CheckInventary()
	selectOptionInventary()
}

func selectOptionInventary() {
	fmt.Printf("\n")
	fmt.Println("=== Inventary Menu ===")
	fmt.Println("1) - Add Cup or Supplies")
	fmt.Println("2) - Return Main Menu")

	fmt.Printf("\nSelect an option: ")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		addInventary()
	case 2:
		utils.ClearScreen()
		Init()
	default:
		utils.ClearScreen()
		fmt.Printf("\nInvalid choice, Please select a valid option.\n")
		selectOptionInventary()
	}
}

func addInventary() {
	fmt.Printf("=== Add Cup or Supplies ===\n")
	fmt.Printf("\n What do you want to add?\n")
	fmt.Println("1) Cup")
	fmt.Println("2) Water")
	fmt.Println("3) Milk")
	fmt.Println("4) Coffee Beans")

	fmt.Printf("\nSelect an option: ")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		addCup()
	case 2:
		addSupply("water")
	case 3:
		addSupply("milk")
	case 4:
		addSupply("coffee_beans")
	default:
		fmt.Printf("\nInvalid choice, Please select a valid option.\n")
		addInventary()
	}

	toInventary()
}

func addCup() {
	fmt.Printf("=== Add Cup ===\n")
	fmt.Println("What size?")
	fmt.Println("1) S")
	fmt.Println("2) M")
	fmt.Println("3) L")

	fmt.Printf("\nSelect an option: ")

	var optionSize int
	fmt.Scanln(&optionSize)

	if optionSize != 1 && optionSize != 2 && optionSize != 3 {
		fmt.Printf("\nInvalid choice, Please select a valid option.\n")
		addCup()
	}

	var quantity int

	// if the user put letter don't break continue to the infiny
	for {
		fmt.Printf("\nHow many?: ")
		_, err := fmt.Scanln(&quantity)
		if err != nil {
			fmt.Printf("\nInvalid choice, Please select a valid option.\n")
		} else {
			break
		}
	}

	// this index is 0 = S, M = 1, L = 2
	index := optionSize - 1

	cups.Add(index, quantity)
}

// the keys are (coffee_beans, milk, water)
func addSupply(key string) {
	fmt.Printf("=== Add %s ===\n", key)

	var quantity int

	for {
		fmt.Printf("\nHow many?: ")
		_, err := fmt.Scanln(&quantity)
		if err != nil {
			fmt.Printf("\nInvalid choice, Please select a valid option.\n")
		} else {
			break
		}
	}

	supplies.Add(key, quantity)
}

func toCashRegister() {
	utils.ClearScreen()
	fmt.Printf("=== Check cash regiter ===\n")
	fmt.Printf("\n Money: $%.3f\n\n", cashregister.Cash)

	fmt.Println("1) - Withdraw funds")
	fmt.Println("2) - Return Main Menu")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		var quantity int

		// if the user put letter don't break continue to the infiny
		for {
			fmt.Printf("\nHow many?: ")
			_, err := fmt.Scanln(&quantity)
			if err != nil {
				fmt.Printf("\nInvalid choice, Please select a valid option.\n")
			} else if float64(quantity) > cashregister.Cash {
				fmt.Printf("\nWe don't have enough\n")
			} else {
				cashregister.Cash -= float64(quantity)
				break
			}
		}

	case 2:
		utils.ClearScreen()
		Init()
	default:
		fmt.Printf("\nInvalid choice, Please select a valid option.\n")
		toCashRegister()
	}
}
