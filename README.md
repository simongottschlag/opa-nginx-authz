# opa-nginx-authz
Lab to test opa authz for nginx

## k6 load test

Start all services: `./run.sh`

Restart the above before each test!

### OPA

Run load test: `k6 run loadtest-opa.js`

```shell
running (25.0s), 00/30 VUs, 223138 complete and 0 interrupted iterations
default ✓ [======================================] 00/30 VUs  25s

     ✓ status is 200
     ✓ result is true

     check_failure_rate.............: 0.00%   ✓ 0      ✗ 223138
     checks.........................: 100.00% ✓ 446276 ✗ 0     
     data_received..................: 27 MB   1.1 MB/s
     data_sent......................: 39 MB   1.6 MB/s
     http_req_blocked...............: avg=2.39µs  min=686ns    med=1.34µs max=8.8ms   p(90)=1.95µs  p(95)=2.47µs 
     http_req_connecting............: avg=73ns    min=0s       med=0s     max=1.99ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=1.77ms  min=126.15µs med=1.06ms max=39.54ms p(90)=4.14ms  p(95)=5.78ms 
       { expected_response:true }...: avg=1.77ms  min=126.15µs med=1.06ms max=39.54ms p(90)=4.14ms  p(95)=5.78ms 
     http_req_failed................: 0.00%   ✓ 0      ✗ 223138
     http_req_receiving.............: avg=29.73µs min=7.66µs   med=14.1µs max=19.58ms p(90)=24.68µs p(95)=38.19µs
     http_req_sending...............: avg=14.66µs min=4.75µs   med=7.69µs max=27.33ms p(90)=12.29µs p(95)=16.67µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s     max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=1.73ms  min=109.21µs med=1.03ms max=39.52ms p(90)=4.07ms  p(95)=5.69ms 
     http_reqs......................: 223138  8925.133978/s
     iteration_duration.............: avg=1.9ms   min=188.17µs med=1.18ms max=39.63ms p(90)=4.34ms  p(95)=6.03ms 
     iterations.....................: 223138  8925.133978/s
     vus............................: 1       min=1    max=30  
     vus_max........................: 30      min=30   max=30 
```

### opa-nginx-external-auth proxy

Run load test: `k6 run loadtest-go-proxy.js`

```shell
running (25.0s), 00/30 VUs, 156276 complete and 0 interrupted iterations
default ✓ [======================================] 00/30 VUs  25s

     ✓ status is 200

     check_failure_rate.............: 0.00%   ✓ 0      ✗ 156276
     checks.........................: 100.00% ✓ 156276 ✗ 0     
     data_received..................: 12 MB   469 kB/s
     data_sent......................: 18 MB   706 kB/s
     http_req_blocked...............: avg=1.92µs  min=645ns    med=1.38µs  max=6.44ms   p(90)=1.74µs  p(95)=2.25µs 
     http_req_connecting............: avg=59ns    min=0s       med=0s      max=900.18µs p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=2.65ms  min=214.85µs med=1.9ms   max=31.37ms  p(90)=5.62ms  p(95)=7.56ms 
       { expected_response:true }...: avg=2.65ms  min=214.85µs med=1.9ms   max=31.37ms  p(90)=5.62ms  p(95)=7.56ms 
     http_req_failed................: 0.00%   ✓ 0      ✗ 156276
     http_req_receiving.............: avg=18.86µs min=6µs      med=13.46µs max=12.78ms  p(90)=19.91µs p(95)=23.22µs
     http_req_sending...............: avg=10.66µs min=3.91µs   med=6.98µs  max=21.41ms  p(90)=9.97µs  p(95)=12.81µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=2.62ms  min=200.82µs med=1.88ms  max=29.53ms  p(90)=5.59ms  p(95)=7.51ms 
     http_reqs......................: 156276  6250.694014/s
     iteration_duration.............: avg=2.73ms  min=261.51µs med=1.97ms  max=31.47ms  p(90)=5.73ms  p(95)=7.67ms 
     iterations.....................: 156276  6250.694014/s
     vus............................: 1       min=1    max=30  
     vus_max........................: 30      min=30   max=30 
```

### opa-nginx-external-auth rego

Run load test: `k6 run loadtest-go-rego.js`

