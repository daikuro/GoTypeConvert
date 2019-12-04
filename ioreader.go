package typeconv

import (
	"bytes"
	"io"
	"os"
)

type IoReaderValue struct {
	A io.Reader
}

func ToIoReader(v interface{}) (r *IoReaderValue) {
	r = &IoReaderValue{}
	if v == nil {
		return
	}
	switch t := v.(type) {
	case string:
		r.A = bytes.NewBufferString(t)
		return
	case []byte:
		r.A = bytes.NewBuffer(t)
		return
	case *os.File:
		r.A = t
		return
	}

	return
}
