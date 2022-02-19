package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wthelper.syncMap.Store(0, 0)
	wthelper.syncMap.Store(1, 1)
	wthelper.syncMap.Store(2, 1)

	manSyncMap := wthelper.ManualMap{
		mapStore: stdMap,
	}
	manSyncMap.Save(0, 0)
	manSyncMap.Save(1, 1)
	manSyncMap.Save(2, 1)

	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 23, 25}
	input2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 90, 98, 100}
	input3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 90, 98, 100}
	wg.Add(len(input))
	wg.Add(len(input2))
	wg.Add(len(input3))
	fmt.Println("Start,End,Duration,FibNo,FibValue,GoRoutineId")
	idx, idy, idz := 100, 200, 300
	for i := 0; i < len(input); i++ {
		idx++
		idy++
		idz++
		var x = input[i]
		var y = input2[i]
		var z = input3[i]
		go process(idx, x, &wg)
		go manualSyncMutex(idy, y, &wg)
		go syncMapProcess(idz, z, &wg)
	}

	// for i := 0; i < len(input3); i++ {

	// }

	// for i := 0; i < len(input2); i++ {

	// }

	wg.Wait()
	// time.Sleep(5000)
}

//fastest
