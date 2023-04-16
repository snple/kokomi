package util

import (
	"bytes"
	"errors"
	"io"

	"github.com/danclive/nson-go"
)

func ReadNsonMessage(reader io.Reader) (nson.Message, error) {
	rb, err := ReadNsonBytes(reader)
	if err != nil {
		return nil, err
	}

	rbuff := bytes.NewBuffer(rb)
	value, err := nson.Message{}.Decode(rbuff)
	if err != nil {
		return nil, err
	}

	return value.(nson.Message), nil
}

func WriteNsonMessage(writer io.Writer, message nson.Message) error {
	wbuff := new(bytes.Buffer)
	err := message.Encode(wbuff)
	if err != nil {
		return err
	}

	_, err = writer.Write(wbuff.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func ReadNsonBytes(reader io.Reader) ([]byte, error) {
	p := make([]byte, 4)
	if _, err := io.ReadFull(reader, p[:]); err != nil {
		return nil, err
	}

	l := int(GetU32(p[:], 0))
	if l < 5 || l > nson.MAX_NSON_SIZE {
		return nil, errors.New("invalid data")
	}

	b := make([]byte, l)
	if _, err := io.ReadFull(reader, b[4:]); err != nil {
		return nil, err
	}

	for i := 0; i < len(p); i++ {
		b[i] = p[i]
	}

	return b, nil
}

func GetU32(b []byte, pos int) uint32 {
	return (uint32(b[pos+0])) |
		(uint32(b[pos+1]) << 8) |
		(uint32(b[pos+2]) << 16) |
		(uint32(b[pos+3]) << 24)
}
