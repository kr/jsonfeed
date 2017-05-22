package jsonfeed

import "encoding/json"

// MarshalJSON has the standard behavior for marshaling a struct,
// except it validates f before marshaling.
func (f *Feed) MarshalJSON() ([]byte, error) {
	err := Validity(f)
	if err != nil {
		return nil, err
	}
	if f.Items == nil {
		f.Items = make([]Item, 0) // avoid emitting JSON 'null'
	}
	type t Feed // get rid of method MarshalJSON to avoid recursion
	return json.Marshal((*t)(f))
}

// UnmarshalJSON has the standard behavior for unmarshaling a struct,
// except that it validates the parsed feed.
func (f *Feed) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, f)
	if err != nil {
		return err
	}
	return Validity(f)
}

// UnmarshalJSON has the standard behavior for unmarshaling a struct,
// except that it allows the id to be of any type,
// converting it if necessary to a string,
// as required by the spec.
func (t *Item) UnmarshalJSON(b []byte) error {
	v := unmarshalItem{Item: t}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	t.ID = string(v.ID)
	return nil
}

type unmarshalItem struct {
	*Item
	ID anyString
}

type anyString string

func (s *anyString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, (*string)(s))
	if err != nil {
		// Not a valid JSON string, so fall back to
		// the JSON text representation.
		*s = anyString(b)
	}
	return nil
}
