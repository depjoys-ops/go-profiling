package main

import (
	"context"
	"log"
	"os"
	"profiling/factorial"
	"runtime/pprof"
)

func main() {
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer cpuFile.Close()

	err = pprof.StartCPUProfile(cpuFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func(){
		pprof.StopCPUProfile()
	}()


	ctx := context.Background()
	pprof.Do(ctx, pprof.Labels("label", "calculate"), func(ctx context.Context) {
		factorial.Calculate()
	})


	memFile, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer memFile.Close()

	err = pprof.WriteHeapProfile(memFile)
	if err != nil {
		log.Fatal(err)
	}

}