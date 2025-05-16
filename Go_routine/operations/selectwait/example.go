package selectwait

import (
	"fmt"
	"math/rand"
	"time"
)

func temperatureSensor(ch chan string) {
	for {
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
		ch <- fmt.Sprintf("Temperature: %d°C", rand.Intn(35)+10)
	}
}

func humiditySensor(ch chan string) {
	for {
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
		ch <- fmt.Sprintf("Humidity: %d%%", rand.Intn(50)+30)
	}
}

func Wheather_Data() {
	rand.Seed(time.Now().UnixNano())

	tempChan := make(chan string)
	humidChan := make(chan string)

	// Start sensors
	go temperatureSensor(tempChan)
	go humiditySensor(humidChan)

	for {
		select {
		case temp := <-tempChan:
			fmt.Println("[Sensor] Received:", temp)
		case humid := <-humidChan:
			fmt.Println("[Sensor] Received:", humid)
		case <-time.After(3 * time.Second):
			fmt.Println("[Warning] No data received in the last 3 seconds!")
		}
	}
}
