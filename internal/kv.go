package internal

import "bytes"

type KV struct {
	log Log
	mem map[string][]byte
}

func (kv *KV) Open() error {
	if err := kv.log.Open(); err != nil {
		return err
	}

	kv.mem = map[string][]byte{}

	ent := Entry{}
	for {
		eof, err := kv.log.Read(&ent)
		if eof {
			break
		}
		if err != nil {
			return err
		}

		if ent.deleted {
			delete(kv.mem, string(ent.key))
		} else {
			kv.mem[string(ent.key)] = ent.val
		}
	}

	return nil
}

func (kv *KV) Close() error { return kv.log.Close() }

func (kv *KV) Get(key []byte) (val []byte, ok bool, err error) {
	val, ok = kv.mem[string(key)]
	return
}

func (kv *KV) Set(key []byte, val []byte) (updated bool, err error) {
	prev, exist := kv.mem[string(key)]
	kv.mem[string(key)] = val
	updated = !exist || !bytes.Equal(prev, val)

	if updated {
		if err := kv.log.Write(&Entry{key, val, false}); err != nil {
			return updated, err
		}
	}

	return
}

func (kv *KV) Del(key []byte) (deleted bool, err error) {
	_, deleted = kv.mem[string(key)]
	delete(kv.mem, string(key))

	if deleted {
		if err := kv.log.Write(&Entry{key, nil, true}); err != nil {
			return deleted, err
		}
	}

	return
}

// QzBQWVJJOUhU https://trialofcode.org/
