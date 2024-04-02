package cups

import "errors"

// Cup represents a cup with its type, quantity, and multiplier.
type Cup struct {
	Type       string  // Type of the cup (e.g., "S", "M", "L")
	Quantity   int     // Quantity of the cup
	Multiplier float64 // Multiplier for the cup
}

// Cups represents a slice of Cup.
var Cups []Cup = []Cup{
	{Type: "S", Quantity: 10, Multiplier: 0.8},
	{Type: "M", Quantity: 10, Multiplier: 1},
	{Type: "L", Quantity: 10, Multiplier: 1.2},
}

// Add increases the quantity of the cup at the specified index.
func Add(index int, quantity int) {
	Cups[index].Quantity += quantity
}

// Less subtracts quantity from the cup at the specified index.
// It returns an error if the quantity to subtract is greater than the current quantity or if the index is out of range.
func Less(index int, quantity int) error {

	if quantity > Cups[index].Quantity {
		return errors.New("quantity to subtract is greater than current quantity")
	}

	if index < 0 || index >= len(Cups) {
		return errors.New("index out of range")
	}

	Cups[index].Quantity -= quantity

	return nil
}

// GetIndex returns the index of the cup with the specified type.
// It also returns a boolean indicating whether the cup was found or not.
func GetIndex(typeCup string) (int, bool) {
	for i, cup := range Cups {
		if cup.Type == typeCup {
			return i, true
		}
	}
	return 0, false
}
