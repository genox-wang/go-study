package singleton

type HungrySingleton struct{}

var inst = &HungrySingleton{}

func GetInstOr() *HungrySingleton {
	return inst
}
