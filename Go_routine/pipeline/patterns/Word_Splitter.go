package patterns

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Sentence Word Splitter

	Given a list of sentences, split each into words.
	Stages: Generator → Splitter → Printer
*/
func Word_Splitter() {
	var wg sync.WaitGroup
	generator := make(chan string)
	printer := make(chan string)
	Generator_Word_Splitter(&wg, generator)
	Splitter(generator, &wg, printer)
	Printer_Word_Splitter(&wg, printer)
	wg.Wait()
}
func Generator_Word_Splitter(wg *sync.WaitGroup, generator chan string) {

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			words := []string{"apple", "banana", "cherry", "date", "elephant", "fig", "grape"}

			rand.Seed(time.Now().UnixNano()) // Seed with current time
			randomWord := words[rand.Intn(len(words))]

			fmt.Println("Random word:", randomWord)

			generator <- randomWord
			time.Sleep(time.Second)
		}
		close(generator)
	}()

}
func Splitter(generator chan string, wg *sync.WaitGroup, printer chan string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range generator {
			for i := 0; i < len(data); i++ {
				printer <- string(data[i])
			}
		}
		close(printer)
	}()
}
func Printer_Word_Splitter(wg *sync.WaitGroup, printer chan string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range printer {
			fmt.Println(data)
		}
	}()
}
