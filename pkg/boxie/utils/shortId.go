// Package utils is a useful utils that boxie use. Methods are usually exported
package utils

import (
	"crypto/rand"
)

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
var namespaceChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-"

// GetShortID makes a short-id
func GetShortID(length int) string {
	ll := len(chars)
	b := make([]byte, length)
	rand.Read(b) // generates len(b) random bytes
	for i := 0; i < length; i++ {
		b[i] = chars[int(b[i])%ll]
	}
	return string(b)
}

// GetShortNamespace makes a string compatable with k8s namespace requirements
func GetShortNamespace(length int) string {
	ll := len(namespaceChars)
	b := make([]byte, length)
	rand.Read(b) // generates len(b) random bytes
	b[0] = 'n'
	for i := 1; i < length-1; i++ {
		b[i] = namespaceChars[int(b[i])%ll]
	}
	b[length-1] = 'e'
	return string(b)
}

func Int32ToString(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
