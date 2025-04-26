### Tools for profiling in Go 
- Benchmarks <br/>
- Profiling to file (runtime/pprof) <br/>
- Web profiling (net/http/pprof) <br/>

> go test -v ./... -bench=. -benchmem -benchtime=100x -count=2