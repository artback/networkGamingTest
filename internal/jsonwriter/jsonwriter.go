package jsonwriter

import "time"

type JsonWriter interface {
	WriteJSON(v interface{}) error
	WriteControl(messageType int, data []byte, deadline time.Time) error
	Close() error
}
