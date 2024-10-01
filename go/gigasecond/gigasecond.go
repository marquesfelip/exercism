// Package gigasecond provides a function to add one gigasecond (1,000,000,000 seconds) to a given time.Time value.
package gigasecond

import "time"

// AddGigasecond adds 1 gigasecond (1,000,000,000 seconds) to the given time.Time value
// and returns the resulting time.Time.
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1e9)

	return t
}
