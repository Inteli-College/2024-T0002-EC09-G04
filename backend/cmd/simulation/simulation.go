package main

import (
	"sync"
	"github.com/Inteli-College/2024-T0002-EC09-G04/pkg/station"
)


func main() {
	numStations := 10
	var wg sync.WaitGroup

	for i := 0; i < numStations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			station.Start("tcp://broker:1891")
		}()
	}
	wg.Wait()
}