Concurrency demo
--
Shows off using workers, channels, and messages

Installation: (Assuming go installed and GOPATH, PATH adjusted)  
    
    go get github.com/billhathaway/workerDemo  

Fetch default URLs with 1 worker goroutine  

    workerDemo
    2014/10/31 21:11:47 {http://google.com true}  
    2014/10/31 21:11:49 {http://cnn.com true}  
    2014/10/31 21:11:49 {http://cposc.org true}  
    2014/10/31 21:11:49 {http://cloudflare.com true}  
    2014/10/31 21:11:51 {http://reddit.com true}  
    2014/10/31 21:11:51 Finished in 3.580 seconds  


Fetch default URLs with 5 worker goroutines  

    workerDemo -n 5  
    2014/10/31 21:11:55 {http://cposc.org true}  
    2014/10/31 21:11:55 {http://reddit.com true}  
    2014/10/31 21:11:55 {http://google.com true}  
    2014/10/31 21:11:55 {http://cnn.com true}  
    2014/10/31 21:11:56 {http://cloudflare.com true}  
    2014/10/31 21:11:56 Finished in 0.517 second  

Fetch URLs from a file and split the work among 10 worker goroutines  

    workerDemo -f urlFile.txt -n 10

