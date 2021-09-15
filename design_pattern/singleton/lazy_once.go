package singleton

import "sync"

type LazyOnceSingleton struct{}

var (
	_onceInst *LazyOnceSingleton
	once      = sync.Once{}
)

func GetLazyOnceSingletonOr() *LazyOnceSingleton {
	once.Do(func() {
		if _onceInst == nil {
			_onceInst = &LazyOnceSingleton{}
		}
	})
	return _onceInst
}
