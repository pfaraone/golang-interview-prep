// learned to use multidimensional slices and sort package

package main

import (
	"sort"
)

func isOverlapping(periodA, periodB []int) bool {
	return periodA[0] <= periodB[1] && periodA[1] >= periodB[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	//     handle edge case of interval
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)

	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		lastSeenInterval := res[len(res)-1:][0]
		if isOverlapping(lastSeenInterval, intervals[i]) {
			lastSeenInterval[0] = min(lastSeenInterval[0], intervals[i][0])
			lastSeenInterval[1] = max(lastSeenInterval[1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

func main() {

}
