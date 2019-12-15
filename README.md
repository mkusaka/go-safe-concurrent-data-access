# go-safe-concurrent-data-access

source code original: https://eli.thegreenplace.net/2019/on-concurrency-in-go-http-servers/
# command
## tool install
brew install homebrew/apache/ab

## bench
data race web server
```bash
❯ go run data_race.go
❯ ab -n 20000 -c 200 "127.0.0.1:8000/inc?name=i"
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)

Test aborted after 10 failures

apr_socket_connect(): Connection reset by peer (54)
Total of 59 requests completed
```

mutex web server

```bash
❯ go run mutex.go
❯ ab -n 20000 -c 200 "127.0.0.1:8000/inc?name=i"
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 2000 requests
Completed 4000 requests
Completed 6000 requests
Completed 8000 requests
Completed 10000 requests
Completed 12000 requests
Completed 14000 requests
Completed 16000 requests
Completed 18000 requests
Completed 20000 requests
Finished 20000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8000

Document Path:          /inc?name=i
Document Length:        3 bytes

Concurrency Level:      200
Time taken for tests:   125.028 seconds
Complete requests:      20000
Failed requests:        0
Total transferred:      2380000 bytes
HTML transferred:       60000 bytes
Requests per second:    159.96 [#/sec] (mean)
Time per request:       1250.283 [ms] (mean)
Time per request:       6.251 [ms] (mean, across all concurrent requests)
Transfer rate:          18.59 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  417 752.9    323   10541
Processing:     2  826 773.4    703   10768
Waiting:        1  631 696.9    540   10472
Total:          6 1243 1080.7   1084   11416

Percentage of the requests served within a certain time (ms)
  50%   1084
  66%   1233
  75%   1411
  80%   1511
  90%   1815
  95%   1974
  98%   2898
  99%   3473
 100%  11416 (longest request)
```

```bash
❯ go run channel.go
❯ ab -n 20000 -c 200 "127.0.0.1:8000/inc?name=i"
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 2000 requests
Completed 4000 requests
Completed 6000 requests
Completed 8000 requests
Completed 10000 requests
Completed 12000 requests
Completed 14000 requests
Completed 16000 requests
Completed 18000 requests
Completed 20000 requests
Finished 20000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8000

Document Path:          /inc?name=i
Document Length:        3 bytes

Concurrency Level:      200
Time taken for tests:   128.576 seconds
Complete requests:      20000
Failed requests:        0
Total transferred:      2380000 bytes
HTML transferred:       60000 bytes
Requests per second:    155.55 [#/sec] (mean)
Time per request:       1285.756 [ms] (mean)
Time per request:       6.429 [ms] (mean, across all concurrent requests)
Transfer rate:          18.08 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  403 684.3    314    8898
Processing:     0  875 746.5    779    9092
Waiting:        0  669 661.5    593    8910
Total:          1 1279 995.8   1131    9640

Percentage of the requests served within a certain time (ms)
  50%   1131
  66%   1277
  75%   1395
  80%   1487
  90%   1671
  95%   1816
  98%   2592
  99%   8879
 100%   9640 (longest request)
```

fasthttp w/ channel
```bash
❯ go run channel-fasthttp.go
❯ ab -n 20000 -c 200 "127.0.0.1:8000/inc?name=i"
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 2000 requests
Completed 4000 requests
Completed 6000 requests
Completed 8000 requests
Completed 10000 requests
Completed 12000 requests
Completed 14000 requests
Completed 16000 requests
Completed 18000 requests
Completed 20000 requests
Finished 20000 requests


Server Software:        fasthttp
Server Hostname:        127.0.0.1
Server Port:            8000

Document Path:          /inc?name=i
Document Length:        3 bytes

Concurrency Level:      200
Time taken for tests:   141.398 seconds
Complete requests:      20000
Failed requests:        0
Total transferred:      3120000 bytes
HTML transferred:       60000 bytes
Requests per second:    141.44 [#/sec] (mean)
Time per request:       1413.978 [ms] (mean)
Time per request:       7.070 [ms] (mean, across all concurrent requests)
Transfer rate:          21.55 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  152 171.7     92    1396
Processing:     2 1258 530.3   1207    4315
Waiting:        1 1164 532.7   1114    4307
Total:         34 1409 489.9   1342    4316

Percentage of the requests served within a certain time (ms)
  50%   1342
  66%   1493
  75%   1592
  80%   1661
  90%   1894
  95%   2415
  98%   2683
  99%   3006
 100%   4316 (longest request)
```
optional, fasthttp w/ channel
```bash
❯ ab -n 200000 -c 2000 "127.0.0.1:8000/inc?name=i"
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 20000 requests
apr_socket_recv: Connection reset by peer (54)
Total of 37452 requests completed
```

# ref

https://eli.thegreenplace.net/2019/on-concurrency-in-go-http-servers/
