package main

import (
	"sync"
	"testing"
)

var routinesFew = 1000
var routinesModerate = 10000
var routinesMany = 100000
var lockRW = sync.RWMutex{}
var lock = sync.Mutex{}

var testInt64 int64

func incRWMutex() {
	lockRW.Lock()
	testInt64++
	lockRW.Unlock()
}

func incMutex() {
	lock.Lock()
	testInt64++
	lock.Unlock()
}

func readRWMutex() int64 {
	lockRW.RLock()
	res := testInt64
	lockRW.RUnlock()
	return res
}

func readMutex() int64 {
	lock.Lock()
	res := testInt64
	lock.Unlock()
	return res
}

func BenchmarkRWMutexMostlyReadFew(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesFew)

	for ri := 0; ri < routinesFew/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incRWMutex()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesFew/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readRWMutex()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkMutexMostlyReadFew(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesFew)

	for ri := 0; ri < routinesFew/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incMutex()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesFew/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readMutex()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkRWMutexMostlyWriteFew(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesFew)

	for ri := 0; ri < routinesFew/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incRWMutex()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesFew/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readRWMutex()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkMutexMostlyWriteFew(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesFew)

	for ri := 0; ri < routinesFew/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incMutex()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesFew/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readMutex()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkRWMutexMostlyReadRoutinesModerate(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesModerate)

	for ri := 0; ri < routinesModerate/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incRWMutex()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesModerate/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readRWMutex()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkMutexMostlyReadRoutinesModerate(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesModerate)

	for ri := 0; ri < routinesModerate/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incMutex()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesModerate/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readMutex()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkRWMutexMostlyWriteRoutinesModerate(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesModerate)

	for ri := 0; ri < routinesModerate/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incRWMutex()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesModerate/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readRWMutex()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkMutexMostlyWriteRoutinesModerate(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesModerate)

	for ri := 0; ri < routinesModerate/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incMutex()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesModerate/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readMutex()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkRWMutexMostlyReadRoutinesMany(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesMany)

	for ri := 0; ri < routinesMany/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incRWMutex()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesMany/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readRWMutex()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkMutexMostlyReadRoutinesMany(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesMany)

	for ri := 0; ri < routinesMany/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incMutex()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesMany/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readMutex()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkRWMutexMostlyWriteRoutinesMany(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesMany)

	for ri := 0; ri < routinesMany/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incRWMutex()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesMany/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readRWMutex()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkMutexMostlyWriteRoutinesMany(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesMany)

	for ri := 0; ri < routinesMany/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incMutex()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesMany/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readMutex()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
