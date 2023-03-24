package util

import "github.com/danclive/nson-go"

func RandomID() string {
	return nson.NewMessageId().Hex()
}

var ZreoMessageId = nson.MessageId(make([]byte, 12))

func MessageIdIsZero(id nson.MessageId) bool {
	for _, n := range id {
		if n != 0 {
			return false
		}
	}

	return true
}
