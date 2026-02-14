package internal

import "go-byo-database/internal/kv"

type DB struct {
	KV kv.KV
}

func (db *DB) Open() error  { return db.KV.Open() }
func (db *DB) Close() error { return db.KV.Close() }

func (db *DB) Select(schema *Schema, row Row) (ok bool, err error) {
	key := row.EncodeKey(schema)

	val, ok, err := db.KV.Get(key)
	if err != nil || !ok {
		return ok, err
	}

	return ok, row.DecodeVal(schema, val)
}

func (db *DB) Insert(schema *Schema, row Row) (updated bool, err error) {
	key := row.EncodeKey(schema)
	val := row.EncodeVal(schema)

	return db.KV.SetEx(key, val, kv.ModeInsert)
}

func (db *DB) Upsert(schema *Schema, row Row) (updated bool, err error) {
	key := row.EncodeKey(schema)
	val := row.EncodeVal(schema)

	return db.KV.SetEx(key, val, kv.ModeUpsert)
}

func (db *DB) Update(schema *Schema, row Row) (updated bool, err error) {
	key := row.EncodeKey(schema)
	val := row.EncodeVal(schema)

	return db.KV.SetEx(key, val, kv.ModeUpdate)
}

func (db *DB) Delete(schema *Schema, row Row) (deleted bool, err error) {
	key := row.EncodeKey(schema)

	return db.KV.Del(key)
}
