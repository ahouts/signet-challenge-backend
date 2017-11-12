package main

import "time"

func timeRangeIntersect(r1Start, r1End, r2Start, r2End time.Time) bool {
	return (r1Start.Before(r2End) && r1Start.After(r2Start)) ||
		(r1End.Before(r2End) && r1End.After(r2Start)) ||
		(r2Start.Before(r1End) && r2Start.After(r1Start)) ||
		(r2End.Before(r1End) && r2End.After(r1Start))
}
