package jsonwriter

type JsonWriter interface {
	WriteJSON(v interface{}) error
	Close() error
}
