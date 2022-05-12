# Go profiling 101

1. create a test to run whatever you want to profile
2. run it like this:
   ```sh
   # go test -cpuprofile [file] -memprofile [file] -bench -v [your package]
   go test -cpuprofile cpu.prof -memprofile mem.prof -bench -v ./pkg/fibonacci
   ```
3. open the profile:
   ```sh
   go tool pprof cpu.prof
   ```
4. play around; e.g.:
   ```sh
   > go tool pprof cpu.prof                                                     
     Type: cpu
     Time: May 12, 2022 at 12:47pm (CEST)
     Duration: 7.82s, Total samples = 9.78s (125.12%)
     Entering interactive mode (type "help" for commands, "o" for options)
   > (pprof) top
     Showing nodes accounting for 7480ms, 76.48% of 9780ms total
     Dropped 90 nodes (cum <= 48.90ms)
     Showing top 10 nodes out of 103
     flat  flat%   sum%        cum   cum%
     3160ms 32.31% 32.31%     3220ms 32.92%  runtime.kevent
     1450ms 14.83% 47.14%     1580ms 16.16%  runtime.pthread_cond_wait
     610ms  6.24% 53.37%      610ms  6.24%  runtime.procyield
     520ms  5.32% 58.69%     2340ms 23.93%  github.com/wernerdweight/tmp/pkg/fibonacci.fibonacciBelow
     490ms  5.01% 63.70%      490ms  5.01%  runtime.pthread_kill
     400ms  4.09% 67.79%      400ms  4.09%  runtime.memclrNoHeapPointers
     260ms  2.66% 70.45%      260ms  2.66%  runtime.libcCall
     220ms  2.25% 72.70%      220ms  2.25%  runtime.pthread_cond_signal
     200ms  2.04% 74.74%     1450ms 14.83%  runtime.mallocgc
     170ms  1.74% 76.48%     1810ms 18.51%  runtime.growslice
   ```
