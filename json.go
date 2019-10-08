package dobaos

import "encoding/json"

type jsonWrapper struct{}

func (jw *jsonWrapper) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (jw *jsonWrapper) Unmarshal(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}
