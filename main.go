//package main
//
//import (
//	"fmt"
//	"math/rand/v2"
//	"sync"
//	"time"
//)
//
//type Order struct {
//	ID      int
//	Topping string
//}
//
//type Pizza struct {
//	OrderID int
//	Topping string
//	Chef    string
//}
//
//var toppings = []string{"pepperoni", "sausage", "veggie", "meatball", "hawaiian"}
//var chefNames = []string{"Alice", "Bob", "Charlie"}
//
//func main() {
//	orders := make(chan Order)
//	pizzas := make(chan Pizza, 8)
//
//	go func() {
//		for i := 1; i <= len(toppings); i++ {
//			orders <- Order{ID: i, Topping: toppings[i%len(toppings)]}
//		}
//		close(orders)
//	}()
//
//	var wg sync.WaitGroup
//	for _, chef := range chefNames {
//		wg.Add(1)
//		go func(c string) {
//			defer wg.Done()
//			prepare(orders, pizzas, c)
//		}(chef)
//	}
//
//	go func() {
//		wg.Wait()
//		close(pizzas)
//	}()
//
//	delivered := 0
//	timeout := time.After(3 * time.Second)
//
//Outer:
//	for {
//		select {
//		case <-timeout:
//			fmt.Printf("No sender")
//			break Outer
//		case p, ok := <-pizzas:
//			if !ok {
//				break Outer
//			}
//			fmt.Printf("Delivering Pizza.. {%d, %s, %s}\n", p.OrderID, p.Topping, p.Chef)
//			delivered++
//			break
//		}
//	}
//
//	fmt.Println("Delivered: ", delivered)
//}
//
//func prepare(orders <-chan Order, pizzas chan<- Pizza, chef string) {
//	for o := range orders {
//		fmt.Printf("Preparing pizza.. {%d, %s, %s}\n", o.ID, o.Topping, chef)
//		time.Sleep(time.Duration(200+rand.IntN(400)) * time.Millisecond)
//		pizzas <- Pizza{OrderID: o.ID, Topping: o.Topping, Chef: chef}
//	}
//}

//type Order struct {
//	ID      int
//	Topping string
//}
//
//type Pizza struct {
//	OrderID int
//	Topping string
//	Chef    string
//}
//
//func main() {
//	orders := make(chan Order)
//	pizzas := make(chan Pizza, 8)
//
//	go func() {
//		for i := 1; i <= 5; i++ {
//			orders <- Order{ID: i, Topping: toppings[i%len(toppings)]}
//		}
//		close(orders)
//	}()
//
//	var wg sync.WaitGroup
//	for _, name := range chefNames {
//		wg.Add(1)
//		go func(n string) {
//			defer wg.Done()
//			chef(n, orders, pizzas)
//		}(name)
//	}
//
//	go func() {
//		wg.Wait()
//		close(pizzas)
//	}()
//
//	// TODO: the shop closes after 3 seconds. Rewrite this loop using a
//	// `select` so we either receive a pizza OR hit the deadline. When the
//	// deadline fires, print "Shop's closed!" and break out. When pizzas is
//	// closed (no more coming), also break out.
//	delivered := 0
//	timeout := time.After(3 * time.Second)
//Outer:
//	for {
//		select {
//		case <-timeout:
//			fmt.Println("no sender")
//			break Outer
//		case p, ok := <-pizzas:
//			if !ok {
//				break Outer
//			}
//			fmt.Printf("Delivered #%d %s by %s\n", p.OrderID, p.Topping, p.Chef)
//			delivered++
//			break
//
//		}
//	}
//
//	fmt.Printf("Total delivered: %d\n", delivered)
//}
//
//func chef(name string, orders <-chan Order, pizzas chan<- Pizza) {
//	for o := range orders {
//		fmt.Printf("%s starts #%d (%s)\n", name, o.ID, o.Topping)
//		time.Sleep(time.Duration(200+rand.Intn(400)) * time.Millisecond)
//		pizzas <- Pizza{OrderID: o.ID, Topping: o.Topping, Chef: name}
//	}
//}
//
//var chefNames = []string{"Mario", "Luigi", "Peach"}
//var toppings = []string{"pepperoni", "margherita", "hawaiian", "veggie"}

//package main
//
//import "fmt"
//
//type Order struct {
//	ID      int
//	Topping string
//}
//
//var toppings = []string{"pepperoni", "mushrooms", "onions", "margherita"}
//
//func main() {
//
//	orders := make(chan Order)
//
//	go func() {
//		var order string
//		for i := 1; i <= len(toppings); i++ {
//			order = toppings[i-1]
//			fmt.Printf("Receiving order.. %s\n", order)
//			orders <- Order{ID: i, Topping: order}
//		}
//		fmt.Println("Order closed")
//		close(orders)
//	}()
//
//	// main reaches this for
//	// no order exist, so the main thread blocks and waits for the sender
//	for o := range orders {
//		fmt.Printf("Order received: %d - %s\n", o.ID, o.Topping)
//	}
//}
