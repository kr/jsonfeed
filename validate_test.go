package jsonfeed

import "testing"

func TestValidFeed(t *testing.T) {
	// a valid feed using all features that have validation rules
	f := &Feed{
		Version: "https://jsonfeed.org/version/1",
		Title:   "title",
		Hubs: []Hub{
			{Type: "type", URL: "url"},
		},
		Items: []Item{
			{
				ID:          "id",
				ContentText: "text",
				Attachments: []Attachment{
					{URL: "url", MIMEType: "mimetype"},
				},
			},
		},
	}
	err := Validity(f)
	if err != nil {
		t.Errorf("Validity(%#v) = %v, want nil", f, err)
	}
}

func TestInvalidFeed(t *testing.T) {
	cases := []*Feed{
		{Title: "title"},                            // no version
		{Version: "https://jsonfeed.org/version/1"}, // no title
		{
			Version: "https://jsonfeed.org/version/1",
			Title:   "title",
			Items: []Item{
				{ID: "id", ContentText: "text"},
				{ID: "id", ContentText: "text"}, // dup item id
			},
		},
		{
			Version: "https://jsonfeed.org/version/1",
			Title:   "title",
			Items:   []Item{{}}, // invalid item
		},
		{
			Version: "https://jsonfeed.org/version/1",
			Title:   "title",
			Hubs:    []Hub{{}}, // invalid hub
		},
	}

	for _, test := range cases {
		err := Validity(test)
		if err == nil {
			t.Errorf("Validity(%v) = nil, want error", test)
		}
	}
}

func TestInvalidItem(t *testing.T) {
	cases := []*Item{
		{},         // no id
		{ID: "id"}, // no content
		{
			ID:          "id",
			ContentText: "text",
			Attachments: []Attachment{{}}, // invalid attachment
		},
	}

	for _, test := range cases {
		err := validItem(test)
		if err == nil {
			t.Errorf("validItem(%v) = nil, want error", test)
		}
	}
}

func TestInvalidAttachment(t *testing.T) {
	cases := []*Attachment{
		{MIMEType: "mimetype"}, // no url
		{URL: "url"},           // no mime_type
	}

	for _, test := range cases {
		err := validAttachment(test)
		if err == nil {
			t.Errorf("validAttachment(%v) = nil, want error", test)
		}
	}
}

func TestInvalidHub(t *testing.T) {
	cases := []*Hub{
		{Type: "type"}, // no url
		{URL: "url"},   // no type
	}

	for _, test := range cases {
		err := validHub(test)
		if err == nil {
			t.Errorf("validHub(%v) = nil, want error", test)
		}
	}
}

func TestInvalidAuthor(t *testing.T) {
	a := &Author{}
	err := validAuthor(a)
	if err == nil {
		t.Errorf("validAuthor(%v) = nil, want error", a)
	}
}
