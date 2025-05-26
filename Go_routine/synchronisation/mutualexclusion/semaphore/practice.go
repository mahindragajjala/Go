package semaphore

import (
	"fmt"
	"sync"
	"time"
)

var Critical_section int

func Prac() {
	var channel = make(chan struct{}, 2)
	for i := 0; i < 5; i++ {
		go Worker(channel)
	}
}
func Worker(ch chan struct{}) {
	ch <- struct{}{}
	Critical_section++
	<-ch
}

//Parking semaphone

func Parking_slot(ch chan struct{}, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Car %v is parking...\n", i)
	ch <- struct{}{} // Occupy slot
}

func Displacing(ch chan struct{}, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Car %v is displacing...\n", i)
	<-ch // Leave slot
}

func Parking() {
	ParkingSlots := make(chan struct{}, 1) // Only 1 parking slot
	NumberOfCars := 6
	var wg sync.WaitGroup

	for i := 1; i <= NumberOfCars; i++ {
		wg.Add(2) // One for Parking_slot, one for Displacing

		carID := i // Capture the loop variable
		go func(id int) {
			Parking_slot(ParkingSlots, id, &wg)
			time.Sleep(2 * time.Second) // Simulate parking duration
			Displacing(ParkingSlots, id, &wg)
		}(carID)
	}

	wg.Wait()
	fmt.Println("All cars have parked and displaced.")
}
