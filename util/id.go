package util

import "github.com/danclive/nson-go"

func RandomID() string {
	return nson.NewId().Hex()
}

var ZreoId = nson.Id(make([]byte, 12))

func IdIsZero(id nson.Id) bool {
	for _, n := range id {
		if n != 0 {
			return false
		}
	}

	return true
}
