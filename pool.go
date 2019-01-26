package main

import (
	"sync"
)

var (
	sWPool *sync.Pool
)

func init() {
	initsWPool()
}

func initsWPool() {
	sWPool = &sync.Pool{
		New: func() interface{} {
			return new(statusWriter)
		},
	}

}
