package main

import (
	"sync"
	"testing"
)

var routinesFew = 1000
var routinesModerate = 10000
var routinesMany = 100000
var syncMap = sync.Map{}
var simpleMap map[int64]int64 = map[int64]int64{1: 1}
var lock = sync.Mutex{}

var testInt64 int64

func writeSyncMap() {
	testInt64++
	syncMap.Store(testInt64, testInt64)
}

func incMutexMap() {
	lock.Lock()
	testInt64++
	simpleMap[testInt64] = testInt64
	lock.Unlock()
}

func readSyncMap() int64 {
	res, ok := syncMap.Load(testInt64)
	if !ok {
		return 0
	}
	return res.(int64)
}

func readMutexMap() int64 {
	lock.Lock()
	res := simpleMap[testInt64]
	lock.Unlock()
	return res
}

func BenchmarkSyncMapMostlyReadRoutineFew(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesFew)

	for ri := 0; ri < routinesFew/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				writeSyncMap()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesFew/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readSyncMap()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkMutexMapMostlyReadFew(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesFew)

	for ri := 0; ri < routinesFew/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incMutexMap()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesFew/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readMutexMap()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkSyncMapMostlyWriteFew(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesFew)

	for ri := 0; ri < routinesFew/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				writeSyncMap()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesFew/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readSyncMap()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkMutexMapMostlyWriteFew(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesFew)

	for ri := 0; ri < routinesFew/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incMutexMap()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesFew/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readMutexMap()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkSyncMapMostlyReadRoutinesModerate(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesModerate)

	for ri := 0; ri < routinesModerate/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				writeSyncMap()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesModerate/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readSyncMap()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkMutexMapMostlyReadRoutinesModerate(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesModerate)

	for ri := 0; ri < routinesModerate/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incMutexMap()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesModerate/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readMutexMap()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkSyncMapMostlyWriteRoutinesModerate(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesModerate)

	for ri := 0; ri < routinesModerate/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				writeSyncMap()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesModerate/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readSyncMap()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkMutexMapMostlyWriteRoutinesModerate(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesModerate)

	for ri := 0; ri < routinesModerate/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incMutexMap()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesModerate/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readMutexMap()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkSyncMutexMostlyReadRoutinesMany(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesMany)

	for ri := 0; ri < routinesMany/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				writeSyncMap()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesMany/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readSyncMap()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkMutexMapMostlyReadRoutinesMany(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesMany)

	for ri := 0; ri < routinesMany/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incMutexMap()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesMany/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readMutexMap()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkSyncMutexMostlyWriteRoutinesMany(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesMany)

	for ri := 0; ri < routinesMany/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				writeSyncMap()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesMany/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readSyncMap()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkMutexMapMostlyWriteRoutinesMany(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesMany)

	for ri := 0; ri < routinesMany/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incMutexMap()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesMany/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readMutexMap()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
