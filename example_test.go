package jsonfeed_test

import (
	"encoding/json"
	"fmt"

	"github.com/kr/jsonfeed"
)

func Example() {
	var f jsonfeed.Feed
	err := json.Unmarshal(microblog, &f)
	if err != nil {
		panic(err)
	}
	fmt.Println(f.Items[0].ID)
	// output:
	// 2347259
}

var microblog = []byte(`
{
    "version": "https://jsonfeed.org/version/1",
    "user_comment": "This is a microblog feed. You can add this to your feed reader using the following URL: https://example.org/feed.json",
    "title": "Brent Simmons’s Microblog",
    "home_page_url": "https://example.org/",
    "feed_url": "https://example.org/feed.json",
    "author": {
        "name": "Brent Simmons",
        "url": "http://example.org/",
        "avatar": "https://example.org/avatar.png"
    },
    "items": [
        {
            "id": "2347259",
            "url": "https://example.org/2347259",
            "content_text": "Cats are neat. \n\nhttps://example.org/cats",
            "date_published": "2016-02-09T14:22:00-07:00"
        }
    ]
}
`)
