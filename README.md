# opa-nginx-authz
Lab to test opa authz for nginx

## k6 load test

Start all services: `./run.sh`

### OPA

Run load test: `k6 run loadtest-opa.js`

```shell
running (25.0s), 00/30 VUs, 130158 complete and 0 interrupted iterations
default ✓ [======================================] 00/30 VUs  25s

     data_received..................: 23 MB  922 kB/s
     data_sent......................: 22 MB  859 kB/s
     http_req_blocked...............: avg=1.77µs  min=689ns    med=1.41µs  max=1.67ms  p(90)=1.7µs   p(95)=2.2µs  
     http_req_connecting............: avg=44ns    min=0s       med=0s      max=1.61ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=3.21ms  min=247.54µs med=2ms     max=68.57ms p(90)=7.48ms  p(95)=10ms   
       { expected_response:true }...: avg=3.21ms  min=247.54µs med=2ms     max=68.57ms p(90)=7.48ms  p(95)=10ms   
     http_req_failed................: 0.00%  ✓ 0    ✗ 130158
     http_req_receiving.............: avg=24.58µs min=8.26µs   med=18.67µs max=19.53ms p(90)=27.23µs p(95)=31.74µs
     http_req_sending...............: avg=12.13µs min=4.53µs   med=8.44µs  max=15.62ms p(90)=11.63µs p(95)=14.79µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=3.18ms  min=219.23µs med=1.97ms  max=68.55ms p(90)=7.44ms  p(95)=9.94ms 
     http_reqs......................: 130158 5206.047259/s
     iteration_duration.............: avg=3.27ms  min=287.63µs med=2.06ms  max=68.64ms p(90)=7.54ms  p(95)=10.06ms
     iterations.....................: 130158 5206.047259/s
     vus............................: 1      min=1  max=29  
     vus_max........................: 30     min=30 max=30 
```

### opa-nginx-external-auth

Run load test: `k6 run loadtest-go.js`

```shell
running (25.0s), 00/30 VUs, 77709 complete and 0 interrupted iterations
default   [======================================] 00/30 VUs  25s

     data_received..................: 5.8 MB 233 kB/s
     data_sent......................: 8.4 MB 336 kB/s
     http_req_blocked...............: avg=1.88µs  min=701ns    med=1.51µs max=3.6ms   p(90)=1.86µs  p(95)=2.41µs 
     http_req_connecting............: avg=100ns   min=0s       med=0s     max=3.54ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=5.44ms  min=429.64µs med=3.82ms max=56.84ms p(90)=11.95ms p(95)=15.34ms
       { expected_response:true }...: avg=5.44ms  min=429.64µs med=3.82ms max=56.84ms p(90)=11.95ms p(95)=15.34ms
     http_req_failed................: 0.00%  ✓ 0    ✗ 77709
     http_req_receiving.............: avg=19.81µs min=6.72µs   med=17.3µs max=7.51ms  p(90)=23.39µs p(95)=27.31µs
     http_req_sending...............: avg=10.78µs min=4.48µs   med=7.88µs max=15.91ms p(90)=11.36µs p(95)=13.55µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s     max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=5.41ms  min=404.28µs med=3.79ms max=56.82ms p(90)=11.91ms p(95)=15.3ms 
     http_reqs......................: 77709  3108.189419/s
     iteration_duration.............: avg=5.49ms  min=477.9µs  med=3.87ms max=56.88ms p(90)=12ms    p(95)=15.39ms
     iterations.....................: 77709  3108.189419/s
     vus............................: 3      min=3  max=29 
     vus_max........................: 30     min=30 max=30 
```

### nginx public

Run load test: `k6 run loadtest-nginx-public.js`

```shell
running (25.0s), 00/30 VUs, 420676 complete and 0 interrupted iterations
default ✓ [======================================] 00/30 VUs  25s

     ✓ status is 200

     check_failure_rate.............: 0.00%   ✓ 0      ✗ 420676
     checks.........................: 100.00% ✓ 420676 ✗ 0     
     data_received..................: 358 MB  14 MB/s
     data_sent......................: 36 MB   1.4 MB/s
     http_req_blocked...............: avg=2.85µs   min=608ns    med=1.15µs   max=20.87ms p(90)=1.83µs p(95)=2.4µs  
     http_req_connecting............: avg=212ns    min=0s       med=0s       max=4.67ms  p(90)=0s     p(95)=0s     
     http_req_duration..............: avg=909.33µs min=84.43µs  med=528.31µs max=35.78ms p(90)=1.99ms p(95)=2.9ms  
       { expected_response:true }...: avg=909.33µs min=84.43µs  med=528.31µs max=35.78ms p(90)=1.99ms p(95)=2.9ms  
     http_req_failed................: 0.00%   ✓ 0      ✗ 420676
     http_req_receiving.............: avg=34.62µs  min=7.74µs   med=13.66µs  max=25.25ms p(90)=33.5µs p(95)=84.12µs
     http_req_sending...............: avg=12.26µs  min=3.47µs   med=5.39µs   max=18.91ms p(90)=9.99µs p(95)=14.15µs
     http_req_tls_handshaking.......: avg=0s       min=0s       med=0s       max=0s      p(90)=0s     p(95)=0s     
     http_req_waiting...............: avg=862.43µs min=66.43µs  med=498.74µs max=34.28ms p(90)=1.91ms p(95)=2.79ms 
     http_reqs......................: 420676  16826.312401/s
     iteration_duration.............: avg=1ms      min=130.57µs med=602.43µs max=36.01ms p(90)=2.14ms p(95)=3.12ms 
     iterations.....................: 420676  16826.312401/s
     vus............................: 0       min=0    max=29  
     vus_max........................: 30      min=30   max=30  
```

### nginx private proxy

Run load test: `k6 run loadtest-nginx-private-proxy.js`

```shell
running (25.0s), 00/30 VUs, 44656 complete and 0 interrupted iterations
default ✓ [======================================] 00/30 VUs  25s

     ✓ status is 200

     check_failure_rate.............: 0.00%   ✓ 0     ✗ 44656
     checks.........................: 100.00% ✓ 44656 ✗ 0    
     data_received..................: 38 MB   1.5 MB/s
     data_sent......................: 5.4 MB  216 kB/s
     http_req_blocked...............: avg=2.65µs  min=1.03µs   med=1.73µs  max=3.06ms  p(90)=2.71µs  p(95)=3.33µs 
     http_req_connecting............: avg=267ns   min=0s       med=0s      max=1.05ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=9.47ms  min=831.38µs med=7.16ms  max=80.39ms p(90)=19.12ms p(95)=24.11ms
       { expected_response:true }...: avg=9.47ms  min=831.38µs med=7.16ms  max=80.39ms p(90)=19.12ms p(95)=24.11ms
     http_req_failed................: 0.00%   ✓ 0     ✗ 44656
     http_req_receiving.............: avg=39.15µs min=13.47µs  med=29.06µs max=21.78ms p(90)=47.54µs p(95)=58.68µs
     http_req_sending...............: avg=14.05µs min=5.3µs    med=9.24µs  max=14.32ms p(90)=14.87µs p(95)=17.83µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=9.41ms  min=798.19µs med=7.12ms  max=80.36ms p(90)=19.07ms p(95)=24.04ms
     http_reqs......................: 44656   1785.999113/s
     iteration_duration.............: avg=9.57ms  min=897.57µs med=7.27ms  max=80.45ms p(90)=19.24ms p(95)=24.23ms
     iterations.....................: 44656   1785.999113/s
     vus............................: 1       min=1   max=29 
     vus_max........................: 30      min=30  max=30 
```