package jsonfeed

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	a := &Attachment{
		DurationInSeconds: 5,
	}
	got := a.Duration()
	want := 5 * time.Second
	if got != want {
		t.Errorf("(%v).Duration() = %v want %v", a, got, want)
	}
}
