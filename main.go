// Showing the use of goroutines and channels
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	defaultURLs = []string{
		"http://google.com",
		"http://cnn.com",
		"http://cposc.org",
		"http://cloudflare.com",
		"http://reddit.com",
	}
)

// urlStatus contains the URL and whether it was reachable/return 200
type urlStatus struct {
	url     string
	working bool
}

// worker reads URL strings from the work chan until it is closed
// each URL is tried, and if it is reachable with a 200 status code
// a urlStatus is returned indicating working, otherwise a urlStatus
// is returned with working false
func worker(work chan string, results chan urlStatus, wg *sync.WaitGroup) {
	// show off defer, could also just use wg.Done() at end of func
	defer wg.Done()
	for url := range work {
		status := urlStatus{url: url}
		response, err := http.Get(url)
		if err != nil {
			results <- status
			continue
		}
		// drain data from the request
		io.Copy(ioutil.Discard, response.Body)
		status.working = (response.StatusCode == 200)
		results <- status
	}
}

// aggregator prints the results it receives
// until it receives a value on the quit channel
// upon which it causes the program to exit
func aggregator(results chan urlStatus, quit chan time.Time) {
	for {
		select {
		case result := <-results:
			fmt.Printf("%s working=%v\n", result.url, result.working)
		case startTime := <-quit:
			fmt.Printf("Finished in %.3f seconds\n", time.Since(startTime).Seconds())
			os.Exit(0)
		}
	}
}

// controller creates 3 channels, and spawns workers and a tabulator
// the URLs are sent to the workers and then the work channel is closed
// once all the workers are finished, the quit channel is sent a message
// so the tabulator knows everything is completed
func controller(urls []string, workers int) {
	wg := &sync.WaitGroup{}

	results := make(chan urlStatus)
	work := make(chan string)
	quit := make(chan time.Time)

	go aggregator(results, quit)
	startTime := time.Now()
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker(work, results, wg)
	}
	// send all the data
	for _, url := range urls {
		work <- url
	}
	close(work)
	wg.Wait()
	quit <- startTime
	select {}
}

// loadFile reads a list of URLs from a file and returns them
func loadFile(file string) []string {
	fd, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fd.Close()

	var urls []string
	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	return urls
}

func main() {
	workers := flag.Int("n", 1, "number of workers")
	urlFile := flag.String("f", "", "file containing URLs")
	flag.Parse()
	var urls []string
	if *urlFile != "" {
		urls = loadFile(*urlFile)
	} else {
		urls = defaultURLs
	}
	controller(urls, *workers)
}
