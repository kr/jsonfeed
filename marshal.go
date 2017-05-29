package jsonfeed

import (
	"encoding/json"
	"time"
)

// MarshalJSON has the standard behavior for marshaling a struct,
// except it validates f before marshaling.
// It always emits the version in Version,
// regardless of the value in f.
func (f *Feed) MarshalJSON() ([]byte, error) {
	// TODO(kr): avoid copying all of f
	f1 := new(Feed)
	*f1 = *f
	f1.Version = Version
	err := validFeed(f1)
	if err != nil {
		return nil, err
	}
	if f1.Items == nil {
		f1.Items = make([]Item, 0) // avoid emitting JSON 'null'
	}
	type t Feed // get rid of method MarshalJSON to avoid recursion
	return json.Marshal((*t)(f1))
}

// UnmarshalJSON has the standard behavior for unmarshaling a struct,
// except that it validates the parsed feed.
func (f *Feed) UnmarshalJSON(b []byte) error {
	type t Feed // get rid of method UnmarshalJSON to avoid recursion
	err := json.Unmarshal(b, (*t)(f))
	if err != nil {
		return err
	}
	return validFeed(f)
}

// UnmarshalJSON has the standard behavior for unmarshaling a struct,
// except that it allows the id to be of any type,
// converting it if necessary to a string,
// as required by the spec,
// and it replaces missing dates with the current time.
func (t *Item) UnmarshalJSON(b []byte) error {
	type T Item // get rid of method UnmarshalJSON to avoid recursion
	v := struct {
		*T
		ID anyString `json:"id"`
	}{T: (*T)(t)}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	t.ID = string(v.ID)
	if t.DatePublished.IsZero() {
		t.DatePublished = time.Now().UTC()
	}
	if t.DateModified.IsZero() {
		t.DateModified = time.Now().UTC()
	}
	return nil
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
