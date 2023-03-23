package hashes

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func CalcRdnSHA256() string {
	r := make([]byte, 255)
	_, e := rand.Read(r)
	if e != nil {
		panic(fmt.Sprintf("Error read random: %s", e.Error()))
	}
	return fmt.Sprintf("%x", sha256.Sum256(r))
}
