package wtutil

import (
	"fmt"
	"sync"
	"time"
)

var StdMap = make(map[int]int, 0)
var SyncMap sync.Map
var ManSyncMap ManualMap
var ReadMapChannel = make(chan int)
var WriteMapChannel = make(chan int)
var SetMapChannel = make(chan MapData)
var StopChannel = make(chan int)

type MapData struct {
	MapKey int
	MapVal int
}

type ManualMap struct {
	sync.Mutex
	MapStore map[int]int
}

// func TestWtTestSyncMapFibonacci() {
// 	WtTestSyncMapFibonacci()
// }

func WtTestSyncMapFibonacci() bool {
	var wg sync.WaitGroup

	SyncMap.Store(0, 0)
	SyncMap.Store(1, 1)

	ManSyncMap = ManualMap{
		MapStore: StdMap,
	}
	ManSyncMap.Save(0, 0)
	ManSyncMap.Save(1, 1)

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
		go NoMapFibo(idx, x, &wg)
		go ManualSyncMutex(idy, y, &wg)
		go SyncMapProcess(idz, z, &wg)
	}

	// for i := 0; i < len(input3); i++ {

	// }

	// for i := 0; i < len(input2); i++ {

	// }

	wg.Wait()
	return true
}

func (c *ManualMap) Get(x interface{}) (int, bool) {
	defer c.Unlock()

	var key int
	var isInt bool
	if key, isInt = x.(int); !isInt {
		return 0, false
	}

	c.Lock()
	var val int
	var found bool
	if val, found = c.MapStore[key]; !found {
		return 0, false
	}
	return val, true
}

func (c *ManualMap) Save(k int, v interface{}) error {
	defer c.Unlock()

	var val int
	var isInt bool
	if val, isInt = v.(int); !isInt {
		return fmt.Errorf("VALUE IS NOT AN INTEGER")
	}
	c.Lock()
	c.MapStore[k] = val

	return nil
}

// var write = make(chan *DataMdl,5)
// var read = make(chan *DataMdl,5)

type DataMdl struct {
	Key   int
	Value int
}

//fastest
func SyncMapProcess(id int, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now().UnixNano()
	result := SyncFib(n)
	end := time.Now().UnixNano()

	elapsed := end - start
	fmt.Printf("%d,%d,%d,%d,%d,%d\n", start, end, elapsed, n, result, id)
	// fmt.Printf("Start %d - End %d| Required %d ns| N= %d Fib is: %d , Using  synch.Map  Processed by goroutine:  %d \n", start, end, elapsed, n, result, id)
}

//fast
func ManualSyncMutex(id int, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now().UnixNano()
	result := ManualSyncFib(n)
	end := time.Now().UnixNano()

	elapsed := end - start
	fmt.Printf("%d,%d,%d,%d,%d,%d\n", start, end, elapsed, n, result, id)

	// fmt.Printf("Start %d - End %d| Required %d ns| N= %d Fib is: %d , Manual sync.Mutex Processed by goroutine:  %d \n", start, end, elapsed, n, result, id)
}

//slow
func NoMapFibo(id int, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now().UnixNano()
	result := Fib(n)
	end := time.Now().UnixNano()

	elapsed := end - start
	fmt.Printf("%d,%d,%d,%d,%d,%d\n", start, end, elapsed, n, result, id)
	// fmt.Printf("Start %d - End %d| Required %d ns| N= %d Fib is: %d , Without Using Map Processed by goroutine:  %d \n", start, end, elapsed, n, result, id)
}

//using sync.Map
func SyncFib(n int) int {
	if n < 1 {
		return 0
	}
	if n < 3 {
		return 1
	}
	var f int
	if v, ok := SyncMap.Load(n); ok {
		var isInt bool
		if f, isInt = v.(int); !isInt {
			return 0
		}

	} else {
		f = SyncFib(n-2) + SyncFib(n-1)
		SyncMap.Store(n, f)
	}

	return f
}

//using struct with sync.Mutex
func ManualSyncFib(n int) int {
	if n < 1 {
		return 0
	}
	if n < 3 {
		return 1
	}
	var f int
	if v, found := ManSyncMap.Get(n); found {
		f = v
	} else {
		f = SyncFib(n-2) + SyncFib(n-1)
		ManSyncMap.Save(n, f) //don't care for error coz we're passing int
	}

	return f
}

//without map
func Fib(n int) int {
	if n < 1 {
		return 0
	}
	if n < 3 {
		return 1
	}

	var f int

	for i := 0; i < n; i++ {
		f = Fib(i) + Fib(i-1)
	}

	return f
}

//using aditional go routine for handle map

func ChannelSyncFibo(id int, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now().UnixNano()
	result := ChannelSyncFib(n)
	end := time.Now().UnixNano()

	elapsed := end - start
	fmt.Printf("%d,%d,%d,%d,%d,%d\n", start, end, elapsed, n, result, id)

	// fmt.Printf("Start %d - End %d| Required %d ns| N= %d Fib is: %d , Without Using Map Processed by goroutine:  %d \n", start, end, elapsed, n, result, id)
}

func ChannelSyncFib(n int) int {
	if n < 1 {
		return 0
	}
	if n < 3 {
		return 1
	}
	var f int
	ReadMapChannel <- n
	process := true
	for process {
		mapVal := <-WriteMapChannel
		if mapVal >= 0 {
			f = mapVal
			process = false
		}
	}
	if f > 0 {
		return f
	}

	if f == 0 {
		// checkFibo := true
		f1, f2 := 0, 0
		f1 = ChannelSyncFib(n - 1)
		f2 = ChannelSyncFib(n - 2)
		f = f1 + f2
		mapData := MapData{
			MapKey: n,
			MapVal: f,
		}
		SetMapChannel <- mapData

	}

	return f
}

func MapRoutine() {
	fmt.Println("Start MapRoutine")
	knownFibo := map[int]int{
		0: 0,
		1: 1,
	}
	var mapData MapData
	running := true
	for running {
		select {
		case mapKey := <-ReadMapChannel:
			{

				found := false
				value := 0
				if value, found = knownFibo[mapKey]; !found {
					// fmt.Println("received", mapKey)
					WriteMapChannel <- 0
				} else {
					WriteMapChannel <- value
				}
			}
		case mapData = <-SetMapChannel:
			{
				knownFibo[mapData.MapKey] = mapData.MapVal
				// fmt.Printf("set key:%d val: %d \n", mapData.MapKey, mapData.MapVal)
			}

		case stopKey := <-StopChannel:
			running = false
			fmt.Println("Stop :", stopKey)
		}
	}
	fmt.Println("Stop MapRoutine")
}

func FibMap(n int, memo *map[int]int) int {
	if n < 1 {
		return 0
	}
	if n < 3 {
		return 1
	}

	var f int

	for i := 0; i < n; i++ {
		f = FibMap(i, memo) + FibMap(i-1, memo)
	}

	return f
}
