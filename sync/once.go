package sync

import (
	"fmt"
	"sync"
)

var m = make(map[string]*sync.Once)

func Once(f func()) {

	s := fmt.Sprint(f)

	if once, existed := m[s]; existed {
		once.Do(f)
	} else {
		once := sync.Once{}
		m[s] = &once
		once.Do(f)
	}

}
