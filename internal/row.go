package internal

import "slices"

type Schema struct {
	Table string
	Cols  []Column
	PKey  []int
}

type Column struct {
	Name string
	Type CellType
}

type Row []Cell

func (schema *Schema) NewRow() Row {
	return make(Row, len(schema.Cols))
}

func (row Row) EncodeKey(schema *Schema) (key []byte) {
	key = append([]byte(schema.Table), 0x00)
	for i := range row {
		if slices.Contains(schema.PKey, i) {
			key = row[i].Encode(key)
		}
	}
	return key
}

func (row Row) EncodeVal(schema *Schema) (val []byte) {
	for i := range row {
		if !slices.Contains(schema.PKey, i) {
			val = row[i].Encode(val)
		}
	}
	return val
}

func (row Row) DecodeKey(schema *Schema, key []byte) (err error) {
	key = key[len(schema.Table)+1:]
	for i := range row {
		if slices.Contains(schema.PKey, i) {
			row[i].Type = schema.Cols[i].Type
			if key, err = row[i].Decode(key); err != nil {
				return err
			}
		}
	}
	return nil
}

func (row Row) DecodeVal(schema *Schema, val []byte) (err error) {
	for i := range row {
		if !slices.Contains(schema.PKey, i) {
			row[i].Type = schema.Cols[i].Type
			if val, err = row[i].Decode(val); err != nil {
				return err
			}
		}
	}
	return nil
}
