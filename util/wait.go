package util

import "sync"

func Block()  {
	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}
