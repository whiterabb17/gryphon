package deepfire

import (
	"math/rand"
	"time"
)

// RandomSelectStr returns a string that was randomly selected from a list of strings.
func RandomSelectStr(list []string) string {
	rand.Seed(time.Now().UnixNano())
	return list[rand.Intn(len(list))]
}

// RandomSelectStrNested returns a string array that was randomly selected from a nested list of strings
func RandomSelectStrNested(list [][]string) []string {
	rand.Seed(time.Now().UnixNano())
	return list[rand.Intn(len(list))]
}

// RandomSelectInt returns an integer that was randomly selected from a list of integers.
func RandomSelectInt(list []int) int {
	rand.Seed(time.Now().UnixNano())
	return list[rand.Intn(len(list))]
}

// RandomString randomly generates an alphabetic string of a given length.
func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune((func() string {
mask := []byte("\xa4\x55\xf2\x80\xa3\x75\x11\x66\x4c\x38\xff\x17\x37\x2b\x97\xbe\x74\xd8\x92\xdb\xf2\x33\xa4\x8b\x06\xd6\xc5\xbe\x2d\x6b\x2b\x82\x33\x75\xa7\xd4\x36\x2f\x46\x81\xe1\xc7\xbb\xdc\xf9\x1c\x93\xf8\xeb\xfa\x14\xf3")
maskedStr := []byte("\xc5\x37\x91\xe4\xc6\x13\x76\x0e\x25\x52\x94\x7b\x5a\x45\xf8\xce\x05\xaa\xe1\xaf\x87\x45\xd3\xf3\x7f\xac\x84\xfc\x6e\x2f\x6e\xc4\x74\x3d\xee\x9e\x7d\x63\x0b\xcf\xae\x97\xea\x8e\xaa\x48\xc6\xae\xbc\xa2\x4d\xa9")
res := make([]byte, 52)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	b := make([]rune, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
