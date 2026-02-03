package internal

import (
	"encoding/binary"
	"io"
)

type Entry struct {
	key []byte
	val []byte
}

func (ent *Entry) Encode() []byte {
	data := make([]byte, 4+4+len(ent.key)+len(ent.val))
	binary.LittleEndian.PutUint32(data[0:4], uint32(len(ent.key)))
	binary.LittleEndian.PutUint32(data[4:8], uint32(len(ent.val)))
	copy(data[8:], ent.key)
	copy(data[8+len(ent.key):], ent.val)

	return data
}

func (ent *Entry) Decode(r io.Reader) error {
	size := make([]byte, 8)
	if _, err := r.Read(size); err != nil {
		return err
	}

	kLen := binary.LittleEndian.Uint32(size[:4])
	vLen := binary.LittleEndian.Uint32(size[4:])
	data := make([]byte, kLen+vLen)

	if _, err := r.Read(data); err != nil {
		return err
	}

	ent.key = data[:kLen]
	ent.val = data[kLen:]
	return nil
}

// QzBQWVJJOUhU https://trialofcode.org/
