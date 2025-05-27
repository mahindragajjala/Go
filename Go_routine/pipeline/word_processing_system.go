package pipeline

import (
	"fmt"
	"strings"
	"time"
)

func Word_processing_system() {
	words := []string{"abc", "defg", "hijkl", "mnop", "qrs", "tuvw", "xyz"}
	sentences := Generate(words)
	split(sentences)
	time.Sleep(2 * time.Second)
}

/*
	func generate(sentences []string) <-chan string {
	    out := make(chan string)
	    go func() {
	        for _, sentence := range sentences {
	            out <- sentence
	        }
	        close(out)
	    }()
	    return out
	}
*/
/*
Explanation:

A goroutine sends each sentence to the channel.

It closes the channel once all data is sent. */
func Generate(sentence []string) (sent chan string) {
	sent = make(chan string)
	go func() {
		for _, value := range sentence {
			sent <- value
		}
		close(sent)
	}()
	return sent
}

/*
Explanation:
For each sentence from the input channel, split it into words.
Send each word to the next stage.

	func split(in <-chan string) <-chan string {
	    out := make(chan string)
	    go func() {
	        for sentence := range in {
	            words := strings.Fields(sentence)
	            for _, word := range words {
	                out <- word
	            }
	        }
	        close(out)
	    }()
	    return out
	}
*/
func split(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for sentence := range in {
			words := strings.Fields(sentence)
			for _, word := range words {
				fmt.Println(word)
				out <- word
			}
		}
		close(out)
	}()
	return out
}
