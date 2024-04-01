package machine

import (
	"coffee_machine/inventary"
	"coffee_machine/utils"
	"fmt"
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
		fmt.Println("Check inventary")
	case 2:
		fmt.Println("Check cash regiter")
	case 3:
		fmt.Println("Brew coffee")
	case 4:
		fmt.Println("Brew coffee")
	default:
		utils.ClearScreen()
		fmt.Printf("\nInvalid choice, Please select a valid option.\n")
		selectOptionMainMenu()
	}

}

func toInventary() {
	inventary.CheckInventary()
	selectOptionInventary()
}

func selectOptionInventary() {
	fmt.Printf("\n")
	fmt.Println("=== Inventary Menu ===")
	fmt.Println("1) - Return Main Menu")
	fmt.Println("2) - Add Cup or Supplies")

	fmt.Printf("\nSelect an option: ")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		utils.ClearScreen()
		Init()
	case 2:
		fmt.Println("Add Cup or Supplies")
	default:
		utils.ClearScreen()
		fmt.Printf("\nInvalid choice, Please select a valid option.\n")
		selectOptionInventary()
	}
}
