// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

const giga = 1e9 * time.Second

// AddGigasecond adds one gigaseconds to its input
func AddGigasecond(t time.Time) time.Time {
	return t.Add(giga)
}
