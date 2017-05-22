package jsonfeed

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/kr/pretty"
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
		if !reflect.DeepEqual(got, test.decoded) {
			t.Errorf("Item.UnmarshalJSON(%q) => %v, want %v", test.encoded, got, test.decoded)
		}
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
		pretty.Ldiff(t, f, saved)
	}
}

func TestMarshalFeedBad(t *testing.T) {
	f := &Feed{} // invalid feed
	_, err := json.Marshal(f)
	if err == nil {
		t.Fatalf("Marshal(%#v) = nil, want error", f)
	}
}
