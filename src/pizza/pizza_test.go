package pizza

import (
	"testing"
)

func TestPizza(t *testing.T) {
	pizza := NewPizzaBuilder().Size("small").AddCheese().AddMushroom().AddTomatoSauce().Build()
	t.Log(pizza)
	if pizza.size != "small" {
		t.Error("Size is not correct")
	}

	if pizza.cheese != true {
		t.Error("Cheese wrongly added")
	}

	if pizza.basil != false {
		t.Error("Basil wrongly added")
	}

	pizza2 := NewPizzaBuilder().Size("large").AddProsciutto().AddSpinach().AddCheese().Build()
	t.Log(pizza2)
	if pizza2.size != "large" {
		t.Error("Size is not correct")
	}

	if pizza2.cheese != true {
		t.Error("Cheese wrongly added")
	}

	if pizza2.prosciutto != true {
		t.Error("Prosciutto wrongly added")
	}
}
