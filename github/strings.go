package github

import (
	"bytes"
	"reflect"
	"sync"
)

var timestampType = reflect.TypeFor[Timestamp]()

var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func Stringify(message any) string { _ = "STUB: not implemented"; return "" }

func stringifyValue(w *bytes.Buffer, val reflect.Value) { _ = "STUB: not implemented"; return }
