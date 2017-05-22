package jsonfeed

import (
	"time"
)

// Feed represents a JSON Feed.
// It can be marshaled and unmarshaled with package encoding/json.
type Feed struct {
	Version     string  `json:"version"`
	Title       string  `json:"title"`
	HomePageURL string  `json:"home_page_url,omitempty"`
	FeedURL     string  `json:"feed_url,omitempty"`
	Description string  `json:"description,omitempty"`
	UserComment string  `json:"user_comment,omitempty"`
	NextURL     string  `json:"next_url,omitempty"`
	Icon        string  `json:"icon,omitempty"`
	Favicon     string  `json:"favicon,omitempty"`
	Author      *Author `json:"author,omitempty"`
	Expired     bool    `json:"expired,omitempty"`
	Hubs        []Hub   `json:"hubs,omitempty"`
	Items       []Item  `json:"items"`
}

// Author represents a JSON Feed author.
type Author struct {
	Name   string `json:"name,omitempty"`
	URL    string `json:"url,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

// Hub represents a JSON Feed hub.
type Hub struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// Item represents a JSON Feed item.
type Item struct {
	ID            string       `json:"id"`
	URL           string       `json:"url,omitempty"`
	ExternalURL   string       `json:"external_url,omitempty"`
	Title         string       `json:"title,omitempty"`
	ContentHTML   string       `json:"content_html,omitempty"`
	ContentText   string       `json:"content_text,omitempty"`
	Summary       string       `json:"summary,omitempty"`
	Image         string       `json:"image,omitempty"`
	BannerImage   string       `json:"banner_image,omitempty"`
	DatePublished time.Time    `json:"date_published,omitempty"`
	DateModified  time.Time    `json:"date_modified,omitempty"`
	Author        *Author      `json:"author,omitempty"`
	Tags          []string     `json:"tags,omitempty"`
	Attachments   []Attachment `json:"attachments,omitempty"`
}

// Attachment represents a JSON Feed attachment.
type Attachment struct {
	URL               string `json:"url"`
	MIMEType          string `json:"mime_type"`
	Title             string `json:"title,omitempty"`
	SizeInBytes       int    `json:"size_in_bytes,omitempty"`
	DurationInSeconds int    `json:"duration_in_seconds,omitempty"`
}

// Duration returns the duration stored in DurationInSeconds
// as a Duration.
func (a *Attachment) Duration() time.Duration {
	return time.Duration(a.DurationInSeconds) * time.Second
}