```shell
running (25.0s), 00/30 VUs, 306568 complete and 0 interrupted iterations
default ✓ [======================================] 00/30 VUs  25s

     ✓ status is 200

     check_failure_rate.............: 0.00%   ✓ 0      ✗ 306568
     checks.........................: 100.00% ✓ 306568 ✗ 0     
     data_received..................: 23 MB   920 kB/s
     data_sent......................: 34 MB   1.4 MB/s
     http_req_blocked...............: avg=2.31µs  min=632ns    med=1.17µs   max=14.95ms p(90)=1.75µs p(95)=2.35µs 
     http_req_connecting............: avg=41ns    min=0s       med=0s       max=3.04ms  p(90)=0s     p(95)=0s     
     http_req_duration..............: avg=1.28ms  min=105.45µs med=783.69µs max=32.76ms p(90)=2.88ms p(95)=4.09ms 
       { expected_response:true }...: avg=1.28ms  min=105.45µs med=783.69µs max=32.76ms p(90)=2.88ms p(95)=4.09ms 
     http_req_failed................: 0.00%   ✓ 0      ✗ 306568
     http_req_receiving.............: avg=21.28µs min=5.61µs   med=9.78µs   max=19.53ms p(90)=17.2µs p(95)=29.57µs
     http_req_sending...............: avg=12.09µs min=3.84µs   med=5.94µs   max=25.31ms p(90)=9.63µs p(95)=14.75µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s       max=0s      p(90)=0s     p(95)=0s     
     http_req_waiting...............: avg=1.25ms  min=92.49µs  med=758.31µs max=32.66ms p(90)=2.83ms p(95)=4.03ms 
     http_reqs......................: 306568  12262.12259/s
     iteration_duration.............: avg=1.38ms  min=150.87µs med=863.46µs max=32.83ms p(90)=3.03ms p(95)=4.3ms  
     iterations.....................: 306568  12262.12259/s
     vus............................: 1       min=1    max=29  
     vus_max........................: 30      min=30   max=30  
```

### nginx public

Run load test: `k6 run loadtest-nginx-public.js`

```shell
running (25.0s), 00/30 VUs, 406248 complete and 0 interrupted iterations
default   [======================================] 00/30 VUs  25s

     ✓ status is 200

     check_failure_rate.............: 0.00%   ✓ 0      ✗ 406248
     checks.........................: 100.00% ✓ 406248 ✗ 0     
     data_received..................: 345 MB  14 MB/s
     data_sent......................: 35 MB   1.4 MB/s
     http_req_blocked...............: avg=2.86µs   min=647ns    med=1.18µs   max=9.38ms  p(90)=1.86µs  p(95)=2.54µs 
     http_req_connecting............: avg=252ns    min=0s       med=0s       max=9.29ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=938.23µs min=91.19µs  med=541.67µs max=51.92ms p(90)=2.03ms  p(95)=2.97ms 
       { expected_response:true }...: avg=938.23µs min=91.19µs  med=541.67µs max=51.92ms p(90)=2.03ms  p(95)=2.97ms 
     http_req_failed................: 0.00%   ✓ 0      ✗ 406248
     http_req_receiving.............: avg=36.96µs  min=7.52µs   med=13.84µs  max=29.03ms p(90)=39.23µs p(95)=90.32µs
     http_req_sending...............: avg=12.7µs   min=3.4µs    med=5.49µs   max=32.56ms p(90)=10.92µs p(95)=15.14µs
     http_req_tls_handshaking.......: avg=0s       min=0s       med=0s       max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=888.55µs min=72.99µs  med=508.26µs max=51.85ms p(90)=1.95ms  p(95)=2.86ms 
     http_reqs......................: 406248  16249.182908/s
     iteration_duration.............: avg=1.04ms   min=135.41µs med=623.29µs max=52.54ms p(90)=2.2ms   p(95)=3.22ms 
     iterations.....................: 406248  16249.182908/s
     vus............................: 3       min=3    max=29  
     vus_max........................: 30      min=30   max=30  
```

### nginx private proxy

Run load test: `k6 run loadtest-nginx-private-proxy.js`

```shell
running (25.0s), 00/30 VUs, 63390 complete and 0 interrupted iterations
default ✓ [======================================] 00/30 VUs  25s

     ✓ status is 200

     check_failure_rate.............: 0.00%   ✓ 0     ✗ 63390
     checks.........................: 100.00% ✓ 63390 ✗ 0    
     data_received..................: 54 MB   2.2 MB/s
     data_sent......................: 7.7 MB  307 kB/s
     http_req_blocked...............: avg=2.38µs  min=884ns    med=1.6µs   max=3.39ms  p(90)=2.46µs  p(95)=2.92µs 
     http_req_connecting............: avg=298ns   min=0s       med=0s      max=2.14ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=6.65ms  min=496.78µs med=4.94ms  max=81.77ms p(90)=12.57ms p(95)=18.07ms
       { expected_response:true }...: avg=6.65ms  min=496.78µs med=4.94ms  max=81.77ms p(90)=12.57ms p(95)=18.07ms
     http_req_failed................: 0.00%   ✓ 0     ✗ 63390
     http_req_receiving.............: avg=32.26µs min=12.51µs  med=24.74µs max=11.04ms p(90)=42.06µs p(95)=55.61µs
     http_req_sending...............: avg=13.69µs min=4.99µs   med=8.19µs  max=31.41ms p(90)=13.69µs p(95)=16.06µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=6.6ms   min=474.84µs med=4.91ms  max=81.66ms p(90)=12.51ms p(95)=18ms   
     http_reqs......................: 63390   2535.166203/s
     iteration_duration.............: avg=6.74ms  min=548.51µs med=5.03ms  max=81.83ms p(90)=12.68ms p(95)=18.23ms
     iterations.....................: 63390   2535.166203/s
     vus............................: 1       min=1   max=29 
     vus_max........................: 30      min=30  max=30 
```

