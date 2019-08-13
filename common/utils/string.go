package utils

import (
	"math/rand"
	"time"
	"unsafe"
)

var src = rand.NewSource(time.Now().UnixNano())

var letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandString(n, letterType int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	switch {
	case letterType == 1:
		letterBytes = "0123456789abcdefghijklmnopqrstuvwxyz"
	case letterType == 2:
		letterBytes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case letterType == 3:
		letterBytes = "abcdefghijklmnopqrstuvwxyz"
	case letterType == 4:
		letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
