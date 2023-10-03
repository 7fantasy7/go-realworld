// todo create dir for garbage code, which I just want to have here :)
package singleton

import (
	"sync"
)

type singleton struct {
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
