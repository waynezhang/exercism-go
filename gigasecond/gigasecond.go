package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(1000000000))
}
