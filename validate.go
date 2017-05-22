package jsonfeed

import (
	"errors"
)

// Validity returns nil if f is a valid JSON Feed Version 1.
// Otherwise, it returns an error describing at least one
// way in which f is invalid.
func Validity(f *Feed) error {
	for i := range f.Hubs {
		if err := validHub(&f.Hubs[i]); err != nil {
			return err
		}
	}
	for i := range f.Items {
		if err := validItem(&f.Items[i]); err != nil {
			return err
		}
	}
	switch "" {
	case f.Title:
		return errors.New("jsonfeed: no title")
	case f.Version:
		return errors.New("jsonfeed: no version")
	}
	ids := make(map[string]bool)
	for _, item := range f.Items {
		if ids[item.ID] {
			return errors.New("jsonfeed: duplicate id " + item.ID)
		}
		ids[item.ID] = true
	}
	return nil
}

func validHub(h *Hub) error {
	switch "" {
	case h.Type:
		return errors.New("jsonfeed: no type in hub")
	case h.URL:
		return errors.New("jsonfeed: no url in hub")
	}
	return nil
}

func validItem(item *Item) error {
	switch "" {
	case item.ID:
		return errors.New("jsonfeed: no id in item")
	}
	if item.ContentHTML == "" && item.ContentText == "" {
		return errors.New("jsonfeed: no content_html or content_text in item " + item.ID)
	}
	for i := range item.Attachments {
		if err := validAttachment(&item.Attachments[i]); err != nil {
			return err
		}
	}
	return nil
}

func validAttachment(a *Attachment) error {
	switch "" {
	case a.URL:
		return errors.New("jsonfeed: no url in attachment")
	case a.MIMEType:
		return errors.New("jsonfeed: no mime_type in attachment")
	}
	return nil
}
