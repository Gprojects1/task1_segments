package main

import (
	"fmt"
	"sort"
)

type Segment struct {
	Start int
	End   int
}

func countPoints(segments []Segment) int {
	sort.Slice(segments, func(i, j int) bool {
		return segments[i].End < segments[j].End
	})

	count := 0
	covered := -1

	for _, segment := range segments {
		if segment.Start > covered {
			count++
			covered = segment.End
		}
	}

	return count
}

func main() {
	var n int
	fmt.Scan(&n)

	segments := make([]Segment, n)
	for i := 0; i < n; i++ {
		var start, end int
		fmt.Scan(&start, &end)
		segments[i] = Segment{start, end}
	}

	fmt.Println(countPoints(segments))
}
