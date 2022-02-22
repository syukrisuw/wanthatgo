package wtutil

// import (
// 	"sync"
// 	"testing"
// )

// func TestWtTestSyncMapFibonacci(t *testing.T) {

// 	result := WtTestSyncMapFibonacci()
// 	if !result {
// 		t.Fatalf(`Testing Failed`)
// 	}

// }

// func TestNoMapFibo(t *testing.T) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			t.Fatalf("Test Failed with err %v", err)
// 		}
// 	}()

// 	var wg sync.WaitGroup
// 	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 23, 25}
// 	wg.Add(len(input))
// 	idx := 100
// 	for i := 0; i < len(input); i++ {
// 		idx++
// 		var x = input[i]
// 		go NoMapFibo(idx, x, &wg)
// 	}
// 	wg.Wait()
// }

// func TestSyncMapProcess(t *testing.T) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			t.Fatalf("Test Failed with err %v", err)
// 		}
// 	}()
// 	SyncMap.Store(0, 0)
// 	SyncMap.Store(1, 1)
// 	SyncMap.Store(2, 1)
// 	var wg sync.WaitGroup
// 	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 23, 25}
// 	wg.Add(len(input))
// 	idx := 100
// 	for i := 0; i < len(input); i++ {
// 		idx++
// 		var x = input[i]
// 		go SyncMapProcess(idx, x, &wg)
// 	}
// 	wg.Wait()
// }

// func TestManualSyncMutex(t *testing.T) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			t.Fatalf("Test Failed with err %v", err)
// 		}
// 	}()
// 	ManSyncMap = ManualMap{
// 		MapStore: StdMap,
// 	}
// 	ManSyncMap.Save(0, 0)
// 	ManSyncMap.Save(1, 1)
// 	ManSyncMap.Save(2, 1)
// 	var wg sync.WaitGroup
// 	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 23, 25}
// 	wg.Add(len(input))
// 	idx := 100
// 	for i := 0; i < len(input); i++ {
// 		idx++
// 		var x = input[i]
// 		go ManualSyncMutex(idx, x, &wg)
// 	}
// 	wg.Wait()
// }

// func TestChannelSyncFibo(t *testing.T) {
// 	defer func() {
// 		StopChannel <- 1
// 		if err := recover(); err != nil {
// 			t.Fatalf("Test Failed with err %v", err)
// 		}
// 	}()

// 	var wg sync.WaitGroup
// 	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 23, 25}
// 	wg.Add(len(input))
// 	idx := 500
// 	go MapRoutine()
// 	for i := 0; i < len(input); i++ {
// 		idx++
// 		var x = input[i]
// 		go ChannelSyncFibo(idx, x, &wg)
// 	}

// 	wg.Wait()
// }

// ManSyncMap = ManualMap{
// 	MapStore: StdMap,
// }
// ManSyncMap.Save(0, 0)
// ManSyncMap.Save(1, 1)
// ManSyncMap.Save(2, 1)

// input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 23, 25}
// input2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 90, 98, 100}
// input3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 90, 98, 100}
// wg.Add(len(input))
// wg.Add(len(input2))
// wg.Add(len(input3))
// fmt.Println("Start,End,Duration,FibNo,FibValue,GoRoutineId")
// idx, idy, idz := 100, 200, 300
// for i := 0; i < len(input); i++ {
// 	idx++
// 	idy++
// 	idz++
// 	var x = input[i]
// 	var y = input2[i]
// 	var z = input3[i]
// 	go NoMapFibo(idx, x, &wg)
// 	go ManualSyncMutex(idy, y, &wg)
// 	go SyncMapProcess(idz, z, &wg)
// }

// // for i := 0; i < len(input3); i++ {

// // }

// // for i := 0; i < len(input2); i++ {

// // }

// wg.Wait()
