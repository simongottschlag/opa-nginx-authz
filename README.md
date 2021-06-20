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
running (25.0s), 00/30 VUs, 119911 complete and 0 interrupted iterations
default ✓ [======================================] 00/30 VUs  25s

     ✓ status is 200

     check_failure_rate.............: 0.00%   ✓ 0      ✗ 119911
     checks.........................: 100.00% ✓ 119911 ✗ 0     
     data_received..................: 102 MB  4.1 MB/s
     data_sent......................: 14 MB   580 kB/s
     http_req_blocked...............: avg=1.99µs  min=766ns    med=1.47µs  max=2.84ms  p(90)=1.88µs  p(95)=2.39µs 
     http_req_connecting............: avg=155ns   min=0s       med=0s      max=1.06ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=3.47ms  min=357.79µs med=2.65ms  max=41.64ms p(90)=7.08ms  p(95)=9.23ms 
       { expected_response:true }...: avg=3.47ms  min=357.79µs med=2.65ms  max=41.64ms p(90)=7.08ms  p(95)=9.23ms 
     http_req_failed................: 0.00%   ✓ 0      ✗ 119911
     http_req_receiving.............: avg=28.19µs min=9.63µs   med=22.23µs max=15.88ms p(90)=31.31µs p(95)=36.38µs
     http_req_sending...............: avg=11.35µs min=4.51µs   med=7.53µs  max=15.18ms p(90)=11.13µs p(95)=13.9µs 
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=3.44ms  min=324.46µs med=2.62ms  max=41.6ms  p(90)=7.03ms  p(95)=9.16ms 
     http_reqs......................: 119911  4796.152669/s
     iteration_duration.............: avg=3.55ms  min=414.45µs med=2.73ms  max=41.71ms p(90)=7.18ms  p(95)=9.34ms 
     iterations.....................: 119911  4796.152669/s
     vus............................: 1       min=1    max=29  
     vus_max........................: 30      min=30   max=30 
```

### nginx private rego

Run load test: `k6 run loadtest-nginx-private-rego.js`

```shell
running (25.0s), 00/30 VUs, 199494 complete and 0 interrupted iterations
default ✓ [======================================] 00/30 VUs  25s

     ✓ status is 200

     check_failure_rate.............: 0.00%   ✓ 0      ✗ 199494
     checks.........................: 100.00% ✓ 199494 ✗ 0     
     data_received..................: 170 MB  6.8 MB/s
     data_sent......................: 24 MB   958 kB/s
     http_req_blocked...............: avg=2.55µs  min=659ns    med=1.33µs  max=10.42ms p(90)=1.84µs p(95)=2.39µs 
     http_req_connecting............: avg=312ns   min=0s       med=0s      max=3.59ms  p(90)=0s     p(95)=0s     
     http_req_duration..............: avg=2.04ms  min=202.62µs med=1.42ms  max=34.16ms p(90)=4.27ms p(95)=5.8ms  
       { expected_response:true }...: avg=2.04ms  min=202.62µs med=1.42ms  max=34.16ms p(90)=4.27ms p(95)=5.8ms  
     http_req_failed................: 0.00%   ✓ 0      ✗ 199494
     http_req_receiving.............: avg=30.62µs min=9.2µs    med=17.18µs max=20.69ms p(90)=27.2µs p(95)=36.21µs
     http_req_sending...............: avg=12.85µs min=3.84µs   med=6.65µs  max=20.59ms p(90)=9.74µs p(95)=13.98µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s     p(95)=0s     
     http_req_waiting...............: avg=1.99ms  min=184.75µs med=1.38ms  max=33.99ms p(90)=4.21ms p(95)=5.71ms 
     http_reqs......................: 199494  7979.292418/s
     iteration_duration.............: avg=2.13ms  min=251.11µs med=1.5ms   max=34.24ms p(90)=4.41ms p(95)=5.98ms 
     iterations.....................: 199494  7979.292418/s
     vus............................: 1       min=1    max=29  
     vus_max........................: 30      min=30   max=30  
```

### nginx private opa

**PLEASE OBSERVE: OPA WILL ALWAYS RESPOND WITH OK HOWEVER THE RESULT IS EVALUATED**

Run load test: `k6 run loadtest-nginx-private-opa.js`

```shell
running (25.0s), 00/30 VUs, 33899 complete and 0 interrupted iterations
default ✓ [======================================] 00/30 VUs  25s

     ✗ status is 200
      ↳  83% — ✓ 28231 / ✗ 5668

     check_failure_rate.............: 16.72% ✓ 5668  ✗ 28231
     checks.........................: 83.27% ✓ 28231 ✗ 5668 
     data_received..................: 28 MB  1.1 MB/s
     data_sent......................: 4.0 MB 161 kB/s
     http_req_blocked...............: avg=31.17µs min=892ns    med=1.58µs  max=15.33ms p(90)=112.21µs p(95)=136.45µs
     http_req_connecting............: avg=21.64µs min=0s       med=0s      max=15.02ms p(90)=71.48µs  p(95)=87.06µs 
     http_req_duration..............: avg=12.49ms min=313.34µs med=10.37ms max=89.58ms p(90)=28.41ms  p(95)=33.98ms 
       { expected_response:true }...: avg=11.2ms  min=313.34µs med=6.33ms  max=89.58ms p(90)=27.19ms  p(95)=33.23ms 
     http_req_failed................: 16.72% ✓ 5668  ✗ 28231
     http_req_receiving.............: avg=41.98µs min=10.64µs  med=26.02µs max=13.65ms p(90)=72.94µs  p(95)=88.66µs 
     http_req_sending...............: avg=18.31µs min=4.4µs    med=8.26µs  max=14.38ms p(90)=43.34µs  p(95)=51.85µs 
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s       p(95)=0s      
     http_req_waiting...............: avg=12.43ms min=286.56µs med=10.31ms max=89.54ms p(90)=28.31ms  p(95)=33.89ms 
     http_reqs......................: 33899  1355.734075/s
     iteration_duration.............: avg=12.61ms min=377.01µs med=10.5ms  max=89.66ms p(90)=28.56ms  p(95)=34.12ms 
     iterations.....................: 33899  1355.734075/s
     vus............................: 1      min=1   max=29 
     vus_max........................: 30     min=30  max=30 
```
