package problems

import "fmt"

//Define a simple struct Book with fields: Title, Author, Price. Write a function that prints a book's details.

type Book struct {
	Title  string
	Author string
	Price  int
}

func Book_details() {
	book1 := Book{Title: "Science", Author: "eisten", Price: 2555}
	fmt.Println(book1)
}

//Define an empty struct and create a map[string]struct{} for storing visited URLs.
// Add and check if "https://golang.org" is present.

func Storing_visited_URLs() {
	// Define a map with string keys and empty struct values
	visited := make(map[string]struct{})

	// Add URLs to the map
	visited["https://golang.org"] = struct{}{}
	visited["https://openai.com"] = struct{}{}
	visited["https://example.com"] = struct{}{}

	// URL to check
	urlToCheck := "https://golang.org"

	// Check if the URL is present in the map
	if _, exists := visited[urlToCheck]; exists {
		fmt.Printf("URL '%s' has been visited.\n", urlToCheck)
	} else {
		fmt.Printf("URL '%s' has NOT been visited.\n", urlToCheck)
	}

}
