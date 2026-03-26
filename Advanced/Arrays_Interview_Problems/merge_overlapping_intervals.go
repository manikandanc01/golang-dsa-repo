package arraysinterviewproblems

/*
Problem Description
Given a collection of intervals, merge all overlapping intervals.

Problem Constraints
1 <= Total number of intervals <= 100000.

Input Format
First argument is a list of intervals.

Output Format
Return the sorted list of intervals after merging all the overlapping intervals.

Example Input
Input 1:
[1,3],[2,6],[8,10],[15,18]

Example Output
Output 1:
[1,6],[8,10],[15,18]

Example Explanation
Explanation 1:
Merge intervals [1,3] and [2,6] -> [1,6].
so, the required answer after merging is [1,6],[8,10],[15,18].
No more overlapping intervals present.
*/

import "sort"

/*
  type Interval struct {
    start, end int
  }
*/

func merge(A []Interval) []Interval {
	// Tc will be O(n log n), Sc will be O(n) if we consider the output space else O(1)
	sort.Slice(A, func(i, j int) bool {
		return A[i].start <= A[j].start
	})

	resultIntervals := make([]Interval, 0)

	prevInterval := A[0]
	for i := 1; i < len(A); i++ {
		if isOverlapping2(prevInterval, A[i]) {
			prevInterval.start = minStart(prevInterval.start, A[i].start)
			prevInterval.end = maxEnd(prevInterval.end, A[i].end)
		} else {
			resultIntervals = append(resultIntervals, prevInterval)
			prevInterval = A[i]
		}
	}

	resultIntervals = append(resultIntervals, prevInterval)
	return resultIntervals

}

func isOverlapping2(i1, i2 Interval) bool {
	if i1.end < i2.start {
		return false
	}

	return true
}

func minStart(i1, i2 int) int {
	if i2 < i1 {
		return i2
	}

	return i1
}

func maxEnd(i1, i2 int) int {
	if i2 > i1 {
		return i2
	}

	return i1
}
