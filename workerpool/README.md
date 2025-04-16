Case Study For Worker Pool:

Shopify increased throughput by 170% by introducing a well-tuned Go worker pool in their Server Pixels service. 

At Shopify, the system was processing over a billion events daily, but performance suffered due to an unbounded number of goroutines spawned during spikes. 
This overloaded the system and caused a lag in processing customer data especially problematic during high-traffic events like Black Friday. 

Their solution?
Implementing a worker pool to limit concurrent goroutines, control resource usage, and stabilize the system under pressure. 

The result: throughput jumped from 7.75K to 21K events/sec per pod. 


For CPU-bound tasks, matching the number of workers to GOMAXPROCS is usually a good baseline.
For I/O-heavy work, you might scale the pool higher since tasks spend time waiting. 
Link for shopify original post: https://shopify.engineering/leveraging-go-worker-pools

Below is a visual that illustrates how a Go Worker Pool processes tasks efficiently using a queue and a fixed number of workers. 

![worker_pool](https://github.com/user-attachments/assets/4d313561-5dcf-4ab8-abbb-c0e6b665802f)
