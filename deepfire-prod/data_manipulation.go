package deepfire

import (
	"fmt"
	"math/rand"
	"net"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/c-robinson/iplib"
)

// RemoveFromSlice removes a string from a list of strings if it exists.
func RemoveFromSlice(slice []string, element string) []string {
	res := []string{}

	for _, e := range slice {
		if e != element {
			res = append(res, e)
		}
	}

	return res
}

// CreateWordList generates possible variations of each word in the wordlist.
func CreateWordlist(words []string) []string {
	wordlist := []string{}
	for _, w := range words {
		word := w
		first_to_upper := strings.ToUpper(string(word[0])) + string(word[1:])
		wordlist = append(wordlist, strings.ToUpper(word))
		wordlist = append(wordlist, Revert(word))
		wordlist = append(wordlist, first_to_upper)
		wordlist = append(wordlist, first_to_upper+(func() string {
mask := []byte("\xbc")
maskedStr := []byte("\x8d")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
		wordlist = append(wordlist, first_to_upper+(func() string {
mask := []byte("\xdc\x0f")
maskedStr := []byte("\xed\x3d")
res := make([]byte, 2)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
		wordlist = append(wordlist, first_to_upper+(func() string {
mask := []byte("\xa7\xd4\xad")
maskedStr := []byte("\x96\xe6\x9e")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
		wordlist = append(wordlist, word+(func() string {
mask := []byte("\x15")
maskedStr := []byte("\x24")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
		wordlist = append(wordlist, word+(func() string {
mask := []byte("\x5e\xeb")
maskedStr := []byte("\x6f\xd9")
res := make([]byte, 2)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
		wordlist = append(wordlist, word+(func() string {
mask := []byte("\xcc\xb9\x46")
maskedStr := []byte("\xfd\x8b\x75")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	}

	return wordlist
}

// RemoveStr removes a given string from a list of strings.
func RemoveStr(slice []string, s string) []string {
	final := []string{}
	for _, e := range slice {
		if e != s {
			final = append(final, e)
		}
	}

	return final
}

// RemoveInt removes a given integer from a list of integers.
func RemoveInt(slice []int, s int) []int {
	final := []int{}
	for _, e := range slice {
		if e != s {
			final = append(final, e)
		}
	}

	return final
}

// SplitJoin splits a string then joins them using given delimiters.
func SplitJoin(s, splittBy, joinBy string) string {
	splitted := strings.Split(s, splittBy)
	joined := strings.Join(splitted, joinBy)

	return joined
}

// RevertSlice reverses a slice type agnostically.
func RevertSlice(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func SplitMultiSep(s string, seps []string) []string {
	f := func(c rune) bool {
		for _, sep := range seps {
			if string(c) == sep {
				return true
			}
		}
		return false
	}
	fields := strings.FieldsFunc(s, f)
	return fields
}

func SplitChunks(s string, chunk int) []string {
	if chunk >= len(s) {
		return []string{s}
	}
	var chunks []string
	c := make([]rune, chunk)
	len := 0
	for _, r := range s {
		c[len] = r
		len++
		if len == chunk {
			chunks = append(chunks, string(c))
			len = 0
		}
	}
	if len > 0 {
		chunks = append(chunks, string(c[:len]))
	}
	return chunks
}

// ExtractIntFromString extracts a list of possible integers from a given string.
func ExtractIntFromString(s string) []int {
	res := []int{}
	re := regexp.MustCompile((func() string {
mask := []byte("\x71\xe3\xdd\xfb\x4b\x99\x7d\x5b\x3f\xe4\xa4\x8b\x59\x8e\x7d\x89\xdf\xff\x84\x0d\xae\xca\xa8\x9b\x8e")
maskedStr := []byte("\x2a\xce\x80\xc4\x17\xfd\x26\x07\x5b\xc8\xf9\xa1\x02\xd2\x53\xd4\xe0\xa4\xd8\x69\xd5\xf8\xd5\xc6\xa4")
res := make([]byte, 25)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	submatchall := re.FindAllString(s, -1)

	for _, element := range submatchall {
		res = append(res, StrToInt(element))
	}

	return res
}

// ShuffleSlice randomly shuffles a list of strings.
func ShuffleSlice(s []string) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})

	return s
}

// IpIncrement increments an IP address by 1.
func IpIncrement(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// Revert returns a reversed string.
func Revert(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Contains is used to check if an element exists in an array type agnostically.
func Contains(s interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(s)
	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}
	return false
}

// StrToWords returns a list of strings which was split by spaces.
func StrToWords(s string) []string {
	words := []string{}
	gr := strings.Split(s, (func() string {
mask := []byte("\xf2")
maskedStr := []byte("\xd2")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	for x := range gr {
		z := gr[x]
		if len(z) != 0 {
			words = append(words, z)
		}
	}
	return words
}

// SizeToBytes converts a human friendly string indicating size into a proper integer.
func SizeToBytes(size string) int {
	period_letter := string(size[len(size)-1])
	intr := string(size[:len(size)-1])
	i, _ := strconv.Atoi(intr)
	switch period_letter {
	case (func() string {
mask := []byte("\x31")
maskedStr := []byte("\x56")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):
		return i * 1024 * 1024 * 1024
	case (func() string {
mask := []byte("\xf6")
maskedStr := []byte("\x9b")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):
		return i * 1024 * 1024
	case (func() string {
mask := []byte("\xf2")
maskedStr := []byte("\x99")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):
		return i * 1024
	}
	return i
}

// IntervalToSeconds converts a human friendly string indicating time into a proper integer.
func IntervalToSeconds(interval string) int {
	period_letter := string(interval[len(interval)-1])
	intr := string(interval[:len(interval)-1])
	i, _ := strconv.Atoi(intr)

	switch period_letter {
	case (func() string {
mask := []byte("\x7e")
maskedStr := []byte("\x0d")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):
		return i
	case (func() string {
mask := []byte("\x13")
maskedStr := []byte("\x7e")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):
		return i * 60
	case (func() string {
mask := []byte("\xda")
maskedStr := []byte("\xb2")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):
		return i * 3600
	case (func() string {
mask := []byte("\xeb")
maskedStr := []byte("\x8f")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):
		return i * 24 * 3600
	}
	return i
}

// RemoveNewLines removes possible newlines from a string.
func RemoveNewlines(s string) string {
	re := regexp.MustCompile((func() string {
mask := []byte("\xe2\xed\xcd\xed\x9f")
maskedStr := []byte("\xbe\x9f\xf2\xb1\xf1")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	s = re.ReplaceAllString(s, (func() string {
mask := []byte("\x38")
maskedStr := []byte("\x18")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	return s
}

// FullRemove removes all instances of a string from another string.
func FullRemove(str string, to_remove string) string {
	return strings.Replace(str, to_remove, (func() string {
mask := []byte("")
maskedStr := []byte("")
res := make([]byte, 0)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), -1)
}

// RemoveDuplicatesStr returns an array of strings that are unique to each other.
func RemoveDuplicatesStr(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// RemoveDuplicatesInt returns an array of integers that are unique to each other.
func RemoveDuplicatesInt(slice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// ContainsAny checks if a string exists within a list of strings.
func ContainsAny(str string, elements []string) bool {
	for element := range elements {
		e := elements[element]
		if strings.Contains(str, e) {
			return true
		}
	}

	return false
}

// Convert an IPv4 address to hex
func IP2Hex(ip string) string {
	ip_obj := net.ParseIP(ip)
	return iplib.IPToHexString(ip_obj)
}

// Convert a port to hex
func Port2Hex(port int) string {
	hexval := fmt.Sprintf((func() string {
mask := []byte("\xc6\x9e\x7f\x7a")
maskedStr := []byte("\xf6\xe6\x5a\x02")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), port)
	hexval_without_prefix := FullRemove(hexval, (func() string {
mask := []byte("\xa1\x83")
maskedStr := []byte("\x91\xfb")
res := make([]byte, 2)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	two_bytes_slice := SplitChunks(hexval_without_prefix, 2)
	return fmt.Sprintf((func() string {
mask := []byte("\x4e\xda\x01\xe3\x35\x41")
maskedStr := []byte("\x7e\xa2\x24\x90\x10\x32")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), two_bytes_slice[1], two_bytes_slice[0])
}
