package coffees

type Coffee struct {
	Name   string
	Milk   int
	Water  int
	Coffee int
}

var Coffees = map[string]Coffee{
	"espresso":   {Name: "Espresso", Milk: 0, Water: 30, Coffee: 7},
	"americano":  {Name: "Americano", Milk: 0, Water: 30, Coffee: 7},
	"cappuccino": {Name: "Cappuccino", Milk: 0, Water: 30, Coffee: 7},
	"latte":      {Name: "Latte", Milk: 0, Water: 30, Coffee: 7},
	"macchiato":  {Name: "Macchiato", Milk: 0, Water: 30, Coffee: 7},
}
