package internal

import (
	"encoding/binary"
	"errors"
	"slices"
)

type CellType uint8

const (
	TypeI64 CellType = 1
	TypeStr CellType = 2
)

type Cell struct {
	Type CellType
	I64  int64
	Str  []byte
}

func (cell *Cell) Encode(toAppend []byte) []byte {
	switch cell.Type {
	case TypeI64:
		return binary.LittleEndian.AppendUint64(toAppend, uint64(cell.I64))
	case TypeStr:
		toAppend = binary.LittleEndian.AppendUint32(toAppend, uint32(len(cell.Str)))
		return append(toAppend, cell.Str...)
	default:
		panic("invalid cell type")
	}
}

func (cell *Cell) Decode(data []byte) (rest []byte, err error) {
	switch cell.Type {
	case TypeI64:
		if len(data) < 8 {
			return data, errors.New("expect more data")
		}
		cell.I64 = int64(binary.LittleEndian.Uint64(data))
		return data[8:], nil
	case TypeStr:
		if len(data) < 4 {
			return data, errors.New("expect more data")
		}
		strLen := int(binary.LittleEndian.Uint32(data))
		if len(data) < 4+strLen {
			return data, errors.New("expect more data")
		}
		cell.Str = slices.Clone(data[4 : 4+strLen])
		return data[4+strLen:], nil
	default:
		panic("invalid cell type")
	}
}
