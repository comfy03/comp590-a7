package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	numChairs    int
	numCustomers int
	barberTime   int
	numBarbers   int
	mutex        sync.Mutex
	waitingRoom  chan int
	wg           sync.WaitGroup
)

func barber(barberID int) {
	for {
		mutex.Lock()
		if len(waitingRoom) == 0 {
			fmt.Printf("Barber %d is sleeping as no customers are waiting.\n\n", barberID)
			mutex.Unlock()
			time.Sleep(2 * time.Second)
		} else {
			customerID := <-waitingRoom
			mutex.Unlock()
			fmt.Printf("Barber %d starts cutting hair for customer %d.\n\n", barberID, customerID)
			time.Sleep(time.Duration(barberTime) * time.Second)
			fmt.Printf("Barber %d finished cutting hair for customer %d.\n\n", barberID, customerID)
			wg.Done()
		}
	}
}

func customer(customerID int) {
	mutex.Lock()
	if len(waitingRoom) == numChairs {
		fmt.Printf("Customer %d leaves as the waiting room is full.\n\n", customerID)
		mutex.Unlock()
		wg.Done()
		return
	}

	fmt.Printf("Customer %d takes a seat in the waiting room.\n\n", customerID)
	waitingRoom <- customerID
	mutex.Unlock()
}

func main() {
	fmt.Println("Enter the number of chairs in the waiting room:")
	fmt.Scan(&numChairs)
	fmt.Println("Enter the number of customers:")
	fmt.Scan(&numCustomers)
	fmt.Println("Enter the time (in seconds) for a barber to cut hair:")
	fmt.Scan(&barberTime)
	fmt.Println("Enter the number of barbers:")
	fmt.Scan(&numBarbers)

	waitingRoom = make(chan int, numChairs)

	for i := 1; i <= numBarbers; i++ {
		go barber(i)
	}

	for i := 1; i <= numCustomers; i++ {
		wg.Add(1)
		go customer(i)
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
	}

	wg.Wait()
	fmt.Println("All customers have been served or left.\n")
}
