package main

import "fmt"

type Order struct {
	ID      int
	Topping string
}

var toppings = []string{"pepperoni", "mushrooms", "onions", "margherita"}

func main() {

	orders := make(chan Order)

	go func() {
		var order string
		for i := 1; i <= len(toppings); i++ {
			order = toppings[i-1]
			fmt.Printf("Receiving order.. %s\n", order)
			orders <- Order{ID: i, Topping: order}
		}
		close(orders)
		fmt.Println("Order closed")
	}()

	for o := range orders {
		fmt.Printf("Order received: %d - %s\n", o.ID, o.Topping)
	}
}
