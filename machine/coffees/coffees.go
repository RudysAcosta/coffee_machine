package coffees

type Coffee struct {
	Name   string
	Milk   int
	Water  int
	Coffee int
	Coust  float64
}

var Coffees = map[int]Coffee{
	1: {Name: "Espresso", Milk: 20, Water: 30, Coffee: 7, Coust: 2.3},
	2: {Name: "Americano", Milk: 30, Water: 30, Coffee: 7, Coust: 3.3},
	3: {Name: "Cappuccino", Milk: 10, Water: 30, Coffee: 7, Coust: 5.3},
	4: {Name: "Latte", Milk: 10, Water: 30, Coffee: 7, Coust: 6.3},
	5: {Name: "Macchiato", Milk: 12, Water: 30, Coffee: 7, Coust: 7.3},
}
