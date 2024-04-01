package inventary

import (
	"coffee_machine/inventary/cups"
	"coffee_machine/inventary/supplies"
	"coffee_machine/utils"
	"fmt"
)

func CheckInventary() {
	utils.ClearScreen()

	fmt.Println("=== Inventary ===")
	fmt.Printf("\n")

	fmt.Println("Cups")
	for _, cup := range cups.Cups {
		fmt.Printf("%s : %d \n", cup.Type, cup.Quantity)
	}

	fmt.Printf("\nSupplies\n")

	for _, supply := range supplies.Supplies {
		fmt.Printf("%s : %d \n", supply.Name, supply.Quantity)
	}
}
