package bruteforce

import (
	"crypto/md5"
	"log"
	"math/rand"
	"testing"
)

var guess = []byte("0000")

func Spawner() func() ([]byte, bool) {
	status := false
	return func() ([]byte, bool) {
		defer func() {
			for i := len(guess) - 1; i >= 0; i-- {
				if guess[i] == '9' {
					guess[i] = 'a'
				} else if guess[i] == 'z' {
					guess[i] = 'A'
				} else if guess[i] == 'Z' {
					guess[i] = '0'
				} else {
					guess[i]++
					break
				}
			}
		}()
		return append([]byte("123123"), guess...), status
	}
}

var hash [16]byte

func Check(b []byte) bool {
	return md5.Sum(b) == hash
}

func BenchmarkBruteForce_Brute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rd := rand.Uint64()
		toapp := make([]byte, 0, 4)
		dict := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for i := 0; i < 4; i++ {
			toapp = append(toapp, dict[rd%uint64(len(dict))])
			rd /= uint64(len(dict))
		}
		hash = md5.Sum(append([]byte("123123"), toapp...))
		bf := New(Spawner(), Check, 16)
		b.ResetTimer()
		val, ok := bf.Brute()
		if ok {
			log.Println(string(val))
		}
	}

}
