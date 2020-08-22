package jsonwriter

import "io"

type JsonWriter interface {
	WriteJSON(v interface{}) error
	Close() error
	NextReader() (messageType int, r io.Reader, err error)
}
