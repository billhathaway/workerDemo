Concurrency demo
--
Shows off using workers, channels, and messages

Installation: (Assuming go installed and GOPATH, PATH adjusted)  
`go get github.com/billhathaway/workerDemo`  

# fetch default URLs with 1 worker goroutine  
`workerDemo`  

# fetch default URLs with 5 worker goroutines  
`workerDemo -n 5`  


# fetch URLs from a file and split the work among 10 worker goroutines  
`workerDemo -f urlFile.txt -n 10`  

