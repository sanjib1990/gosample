#### Run benchmark test

Like unit tests in Go, benchmark functions are placed in a _test.go file, and each benchmark function is expected to have func BenchmarkXxx(*testing.B) as a signature, with the testing.B type managing the benchmark’s timing.

b.N specifies the number of iterations; the value is not fixed, but dynamically allocated, ensuring that the benchmark runs for at least one second by default.

To run all benchmarks, use -bench=.

```bash
 go test -bench=.
```

To verify that the benchmark produces a consistent result, you can run it multiple times by passing a number to the -count flag:

```bash
 go test -bench=. -count=5
```

#### Skipping unit tests
To avoid executing any test functions in the test files, pass a regular expression to the -run flag:

```bash
 go test -bench=. -count 5 -run=^#
```

#### Adjusting the minimum time
we can increase the minimum amount of time that the benchmark should run using the -benchtime flag

```bash
 go test -bench=. -run=^# -benchtime=10s
```

#### Display memory allocation statistics
To include memory allocation statistics in the benchmark output, add the -benchmem flag while running the benchmarks

```bash
 go test -bench=. -run=^# -benchtime=10s -benchmem
```

#### Comparing benchmark results
To compare the output of both implementations of our benchmark with benchstat

```bash
 go test -bench=<oldimpl> -count 5 | tee old.txt
 go install golang.org/x/perf/cmd/benchstat@latest
 go test -bench=<new impl> -count 5 | tee new.txt
 benchstat old.txt new.txt
```



