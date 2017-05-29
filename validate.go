package jsonfeed

import (
	"errors"
)

// validFeed returns nil if f is a valid JSON Feed Version 1.
// Otherwise, it returns an error describing at least one
// way in which f is invalid.
func validFeed(f *Feed) error {
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
	switch {
	case f.Version == "":
		return errors.New("jsonfeed: no version")
	case f.Title == "":
		return errors.New("jsonfeed: no title")
	}
	ids := make(map[string]bool)
	for _, t := range f.Items {
		if ids[t.ID] {
			return errors.New("jsonfeed: duplicate id " + t.ID)
		}
		ids[t.ID] = true
	}
	return validAuthor(f.Author)
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

func validItem(t *Item) error {
	switch "" {
	case t.ID:
		return errors.New("jsonfeed: no id in item")
	}
	if t.ContentHTML == "" && t.ContentText == "" {
		return errors.New("jsonfeed: no content_html or content_text in item " + t.ID)
	}
	for i := range t.Attachments {
		if err := validAttachment(&t.Attachments[i]); err != nil {
			return err
		}
	}
	return validAuthor(t.Author)
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

func validAuthor(a *Author) error {
	if a == nil {
		return nil
	}
	if a.Name == "" && a.URL == "" && a.Avatar == "" {
		return errors.New("jsonfeed: author must provide name or url or avatar")
	}
	return nil
}
