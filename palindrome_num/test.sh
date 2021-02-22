#!/bin/bash

echo "Test Bench"

go test -bench=. -cpuprofile=cpu.test -memprofile=mem.test

go tool pprof cpu.test
