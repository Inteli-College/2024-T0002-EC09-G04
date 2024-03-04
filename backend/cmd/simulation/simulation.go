package main

import (
	"sync"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/pkg/station"
)


func main() {
	numStations := 10
	var wg sync.WaitGroup

	for i := 0; i < numStations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			station.Start("ssl://908447a66faf43129ef280ff434012e6.s1.eu.hivemq.cloud:8883/mqtt:1883")
		}()
	}
	wg.Wait()
}