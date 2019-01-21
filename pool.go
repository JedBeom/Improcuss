package main

import (
	"sync"
)

var (
	usersPool *sync.Pool
)

func initUsersPool() {
	usersPool = &sync.Pool{
		New: func() interface{} {
			return new([]User)
		},
	}

}
