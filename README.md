# throttling-sliding-window
It's a small practice using Golang, to solve an interesting scenario related to throttling/rate limiting. 

Scenario:
* The incoming requests volume is overwelmingly large. In this example we have 1000 requests, and we can only process at most 100 at a time. 
* Started 16 goroutines (threads) to process that 100 jobs in parallel. When some jobs are processed, the jobs channel can accept more requests. 
* The goroutines will write the results to the results channel. 

Small experiment:
On my machine, when I started 16 threads the total run time is less than 7 seconds. When I changed to 1 thread only, the total run time is around 1 minute 44 seconds. 
