package jsonfeed

import (
	"time"
)

// Version is the version of JSON Feed
// generated and recognized by this package.
// It is automatically inserted into a feed when marshaling,
// and examined when checking validity.
// Future editions of this package that support subsequent versions
// of the spec will recognize all past versions;
// since this is currently the only version, it must match exactly.
const Version = "https://jsonfeed.org/version/1"

// Feed represents a JSON Feed.
// It can be marshaled and unmarshaled with package encoding/json.
//
// Documentation for each field is taken from the
// JSON Feed spec.
type Feed struct {
	// Version is the URL of the version of the format the
	// feed uses. This should appear at the very top, though
	// we recognize that not all JSON generators allow for
	// ordering.
	Version string `json:"version"`

	// Title is the name of the feed, which will often
	// correspond to the name of the website (blog, for
	// instance), though not necessarily.
	Title string `json:"title"`

	// HomePageURL (strongly recommended) is the URL of the
	// resource that the feed describes. This resource may
	// or may not actually be a “home” page, but it should
	// be an HTML page. If a feed is published on the public
	// web, this should be considered as required. But it
	// may not make sense in the case of a file created on a
	// desktop computer, when that file is not shared or is
	// shared only privately.
	HomePageURL string `json:"home_page_url,omitempty"`

	// FeedURL (strongly recommended) is the URL of the
	// feed, and serves as the unique identifier for the
	// feed. As with HomePageURL, this should be considered
	// required for feeds on the public web.
	FeedURL string `json:"feed_url,omitempty"`

	// Description provides more detail, beyond the  title,
	// on what the feed is about. A feed reader may display
	// this text.
	Description string `json:"description,omitempty"`

	// UserComment is a description of the purpose of the
	// feed. This is for the use of people looking at the
	// raw JSON, and should be ignored by feed readers.
	UserComment string `json:"user_comment,omitempty"`

	// NextURL is the URL of a feed that provides the next n
	// items, where n is determined by the publisher. This
	// allows for pagination, but with the expectation that
	// reader software is not required to use it and
	// probably won’t use it very often. NextURL must not
	// be the same as FeedURL, and it must not be the same
	// as a previous NextURL (to avoid infinite loops).
	NextURL string `json:"next_url,omitempty"`

	// Icon is the URL of an image for the feed suitable to
	// be used in a timeline, much the way an avatar might
	// be used. It should be square and relatively large —
	// such as 512 x 512 — so that it can be scaled-down and
	// so that it can look good on retina displays. It
	// should use transparency where appropriate, since it
	// may be rendered on a non-white background.
	Icon string `json:"icon,omitempty"`

	// Favicon is the URL of an image for the feed suitable
	// to be used in a source list. It should be square and
	// relatively small, but not smaller than 64 x 64 (so
	// that it can look good on retina displays). As with
	// icon, this image should use transparency where
	// appropriate, since it may be rendered on a non-white
	// background.
	Favicon string `json:"favicon,omitempty"`

	// Author is the author of the feed.
	Author *Author `json:"author,omitempty"`

	// Expired says whether or not the feed is finished —
	// that is, whether or not it will ever update again. A
	// feed for a temporary event, such as an instance of
	// the Olympics, could expire. If the value is true,
	// then it’s expired. Any other value, or the absence of
	// expired, means the feed may continue to update.
	Expired bool `json:"expired,omitempty"`

	// Hubs describes endpoints that can be used to
	// subscribe to real-time notifications from the
	// publisher of this feed.
	Hubs []Hub `json:"hubs,omitempty"`

	// Items contains the items in the feed.
	Items []Item `json:"items"`
}

// Author represents a JSON Feed author.
// Its fields are all optional — but if you provide an author
// object, then at least one is required.
//
// Documentation for each field is taken from the
// JSON Feed spec.
type Author struct {
	// Name is the author’s name.
	Name string `json:"name,omitempty"`

	// URL is the URL of a site owned by the author. It
	// could be a blog, micro-blog, Twitter account, and so
	// on. Ideally the linked-to page provides a way to
	// contact the author, but that’s not required. The URL
	// could be a mailto: link, though we suspect that will
	// be rare.
	URL string `json:"url,omitempty"`

	// Avatar is the URL for an image for the author. As
	// with icon, it should be square and relatively large —
	// such as 512 x 512 — and should use transparency where
	// appropriate, since it may be rendered on a non-white
	// background.
	Avatar string `json:"avatar,omitempty"`
}

// Hub represents a JSON Feed hub.
//
// Documentation for each field is taken from the
// JSON Feed spec.
type Hub struct {
	// Type describes the protocol used to talk with the
	// hub, such as “rssCloud” or “WebSub.” When using
	// WebSub, the value for the JSON Feed’s FeedURL is
	// passed for the hub.topic parameter.
	Type string `json:"type"`

	// URL specifies the location of the hub.
	URL string `json:"url"`
}

