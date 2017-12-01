# Concurrent_Go_Tour
A tour of the concurrent features of go.

# Report and Presentation
GoLang_Report.docx -report
Golang.pptx -presentation

# Examples Of Primitives
## Goroutines
concurrent_sequential.go -A Simple loops that sleeps for the iteration its on

concurrent_one_thread.go -A loop that creates several goroutines to work concurrently

concurrent_four_threads.go -A loop that creates several goroutines and runs them in 4 threads


## Channels
channel.go -creates a simple channel and passes it with a goroutine

channelConc.go -creates two channels and reads them in a particular order

channelLoop.go -creates a channel and loops through 10 times to send through the channel


## Atomic
atomic.go -creates an int pointer and adds to it using an atomic operation

## Mutex
mutex.go -shows an example of the mutex type in go

# Examples Of Applications
These examples apply the primitives above to try and do something more interesting

## Reduction
reduction_mutex.go -creates a reduction for finding the max in a list with a mutex lock

reduction_channel.go -creates a reduction for finding themax in a list with


## WebCrawler
crawler.go -explores all links up to n websites from the starting website

crawler_concurrent.go -explores all links up to n websites by using goroutines
