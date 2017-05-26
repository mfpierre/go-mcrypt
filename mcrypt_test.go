package mcrypt

import "testing"
import "reflect"

type encryptionTest struct {
	algo      string
	mode      string
	key       string
	ivSize    int
	plaintext string
	encrypted []byte
}

var tests = []encryptionTest{
	{"cast-256", "ecb", "here is a random key of 32 bytes", 16, "hello", []byte{53, 190, 73, 3, 192, 5, 42, 202, 55, 77, 85, 218, 189, 169, 253, 147}},
	{"rijndael-256", "cbc", "here is a random key of 32 bytes", 32, "hello", []byte{196, 120, 32, 94, 41, 142, 113, 20, 109, 123, 95, 30, 255, 227, 63, 219, 32, 175, 35, 54, 65, 24, 130, 254, 2, 88, 201, 226, 1, 2, 235, 252}},
}

func TestEncrypt(t *testing.T) {
	for _, pair := range tests {
		key := []byte(pair.key) // 32 bytes
		plaintext := []byte(pair.plaintext)
		iv := make([]byte, pair.ivSize)

		s, _ := encrypt(key, iv, plaintext, pair.algo, pair.mode)
		if !reflect.DeepEqual(s, pair.encrypted) {
			t.Error(
				"For", pair.plaintext,
				"expected", pair.encrypted,
				"got", s,
			)
		}
	}
}