### nginx private rego

Run load test: `k6 run loadtest-nginx-private-rego.js`

```shell
running (25.0s), 00/30 VUs, 50707 complete and 0 interrupted iterations
default ✓ [======================================] 00/30 VUs  25s

     ✓ status is 200

     check_failure_rate.............: 0.00%   ✓ 0     ✗ 50707
     checks.........................: 100.00% ✓ 50707 ✗ 0    
     data_received..................: 43 MB   1.7 MB/s
     data_sent......................: 6.1 MB  243 kB/s
     http_req_blocked...............: avg=2.28µs  min=810ns    med=1.52µs  max=1.68ms  p(90)=2.61µs  p(95)=2.97µs 
     http_req_connecting............: avg=252ns   min=0s       med=0s      max=1.47ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=8.34ms  min=360.45µs med=3.56ms  max=67.39ms p(90)=22.38ms p(95)=28.7ms 
       { expected_response:true }...: avg=8.34ms  min=360.45µs med=3.56ms  max=67.39ms p(90)=22.38ms p(95)=28.7ms 
     http_req_failed................: 0.00%   ✓ 0     ✗ 50707
     http_req_receiving.............: avg=33.82µs min=9.31µs   med=23.83µs max=21.97ms p(90)=52.27µs p(95)=59.67µs
     http_req_sending...............: avg=12.55µs min=4.36µs   med=7.68µs  max=15.2ms  p(90)=13.58µs p(95)=15.63µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=8.29ms  min=328.54µs med=3.51ms  max=67.36ms p(90)=22.32ms p(95)=28.64ms
     http_reqs......................: 50707   2028.021526/s
     iteration_duration.............: avg=8.43ms  min=420.46µs med=3.65ms  max=67.48ms p(90)=22.49ms p(95)=28.82ms
     iterations.....................: 50707   2028.021526/s
     vus............................: 1       min=1   max=29 
     vus_max........................: 30      min=30  max=30 
```

### nginx private opa

**PLEASE OBSERVE: OPA WILL ALWAYS RESPOND WITH OK HOWEVER THE RESULT IS EVALUATED**

Run load test: `k6 run loadtest-nginx-private-opa.js`

```shell
running (25.0s), 00/30 VUs, 51367 complete and 0 interrupted iterations
default ✓ [======================================] 00/30 VUs  25s

     ✓ status is 200

     check_failure_rate.............: 0.00%   ✓ 0     ✗ 51367
     checks.........................: 100.00% ✓ 51367 ✗ 0    
     data_received..................: 44 MB   1.7 MB/s
     data_sent......................: 6.1 MB  244 kB/s
     http_req_blocked...............: avg=2.6µs   min=708ns    med=1.5µs   max=3.58ms   p(90)=2.66µs  p(95)=2.99µs 
     http_req_connecting............: avg=560ns   min=0s       med=0s      max=3.42ms   p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=8.22ms  min=319.62µs med=3.36ms  max=128.61ms p(90)=20.47ms p(95)=36.5ms 
       { expected_response:true }...: avg=8.22ms  min=319.62µs med=3.36ms  max=128.61ms p(90)=20.47ms p(95)=36.5ms 
     http_req_failed................: 0.00%   ✓ 0     ✗ 51367
     http_req_receiving.............: avg=32.74µs min=11.03µs  med=23.11µs max=12.56ms  p(90)=52.05µs p(95)=59.48µs
     http_req_sending...............: avg=11.88µs min=4.61µs   med=7.7µs   max=9.5ms    p(90)=13.72µs p(95)=15.89µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=8.18ms  min=288.67µs med=3.32ms  max=128.57ms p(90)=20.41ms p(95)=36.44ms
     http_reqs......................: 51367   2054.510551/s
     iteration_duration.............: avg=8.32ms  min=382.99µs med=3.46ms  max=128.67ms p(90)=20.59ms p(95)=36.58ms
     iterations.....................: 51367   2054.510551/s
     vus............................: 1       min=1   max=29 
     vus_max........................: 30      min=30  max=30 
```
