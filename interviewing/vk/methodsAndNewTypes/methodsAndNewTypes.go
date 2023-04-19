package main

import "sync"

type myMutex sync.Mutex
type myLocker sync.Locker
type myMutex1 struct{ sync.Mutex }

func main() {
	var mtx myMutex
	mtx.Lock()   // ошибка
	mtx.Unlock() // ошибка

	var mtx1 myMutex1
	mtx1.Lock()   // ok
	mtx1.Unlock() // ok

	var lock myLocker = new(sync.Mutex)
	lock.Lock()   // ok
	lock.Unlock() // ok
}
