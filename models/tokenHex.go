package models

import (
	"crypto/rand"
	"encoding/hex"
	mrand "math/rand"
	"strconv"
	"time"
)

//RandomHex generate and return a random hex string
func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func RandomDigits(length int) string {
        digits := ""
        mrand.Seed(time.Now().UnixNano())
        for i:=0; i<length; i++ {
                digits += strconv.Itoa(mrand.Intn(9))
        }
        return digits
}