// Item represents a JSON Feed item.
//
// Documentation for each field is taken from the
// JSON Feed spec.
type Item struct {
	// ID is unique for that item for that feed over time.
	// If an item is ever updated, the id should be
	// unchanged. New items should never use a
	// previously-used id. If an id is presented as a number
	// or other type, a JSON Feed reader must coerce it to a
	// string. Ideally, the id is the full URL of the
	// resource described by the item, since URLs make great
	// unique identifiers.
	ID string `json:"id"`

	// URL is the URL of the resource described by the item.
	// It’s the permalink. This may be the same as the id —
	// but should be present regardless.
	URL string `json:"url,omitempty"`

	// ExternalURL is the URL of a page elsewhere. This is
	// especially useful for linkblogs. If url links to
	// where you’re talking about a thing, then ExternalURL
	// links to the thing you’re talking about.
	ExternalURL string `json:"external_url,omitempty"`

	// Title is plain text. Microblog items in particular
	// may omit titles.
	Title string `json:"title,omitempty"`

	// ContentHTML is the HTML of the item. Important: the
	// only place HTML is allowed in this format is in
	// ContentHTML.
	//
	// ContentHTML and ContentText are each optional strings
	// — but one or both must be present. A Twitter-like
	// service might use ContentText, while a blog might use
	// ContentHTML. Use whichever makes sense for your
	// resource. (It doesn’t even have to be the same for
	// each item in a feed.)
	ContentHTML string `json:"content_html,omitempty"`

	// ContentText is the plain text of the item.
	//
	// ContentHTML and ContentText are each optional strings
	// — but one or both must be present. A Twitter-like
	// service might use ContentText, while a blog might use
	// ContentHTML. Use whichever makes sense for your
	// resource. (It doesn’t even have to be the same for
	// each item in a feed.)
	ContentText string `json:"content_text,omitempty"`

	// Summary is a plain text sentence or two describing
	// the item. This might be presented in a timeline, for
	// instance, where a detail view would display all of
	// ContentHTML or ContentText.
	Summary string `json:"summary,omitempty"`

	// Image is the URL of the main image for the item. This
	// image may also appear in the ContentHTML — if so,
	// it’s a hint to the feed reader that this is the main,
	// featured image. Feed readers may use the image as a
	// preview (probably resized as a thumbnail and placed
	// in a timeline).
	Image string `json:"image,omitempty"`

	// BannerImage is the URL of an image to use as a
	// banner. Some blogging systems display a different
	// banner image chosen to go with each post, but that
	// image wouldn’t otherwise appear in the ContentHTML. A
	// feed reader with a detail view may choose to show
	// this banner image at the top of the detail view,
	// possibly with the title overlaid.
	BannerImage string `json:"banner_image,omitempty"`

	// DatePublished specifies the date in RFC 3339 format.
	// (Example: 2010-02-07T14:04:00-05:00.)
	DatePublished time.Time `json:"date_published,omitempty"`

	// DateModified specifies the modification date in RFC
	// 3339 format.
	DateModified time.Time `json:"date_modified,omitempty"`

	// Author is the author of this item. If not specified
	// in an item, then the top-level author, if present, is
	// the author of the item.
	Author *Author `json:"author,omitempty"`

	// Tags can have any plain text values you want. Tags
	// tend to be just one word, but they may be anything.
	// Note: they are not the equivalent of Twitter
	// hashtags. Some blogging systems and other feed
	// formats call these categories.
	Tags []string `json:"tags,omitempty"`

	// Attachments lists related resources. Podcasts, for
	// instance, would include an attachment that’s an audio
	// or video file.
	Attachments []Attachment `json:"attachments,omitempty"`
}

// Attachment represents a JSON Feed attachment.
//
// Documentation for each field is taken from the
// JSON Feed spec.
type Attachment struct {
	// URL specifies the location of the attachment.
	URL string `json:"url"`

	// MIMEType specifies the type of the attachment, such
	// as “audio/mpeg”.
	MIMEType string `json:"mime_type"`

	// Title is a name for the attachment. Important: if
	// there are multiple attachments, and two or more have
	// the exact same title (when title is present), then
	// they are considered as alternate representations of
	// the same thing. In this way a podcaster, for
	// instance, might provide an audio recording in
	// different formats.
	Title string `json:"title,omitempty"`

	// SizeInBytes specifies how large the file is.
	SizeInBytes int `json:"size_in_bytes,omitempty"`

	// DurationInSeconds specifies how long it takes to
	// listen to or watch, when played at normal speed.
	DurationInSeconds int `json:"duration_in_seconds,omitempty"`
}

// Duration returns the duration stored in DurationInSeconds
// as a Duration.
func (a *Attachment) Duration() time.Duration {
	return time.Duration(a.DurationInSeconds) * time.Second
}
