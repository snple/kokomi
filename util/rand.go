package util

import (
	"math/rand"
	"strings"
	"time"
)

func RandString(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[r.Intn(len(chars))])
	}

	return b.String()
}
