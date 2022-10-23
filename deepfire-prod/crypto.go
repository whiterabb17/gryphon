package deepfire

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	crand "crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateKey() []byte {
	random_bytes := make([]byte, 32)
	_, err := crand.Read(random_bytes) // Generates 32 cryptographically secure random bytes
	if err != nil {
		println((func() string {
mask := []byte("\xfb\xc8\xf0\x08\x25\xe6\x4c\x27\x62\xfa\xc9\xfe\x63\xe2\x42\x03\x7a\x8b\xc4\x03\x69\x6e\x07\x33\x42\x37\x10")
maskedStr := []byte("\xbd\xa9\x99\x64\x40\x82\x6c\x53\x0d\xda\xae\x9b\x0d\x87\x30\x62\x0e\xee\xe4\x77\x01\x0b\x27\x58\x27\x4e\x3e")
res := make([]byte, 27)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
		return nil
	}
	return random_bytes
}

func GenerateIV() []byte {
	random_bytes := make([]byte, 16)
	_, err := crand.Read(random_bytes) // Generates 16 cryptographically secure random bytes
	if err != nil {
		println((func() string {
mask := []byte("\x4c\x5e\xc5\x64\x2c\xa3\xc1\xc2\x55\x0a\x87\x16\xa9\x28\xe7\xcd\xba\xea\xc9\xbf\x25\x00")
maskedStr := []byte("\x0a\x3f\xac\x08\x49\xc7\xe1\xb6\x3a\x2a\xe0\x73\xc7\x4d\x95\xac\xce\x8f\xe9\xf6\x73\x2e")
res := make([]byte, 22)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
		return nil
	}
	return random_bytes
}

func EncryptBytes(secret_message []byte, key []byte) []byte {
	cipher_block, err := aes.NewCipher(key)
	if err != nil {
		println((func() string {
mask := []byte("\x98\x49\x45\xcc\xe5\xda\x4c\x6f\xb4\x82\x9d\xbb\x66\x9f\x61\x44\x40\xf7\x19\x2e\x14\xea\x5f\xba\x49\x2c\x95\x56")
maskedStr := []byte("\xdd\x3b\x37\xa3\x97\xfa\x23\x0c\xd7\xf7\xef\xde\x02\xb3\x41\x27\x21\x99\x3e\x5a\x34\x8f\x31\xd9\x3b\x55\xe5\x22")
res := make([]byte, 28)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
		return nil
	}

	length_to_bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(length_to_bytes, uint32(len(secret_message)))

	length_and_secret := append(length_to_bytes, secret_message...)

	IV := GenerateIV()
	if len(length_and_secret)%16 != 0 {
		appending := make([]byte, (16 - len(length_and_secret)%16))
		corrected := append(length_and_secret, appending...)
		length_and_secret = corrected
	}

	c := cipher.NewCBCEncrypter(cipher_block, IV)
	encrypted := make([]byte, len(length_and_secret))
	c.CryptBlocks(encrypted, length_and_secret)

	return append(IV, encrypted...)
}

func DecryptBytes(encrypted_message []byte, key []byte) []byte {
	IV := encrypted_message[0:16]

	actual_ciphertext := encrypted_message[16:]

	cipher_block, err := aes.NewCipher(key)
	if err != nil {
		println((func() string {
mask := []byte("\xd7\x40\x88\xa9\x13\x2e\x4a\x5c\x13\xcc\x15\x9a\xa6\xeb\x4a\x0e\x9b\x96\x3c\xad\xd1\x6e\x9c\x2d\xba\xfd\xce\x26")
maskedStr := []byte("\x92\x32\xfa\xc6\x61\x0e\x25\x3f\x70\xb9\x67\xff\xc2\xc7\x6a\x6d\xfa\xf8\x1b\xd9\xf1\x0a\xf9\x4e\xc8\x84\xbe\x52")
res := make([]byte, 28)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	}
	c := cipher.NewCBCDecrypter(cipher_block, IV)
	decrypted := make([]byte, len(actual_ciphertext))
	c.CryptBlocks(decrypted, actual_ciphertext)

	length_bytes := decrypted[0:4]
	length := binary.LittleEndian.Uint32(length_bytes)
	decrypted = decrypted[4:]
	return decrypted[:length]
}

func EncryptString(message string, key []byte) []byte {
	return DecryptBytes([]byte(message), key)
}

func DecryptString(encrypted_message []byte, key []byte) string {
	return string(DecryptBytes(encrypted_message, key))
}

// MD5Hash hashes a given string using the MD5.
func Md5Hash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))

	return hex.EncodeToString(hasher.Sum(nil))
}

//SHA1Hash hashes a given string using the SHA1.
func Sha1Hash(str string) string {
	hasher := sha1.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Sha256Hash(str string) string {
	hasher := sha256.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

// B64D decodes a given string encoded in Base64.
func B64D(str string) string {
	raw, _ := base64.StdEncoding.DecodeString(str)

	return fmt.Sprintf((func() string {
mask := []byte("\xc7\x18")
maskedStr := []byte("\xe2\x6b")
res := make([]byte, 2)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), raw)
}

// B64E encodes a string in Base64.
func B64E(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Rot13(str string) string {
	var finaldata bytes.Buffer
	for _, character := range str {
		if character >= 'a' && character <= 'z' {
			if character > 'm' {
				character_tmp := character - 13
				finaldata.WriteString(string(character_tmp))
			} else if character == 'm' {
				character_tmp := 'z'
				finaldata.WriteString(string(character_tmp))
			} else if character == 'z' {
				character_tmp := 'm'
				finaldata.WriteString(string(character_tmp))
			} else {
				character_tmp := character + 13
				finaldata.WriteString(string(character_tmp))
			}
		} else if character >= 'A' && character <= 'Z' {
			if character > 'M' {
				character_tmp := character - 13
				finaldata.WriteString(string(character_tmp))
			} else if character == 'M' {
				character_tmp := 'Z'
				finaldata.WriteString(string(character_tmp))
			} else if character == 'Z' {
				character_tmp := 'M'
				finaldata.WriteString(string(character_tmp))
			} else {
				character_tmp := character + 13
				finaldata.WriteString(string(character_tmp))
			}
		}
	}
	return finaldata.String()
}

func UnixToTime(time_num int64) string {
	return time.Unix(time_num, 0).String()
}
