package supplies

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
func Less(supply string, quantity int) {
	if s, ok := Supplies[supply]; ok {
		s.Quantity -= quantity
		Supplies[supply] = s
	}
}
