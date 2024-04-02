package supplies

import "errors"

type Supply struct {
	Name     string
	Quantity int
}

var Supplies = map[string]Supply{
	"coffee_beans": {Name: "Coffee Beans", Quantity: 1000},
	"milk":         {Name: "Milk", Quantity: 100},
	"water":        {Name: "Water", Quantity: 100},
}

// add to supply
func Add(supply string, quantity int) {
	if s, ok := Supplies[supply]; ok {
		s.Quantity += quantity
		Supplies[supply] = s
	}
}

// Less quantity
func Less(supply string, quantity int) error {

	if quantity > Supplies[supply].Quantity {
		return errors.New("quantity to subtract is greater than current quantity")
	}

	if s, ok := Supplies[supply]; ok {
		s.Quantity -= quantity
		Supplies[supply] = s
	} else {
		return errors.New("index out of range")
	}

	return nil
}
