package jsonwriter

import "io"

type Mock struct {
	Err error
}

func (j Mock) WriteJSON(v interface{}) error {
	if j.Err != nil {
		return j.Err
	}
	return nil
}
func (j Mock) Close() error {
	if j.Err != nil {
		return j.Err
	}
	return nil
}

func (j Mock) NextReader() (messageType int, r io.Reader, err error) {
	if j.Err != nil {
		return 0, nil, j.Err
	}
	return 0, nil, nil
}
