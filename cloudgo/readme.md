框架选择：gorilla/mux+codegangsta/negroni
原因：

curl 请求：
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 9090 (#0)
> GET /hello/testuser HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.62.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=UTF-8
< Date: Thu, 15 Nov 2018 19:09:04 GMT
< Content-Length: 31
<
{
  "Test": "Hello testuser"
}
* Connection #0 to host localhost left intact

ab testresult:

ab -n 1000 -c 100 http://localhost:9090/hello/your
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            9090

Document Path:          /hello/your
Document Length:        27 bytes

Concurrency Level:      100
Time taken for tests:   0.756 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      150000 bytes
HTML transferred:       27000 bytes
Requests per second:    1322.73 [#/sec] (mean)
Time per request:       75.601 [ms] (mean)
Time per request:       0.756 [ms] (mean, across all concurrent requests)
Transfer rate:          193.76 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0       1
Processing:     1   72  15.8     73     168
Waiting:        0   72  15.9     73     167
Total:          1   72  15.8     73     168

Percentage of the requests served within a certain time (ms)
  50%     73
  66%     75
  75%     77
  80%     79
  90%     84
  95%     92
  98%     96
  99%     97
 100%    168 (longest request)

 参数解析：
 -n:    一共发多少个请求
 -c:    一次性最多同时发多少个请求：