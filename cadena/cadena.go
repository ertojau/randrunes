//write one function that takes a string and return n bytes
//input: string and n
//output: string

package cadena

import (
	"math/rand"
)

func Cadena(s string, n int) string {
	var b []byte
	for i := 0; i < n; i++ {
		b = append(b, s[rand.Intn(len(s))])
	}
	return string(b)

}
