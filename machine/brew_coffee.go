package machine

import (
	cashregister "coffee_machine/cash_register"
	"coffee_machine/inventary/cups"
	"coffee_machine/inventary/supplies"
	"coffee_machine/machine/coffees"
	"coffee_machine/utils"
	"fmt"
	"os"
)

func startBrewCoffee() bool {
	var coffee coffees.Coffee = menu()
	var cup cups.Cup = sizeMenu(&coffee)

	if cup.Quantity == 0 {
		fmt.Println("We don't have enough cup")
		startBrewCoffee()
		return false
	}

	var cosumable map[string]float64 = getCosumable(&coffee, &cup)

	insufficient_supply, all_available := validator(&cosumable)

	if !all_available {
		fmt.Printf("Insuficiente %s", insufficient_supply)
		startBrewCoffee()
		return false
	}

	var is_brewed bool = brewCoffee(&cup, &cosumable)
	cashregister.Cash += float64(coffee.Coust)
	if !is_brewed {
		return false
	}

	fmt.Printf("Perfection! The coffee (%s) size %s has been brewed flawlessly.", coffee.Name, cup.Type)
	fmt.Println("Press any key for returne main menu")

	return true

}

func brewCoffee(cup *cups.Cup, cosumable *map[string]float64) bool {

	index_cup, _ := cups.GetIndex(string(cup.Type))

	err := cups.Less(index_cup, 1)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	for key, value := range *cosumable {
		if value > 0 {
			// TODO: Rounding, but it's a bad practice here since decimals are lost.
			err := supplies.Less(key, int(value))

			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
		}
	}

	return true
}

func validator(cosumable *map[string]float64) (string, bool) {
	my_supplies := supplies.Supplies

	for i, supply := range my_supplies {
		if (*cosumable)[i] > float64(supply.Quantity) {
			return supply.Name, false
		}
	}

	return "", true
}

func getCosumable(coffee *coffees.Coffee, cup *cups.Cup) map[string]float64 {
	var cosumable = make(map[string]float64)

	var milk_to_consume float64 = float64(coffee.Milk) * float64(cup.Multiplier)
	var water_to_consume float64 = float64(coffee.Water) * float64(cup.Multiplier)
	var coffee_to_consume float64 = float64(coffee.Coffee) * float64(cup.Multiplier)

	cosumable["milk"] = milk_to_consume
	cosumable["water"] = water_to_consume
	cosumable["coffee_beans"] = coffee_to_consume

	return cosumable
}

// menu displays the menu options for selecting a type of coffee and
// returns the selected option.
// It returns an integer representing the option selected by the user.
func menu() coffees.Coffee {
	var optionSize int

	for {
		showMenu()
		fmt.Scanln(&optionSize)

		if optionSize >= 1 && optionSize < len(coffees.Coffees) {
			break
		} else {
			fmt.Printf("\nInvalid choice, Please select a valid option.\n")
		}
	}

	return coffees.Coffees[optionSize]
}

// showMenu displays the menu options for available types of coffee.
// It takes no arguments and doesn't return any value.
func showMenu() {
	utils.ClearScreen()
	fmt.Println("=== Brew Menu ===")
	fmt.Printf("\nWhat coffee do you want?\n\n")

	for i, coffee := range coffees.Coffees {
		fmt.Printf("%d) - %s\n", i, coffee.Name)
	}

	fmt.Printf("\nSelect an option: ")
}

// sizeMenu displays the menu options for selecting a size (small, medium, large) and returns the selected option.
// It returns an integer representing the option selected by the user.
func sizeMenu(coffee *coffees.Coffee) cups.Cup {
	var option int

	for {
		showSizeMenu(coffee)

		fmt.Scanln(&option)

		if option >= 1 && option <= 3 {
			break
		} else {
			fmt.Printf("\nInvalid choice, Please select a valid option.\n")
		}
	}

	// TODO: cups.Cups in the future, we'll need to changer de first index
	// need to be 1 not 0
	return cups.Cups[option-1]
}

// showSizeMenu displays the menu options for selecting a size (small, medium, large).
// It takes no arguments 3and doesn't return any value.
func showSizeMenu(coffee *coffees.Coffee) {
	utils.ClearScreen()

	fmt.Printf("\nWhat size would you like for your %s (small, medium, large)?\n", coffee.Name)
	fmt.Printf("1) Small\n")
	fmt.Printf("2) Medium\n")
	fmt.Printf("3) Large\n")

	fmt.Printf("\nSelect an option: ")
}
