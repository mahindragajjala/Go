package bufferedchannel

import "fmt"

func Worker_Pool(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "processing job", j)
		results <- j * 2
	}
}
