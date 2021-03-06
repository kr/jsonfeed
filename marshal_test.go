package jsonfeed

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestAnyString(t *testing.T) {
	cases := []struct{ encoded, decoded string }{
		{`"s"`, `s`},
		{`true`, `true`},
		{`12345`, `12345`},
		{`[1, 2]`, `[1, 2]`},
		{`{"a": 1}`, `{"a": 1}`},
	}

	for _, test := range cases {
		var got anyString
		err := got.UnmarshalJSON([]byte(test.encoded))
		if err != nil {
			t.Errorf("anyString.UnmarshalJSON(%q) = %v, want nil", test.encoded, err)
			continue
		}
		want := anyString(test.decoded)
		if got != want {
			t.Errorf("anyString.UnmarshalJSON(%q) => %q, want %q", test.encoded, got, want)
		}
	}
}

func TestUnmarshalItemOk(t *testing.T) {
	cases := []struct {
		encoded string
		decoded Item
	}{
		{
			`{"id": "id", "content_text": "text"}`,
			Item{ID: "id", ContentText: "text"},
		},
		{
			`{"id": 12345, "content_text": "text"}`,
			Item{ID: "12345", ContentText: "text"},
		},
	}

	for _, test := range cases {
		var got Item
		err := json.Unmarshal([]byte(test.encoded), &got)
		if err != nil {
			t.Errorf("Item.UnmarshalJSON(%q) = %v, want nil", test.encoded, err)
			continue
		}
		got.DatePublished = time.Time{} // tested below
		got.DateModified = time.Time{}  // tested below
		if !reflect.DeepEqual(got, test.decoded) {
			t.Errorf("Item.UnmarshalJSON(%q) => %v, want %v", test.encoded, got, test.decoded)
		}
	}
}

func TestUnmarshalItemDateExplicit(t *testing.T) {
	date := time.Date(1985, 10, 26, 1, 21, 0, 0, time.FixedZone("", -28800))
	b := []byte(`{
		"id": "id",
		"content_text": "text",
		"date_published": "1985-10-26T01:21:00-08:00",
		"date_modified": "1985-10-26T01:21:00-08:00"
	}`)
	var got Item
	err := json.Unmarshal(b, &got)
	if err != nil {
		t.Fatalf("Item.UnmarshalJSON(%q) = %v, want nil", b, err)
	}

	want := Item{
		ID:            "id",
		ContentText:   "text",
		DatePublished: date,
		DateModified:  date,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Item.UnmarshalJSON(%q) => %v, want %v", b, got, want)
	}
}

func TestUnmarshalItemDateDefault(t *testing.T) {
	b := []byte(`{"id": "id", "content_text": "text"}`)
	var got Item
	before := time.Now().UTC()
	err := json.Unmarshal(b, &got)
	after := time.Now().UTC()
	if err != nil {
		t.Fatalf("Item.UnmarshalJSON(%q) = %v, want nil", b, err)
	}
	if g := got.DatePublished; g.Before(before) || g.After(after) {
		t.Errorf("Item.UnmarshalJSON(%q) => bad date_published %v", b, g)
	}
	if g := got.DateModified; g.Before(before) || g.After(after) {
		t.Errorf("Item.UnmarshalJSON(%q) => bad date_modified %v", b, g)
	}
}

func TestUnmarshalItemBad(t *testing.T) {
	b := []byte(`xxx`) // invalid JSON
	var got Item
	err := got.UnmarshalJSON(b)
	if err == nil {
		t.Fatalf("Item.UnmarshalJSON(%q) = nil, want error", b)
	}
}

func TestUnmarshalFeedBad(t *testing.T) {
	b := []byte(`xxx`) // invalid JSON
	var got Feed
	err := got.UnmarshalJSON(b)
	if err == nil {
		t.Fatalf("Feed.UnmarshalJSON(%q) = nil, want error", b)
	}
}

func TestMarshalFeedOk(t *testing.T) {
	f := &Feed{
		Version: Version,
		Title:   "title",
		Items: []Item{{
			ID:          "id",
			ContentText: "text",
		}},
	}
	got, err := json.Marshal(f)
	if err != nil {
		t.Fatalf("Marshal(%#v) = %v, want nil", f, err)
	}
	want := []byte(`{"version":"https://jsonfeed.org/version/1","title":"title","items":[{"id":"id","content_text":"text","date_published":"0001-01-01T00:00:00Z","date_modified":"0001-01-01T00:00:00Z"}]}`)
	if !bytes.Equal(got, want) {
		t.Errorf("Marshal(%#v) => %#q, want %#q", f, got, want)
	}
}

func TestMarshalFeedFixups(t *testing.T) {
	f := &Feed{
		Version: "", // test replacing version string
		Title:   "title",
		Items:   nil, // test encoding as "[]"
	}
	saved := new(Feed)
	*saved = *f
	got, err := json.Marshal(f)
	if err != nil {
		t.Fatalf("Marshal(%#v) = %v, want nil", f, err)
	}
	want := []byte(`{"version":"https://jsonfeed.org/version/1","title":"title","items":[]}`)
	if !bytes.Equal(got, want) {
		t.Errorf("Marshal(%#v) => %#q, want %#q", f, got, want)
	}
	if !reflect.DeepEqual(f, saved) {
		t.Error("MarshalJSON mutated f")
	}
}

func TestMarshalFeedBad(t *testing.T) {
	f := &Feed{} // invalid feed
	_, err := json.Marshal(f)
	if err == nil {
		t.Fatalf("Marshal(%#v) = nil, want error", f)
	}
}
