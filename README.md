# opa-nginx-authz
Lab to test opa authz for nginx

## k6 load test

Start all services: `./run.sh`

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
running (25.0s), 00/30 VUs, 218135 complete and 0 interrupted iterations
default ✓ [======================================] 00/30 VUs  25s

     ✓ status is 200

     check_failure_rate.............: 0.00%   ✓ 0      ✗ 218135
     checks.........................: 100.00% ✓ 218135 ✗ 0     
     data_received..................: 16 MB   654 kB/s
     data_sent......................: 24 MB   977 kB/s
     http_req_blocked...............: avg=2.57µs  min=617ns    med=1.26µs  max=24.04ms p(90)=1.81µs  p(95)=2.43µs 
     http_req_connecting............: avg=58ns    min=0s       med=0s      max=3.05ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=1.85ms  min=167.68µs med=1.11ms  max=35.96ms p(90)=4.21ms  p(95)=6.01ms 
       { expected_response:true }...: avg=1.85ms  min=167.68µs med=1.11ms  max=35.96ms p(90)=4.21ms  p(95)=6.01ms 
     http_req_failed................: 0.00%   ✓ 0      ✗ 218135
     http_req_receiving.............: avg=21.19µs min=5.67µs   med=10.83µs max=20.58ms p(90)=18.01µs p(95)=23.66µs
     http_req_sending...............: avg=12.72µs min=3.33µs   med=6.33µs  max=20.54ms p(90)=9.91µs  p(95)=14.3µs 
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=1.82ms  min=148.76µs med=1.08ms  max=35.68ms p(90)=4.16ms  p(95)=5.94ms 
     http_reqs......................: 218135  8725.0431/s
     iteration_duration.............: avg=1.95ms  min=223.15µs med=1.19ms  max=36.05ms p(90)=4.36ms  p(95)=6.2ms  
     iterations.....................: 218135  8725.0431/s
     vus............................: 1       min=1    max=30  
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
running (25.0s), 00/30 VUs, 42288 complete and 0 interrupted iterations
default   [======================================] 00/30 VUs  25s

     ✓ status is 200

     check_failure_rate.............: 0.00%   ✓ 0     ✗ 42288
     checks.........................: 100.00% ✓ 42288 ✗ 0    
     data_received..................: 36 MB   1.4 MB/s
     data_sent......................: 5.1 MB  203 kB/s
     http_req_blocked...............: avg=2.62µs  min=933ns    med=1.69µs  max=6.06ms   p(90)=2.94µs  p(95)=3.42µs 
     http_req_connecting............: avg=363ns   min=0s       med=0s      max=6ms      p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=10.01ms min=428.89µs med=4.46ms  max=121.65ms p(90)=27.43ms p(95)=36.73ms
       { expected_response:true }...: avg=10.01ms min=428.89µs med=4.46ms  max=121.65ms p(90)=27.43ms p(95)=36.73ms
     http_req_failed................: 0.00%   ✓ 0     ✗ 42288
     http_req_receiving.............: avg=34.16µs min=10.92µs  med=25.95µs max=8.94ms   p(90)=50.72µs p(95)=59.79µs
     http_req_sending...............: avg=13.47µs min=4.24µs   med=8.62µs  max=15.84ms  p(90)=14.87µs p(95)=17.04µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=9.96ms  min=402.69µs med=4.41ms  max=121.55ms p(90)=27.35ms p(95)=36.68ms
     http_reqs......................: 42288   1691.377648/s
     iteration_duration.............: avg=10.1ms  min=485.79µs med=4.56ms  max=121.81ms p(90)=27.53ms p(95)=36.83ms
     iterations.....................: 42288   1691.377648/s
     vus............................: 1       min=1   max=29 
     vus_max........................: 30      min=30  max=30 
```