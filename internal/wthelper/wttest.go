package helper

import (
	"fmt"
	"sync"
	"time"
)

var stdMap = make(map[int]int, 0)
var syncMap sync.Map
var manSyncMap ManualMap

type ManualMap struct {
	sync.Mutex
	mapStore map[int]int
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
	if val, found = c.mapStore[key]; !found {
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
	c.mapStore[k] = val

	return nil
}

// var write = make(chan *DataMdl,5)
// var read = make(chan *DataMdl,5)

type DataMdl struct {
	Key   int
	Value int
}

func syncMapProcess(id int, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now().UnixNano()
	result := SyncFib(n)
	end := time.Now().UnixNano()

	elapsed := end - start
	fmt.Printf("%d,%d,%d,%d,%d,%d\n", start, end, elapsed, n, result, id)
	// fmt.Printf("Start %d - End %d| Required %d ns| N= %d Fib is: %d , Using  synch.Map  Processed by goroutine:  %d \n", start, end, elapsed, n, result, id)
}

//fast
func manualSyncMutex(id int, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now().UnixNano()
	result := ManualSyncFib(n)
	end := time.Now().UnixNano()

	elapsed := end - start
	fmt.Printf("%d,%d,%d,%d,%d,%d\n", start, end, elapsed, n, result, id)

	// fmt.Printf("Start %d - End %d| Required %d ns| N= %d Fib is: %d , Manual sync.Mutex Processed by goroutine:  %d \n", start, end, elapsed, n, result, id)
}

//slow
func process(id int, n int, wg *sync.WaitGroup) {
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
	if v, ok := syncMap.Load(n); ok {
		var isInt bool
		if f, isInt = v.(int); !isInt {
			return 0
		}

	} else {
		f = SyncFib(n-2) + SyncFib(n-1)
		syncMap.Store(n, f)
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
	if v, found := manSyncMap.Get(n); found {
		f = v
	} else {
		f = SyncFib(n-2) + SyncFib(n-1)
		manSyncMap.Save(n, f) //don't care for error coz we're passing int
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
