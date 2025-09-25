package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func FetchTitle(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error reading body of %s: %v", url, err)
		return
	}

	titleStart := strings.Index(string(body), "<title>")
	if titleStart == -1 {
		ch <- fmt.Sprintf("No title found for %s", url)
		return
	}

	titleEnd := strings.Index(string(body), "</title>")
	if titleEnd == -1 {
		ch <- fmt.Sprintf("No title found for %s", url)
		return
	}

	title := string(body[titleStart+len("<title>") : titleEnd])

	ch <- fmt.Sprintf("URL: %s -> Title: %s", url, title)

}

func main() {
	urls := []string{
		"https://go.dev/",
		"https://golang.org/",
		"https://www.google.com/",
		"https://www.youtube.com/",
	}

	ch := make(chan string)

	for _, url := range urls {
		go FetchTitle(url, ch)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}
}
