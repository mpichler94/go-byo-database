package internal

type KV struct {
	mem map[string][]byte
}

func (kv *KV) Open() error {
	kv.mem = map[string][]byte{} // empty
	return nil
}

func (kv *KV) Close() error { return nil }

func (kv *KV) Get(key []byte) (val []byte, ok bool, err error) {
	keyStr := string(key)
	val = kv.mem[keyStr]
	ok = val != nil

	return val, ok, nil
}

func (kv *KV) Set(key []byte, val []byte) (updated bool, err error) {
	keyStr := string(key)
	updated = true
	kv.mem[keyStr] = val

	return updated, nil
}

func (kv *KV) Del(key []byte) (deleted bool, err error) {
	keyStr := string(key)
	deleted = kv.mem[keyStr] != nil
	delete(kv.mem, keyStr)

	return deleted, nil
}

// QzBQWVJJOUhU https://trialofcode.org/
