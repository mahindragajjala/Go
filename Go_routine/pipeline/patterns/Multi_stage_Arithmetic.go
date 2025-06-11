package patterns

import (
	"fmt"
	"sync"
	"time"
)

func Multi_stage_Arithmetic() {
	var wga sync.WaitGroup
	generatordata := make(chan int)
	squaredata := make(chan int)
	add10data := make(chan int)
	halvedata := make(chan int)
	Multi_stage_Arithmetic_Generator(generatordata, &wga)
	Square(generatordata, &wga, squaredata)
	Add10(&wga, squaredata, add10data)
	HalveStage(&wga, add10data, halvedata)
	Printerstage(&wga, halvedata)
	wga.Wait()
}
func Multi_stage_Arithmetic_Generator(generatordata chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			generatordata <- i
			time.Sleep(time.Second)
		}
		close(generatordata)
	}()

}
func Square(generatordata chan int, wg *sync.WaitGroup, squaredata chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range generatordata {
			v = v * v
			squaredata <- v
		}
		close(squaredata)
	}()
}
func Add10(wg *sync.WaitGroup, squaredata chan int, add10data chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range squaredata {
			v = v + 10
			add10data <- v
		}
		close(add10data)
	}()
}
func HalveStage(wg *sync.WaitGroup, add10data chan int, halvedata chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range add10data {
			v = v / 2
			halvedata <- v
		}
		close(halvedata)
	}()
}
func Printerstage(wg *sync.WaitGroup, halvedata chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range halvedata {
			fmt.Println(v)
		}
	}()
}
