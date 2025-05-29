package semaphore

import (
	"fmt"
	"time"
)

/*
Task:
Write a Go program where multiple goroutines increment a shared counter using a "binary semaphore" to ensure mutual exclusion.

Goal:
Prevent race conditions using a buffered channel of size 1.
*/
func Counter() {
	var count int
	var channel = make(chan struct{}, 1)
	go func() {
		channel <- struct{}{}
		count++
		<-channel
	}()
	go func() {
		channel <- struct{}{}
		count++
		<-channel
	}()
	go func() {
		channel <- struct{}{}
		count++
		<-channel
	}()
	go func() {
		channel <- struct{}{}
		count++
		<-channel
	}()
	go func() {
		channel <- struct{}{}
		count++
		<-channel
	}()
	go func() {
		channel <- struct{}{}
		count++
		<-channel
	}()
	go func() {
		channel <- struct{}{}
		count++
		<-channel
	}()
	go func() {
		channel <- struct{}{}
		count++
		<-channel
	}()
	go func() {
		channel <- struct{}{}
		count++
		<-channel
	}()
	go func() {
		channel <- struct{}{}
		count++
		<-channel
	}()
	go func() {
		channel <- struct{}{}
		count++
		<-channel
	}()
	go func() {
		channel <- struct{}{}
		count++
		<-channel
	}()
	go func() {
		channel <- struct{}{}
		count++
		<-channel
	}()
	go func() {
		channel <- struct{}{}
		count++
		<-channel
	}()
	go func() {
		channel <- struct{}{}
		count++
		<-channel
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("count:", count)
}
