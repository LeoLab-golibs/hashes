package hashes

import (
	"crypto/rand"
	"fmt"
)

func GetRandBytes(cnt int) []byte {
	r := make([]byte, cnt)
	_, e := rand.Read(r)
	if e != nil {
		panic(fmt.Sprintf("Error read random: %s", e.Error()))
	}
	return r
}
