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

// UnmarshalJSON has the standard behavior for unmarshaling
// a struct, except for two things:
// it allows an item id to be of any type,
// converting it if necessary to a string,
// as required by the spec,
// and it validates the parsed feed.
func (f *Feed) UnmarshalJSON(b []byte) error {
	v := unmarshalFeed{Feed: f}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	f.Items = make([]Item, 0, len(v.Items))
	for _, item := range v.Items {
		item.Item.ID = string(item.ID)
		f.Items = append(f.Items, *item.Item)
	}
	return Validity(f)
}

type unmarshalFeed struct {
	*Feed
	Items []unmarshalItem
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
