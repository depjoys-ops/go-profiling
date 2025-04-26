### Tools for profiling in Go 
- Benchmarks <br/>
- Profiling to file (runtime/pprof) <br/>
- Web profiling (net/http/pprof) <br/>

## Benchmarks

> go test -v ./... -bench='Benchmark(Recursive|Dynamic)' -benchmem -benchtime=100x -count=2

![screenshot](./imageFolder/screenshot1.png)

```
go test -v ./... -bench="BenchmarkCalculate" -run=^# -count=10 | tee CalcRecursive.txt
go test -v ./... -bench="BenchmarkCalculate" -run=^# -count=10 | tee CalcDynamic.txt
go install golang.org/x/perf/cmd/benchstat@latest
benchstat CalcRecursive.txt CalcDynamic.txt
```
![screenshot](./imageFolder/screenshot2.png)

## Profiling

### Dump from benchmarks cpu.prof and mem.prof for analyze
> go test -v ./... -bench="BenchmarkCalculate" -cpuprofile='cpu.prof' -memprofile='mem.prof'

### To analyze from the console, use
> go tool pprof cpu.prof

![screenshot](./imageFolder/screenshot3.png)

> go tool pprof mem.prof

![screenshot](./imageFolder/screenshot4.png)
