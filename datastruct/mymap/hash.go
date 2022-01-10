package mymap

import (
	"crypto/sha256"
	"sync"
)

func hashStr(s string) []byte {
	bytes := sha256.New().Sum([]byte(s))
	return bytes

	m := sync.Map{}
	m.Store()
}
