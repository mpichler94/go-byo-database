package file

import (
	"errors"
	"go-byo-database/internal/kv"
	"io"
	"os"
)

type Log struct {
	FileName string
	fp       *os.File
}

func (log *Log) Open() (err error) {
	log.fp, err = createFileSync(log.FileName)
	return err
}

func (log *Log) Close() error {
	return log.fp.Close()
}

func (log *Log) Write(ent *kv.Entry) error {
	if _, err := log.fp.Write(ent.Encode()); err != nil {
		return err
	}
	return log.fp.Sync()
}

func (log *Log) Read(ent *kv.Entry) (eof bool, err error) {
	err = ent.Decode(log.fp)
	if err == io.EOF || errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, kv.ErrBadSum) {
		return true, nil
	} else if err != nil {
		return false, err
	} else {
		return false, nil
	}
}
