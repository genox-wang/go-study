package singleton

import "sync"

type LazySingleton struct{}

var _inst *LazySingleton

var mu = sync.Mutex{}

func GetLazySingletonOr() *LazySingleton {
	if _inst == nil {
		mu.Lock()
		if _inst == nil {
			_inst = &LazySingleton{}
		}
		mu.Unlock()
	}
	return _inst
}
