package jsonwriter

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
