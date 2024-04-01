package cups

type Cup struct {
	Type       string
	Quantity   int
	Multiplier float64
}

var Cups []Cup = []Cup{
	{Type: "S", Quantity: 10, Multiplier: 0.8},
	{Type: "M", Quantity: 10, Multiplier: 1},
	{Type: "L", Quantity: 10, Multiplier: 1.2},
}

func Add(index int, quantity int) {
	Cups[index].Quantity += quantity
}

func Get(typeCup string) *Cup {
	for _, cup := range Cups {
		if cup.Type == typeCup {
			return &cup
		}
	}

	return nil
}
