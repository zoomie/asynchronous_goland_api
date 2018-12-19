# Go API

All the action is in the app.go 

Change of plan, I still want it to be async but I also want it to be a Docker image built from scratch.

Done: Add Redis database next and see how that integrates with scratch, I am probably use docker compose to join them.

```
ab -n 10000 -c 20 "http://localhost:5000/SetValue?key=key1&value=value1"
```
```
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            5000

Document Path:          /SetValue?key=key1&value=value1
Document Length:        4 bytes

Concurrency Level:      20
Time taken for tests:   3.137 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1200000 bytes
HTML transferred:       40000 bytes
Requests per second:    3188.13 [#/sec] (mean)
Time per request:       6.273 [ms] (mean)
Time per request:       0.314 [ms] (mean, across all concurrent requests)
Transfer rate:          373.61 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   7.8      2     553
Processing:     0    4  22.7      3     553
Waiting:        0    3  20.6      2     551
Total:          1    6  24.0      5     555

Percentage of the requests served within a certain time (ms)
  50%      5
  66%      5
  75%      6
  80%      6
  90%      7
  95%      7
  98%      8
  99%      8
 100%    555 (longest request)
```
