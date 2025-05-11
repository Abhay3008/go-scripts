package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Printf("Hey! welcome to booking system!\n")
	const numberoftickets int = 123
	var price int
	var remainingtickets int
	var ticketspurchased int
	remainingtickets = 76
	price = 199
	var user string
	var bookings = make(map[string]int)

	for remainingtickets > 0 {
		fmt.Printf("We have %v tickets out of which %v tickets are left, grab yours now, the price is %vrs only\n", numberoftickets, remainingtickets, price)
		fmt.Printf("Do you want a Ticket\nTell me yoour username:\n")
		fmt.Scan(&user)
		fmt.Printf("How many tickets you want:\n")
		fmt.Scan(&ticketspurchased)
		bookings[user] = ticketspurchased
		if !checktickets(ticketspurchased, remainingtickets) {
			fmt.Printf("Sorry! we have only %v tickets left\n", remainingtickets)
			continue
		}
		fmt.Printf("user %v purchased %v tickets\n", user, ticketspurchased)
		wg.Add(1)
		go sendingNotification(user, ticketspurchased)
		remainingtickets -= ticketspurchased
	}
	fmt.Printf("Tickets have been sold to these users: %v\n", bookings)
	wg.Wait()

}

func checktickets(purchased int, remaining int) bool {

	return remaining >= purchased

}

func sendingNotification(user string, tickets int) {
	time.Sleep(10 * time.Second)
	fmt.Printf("*******\n")
	fmt.Printf("Notifications has been sent to user %v for the purchase of %v tickets\n", user, tickets)
	fmt.Printf("*******\n")
	wg.Done()
}
