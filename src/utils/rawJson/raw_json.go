package rawJson

import "errors"

type RawJson []byte

// MarshalJSON returns m as the JSON encoding of m.
func (r RawJson) MarshalJSON() ([]byte, error) {
	return r, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (r *RawJson) UnmarshalJSON(data []byte) error {
	if r == nil {
		return errors.New("UnmarshalJSON on nil pointer")
	}
	*r = data
	return nil
}

type RawJsonStr string

// MarshalJSON returns m as the JSON encoding of m.
func (r RawJsonStr) MarshalJSON() ([]byte, error) {
	return []byte(r), nil
}

// UnmarshalJSON sets *m to a copy of data.
func (r *RawJsonStr) UnmarshalJSON(data []byte) error {
	if r == nil {
		return errors.New("UnmarshalJSON on nil pointer")
	}
	*r = RawJsonStr(data)
	return nil
}
