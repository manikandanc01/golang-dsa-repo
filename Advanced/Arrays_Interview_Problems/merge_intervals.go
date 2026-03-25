package arraysinterviewproblems

/*
Problem Description
You have a set of non-overlapping intervals. You are given a new interval [start, end], insert this new interval into the set of intervals (merge if necessary).
You may assume that the intervals were initially sorted according to their start times.

Problem Constraints
0 <= |intervals| <= 105

Input Format
First argument is the vector of intervals
second argument is the new interval to be merged

Output Format
Return the vector of intervals after merging

Example Input
Input 1:
Given intervals [1, 3], [6, 9] insert and merge [2, 5] .

Input 2:
Given intervals [1, 3], [6, 9] insert and merge [2, 6] .

Example Output
Output 1:
 [ [1, 5], [6, 9] ]

Output 2:
 [ [1, 9] ]

Example Explanation
Explanation 1:
(2,5) does not completely merge the given intervals

Explanation 2:
(2,6) completely merges the given intervals
*/

/**
* Definition of Structure
* type Interval struct{
*   start, end int
* }
 */

type Interval struct {
	start, end int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	// First check the 2 intervals are overlapping or not. 
	// If it is overlapping, merge the intervals, (1,3) and (2,5) => merged to be (1,5)
	// If it is non-overlapping, now append the minimum start value interval to the result intervals. 
	// Tc is O(n), Sc is O(1)
	resultIntervals := make([]Interval, 0)
	// isMerged := false

	for i := 0; i < len(intervals); i++ {
		if isOverlapping(intervals[i], newInterval) {
			newStart := intervals[i].start
			if newInterval.start < newStart {
				newStart = newInterval.start
			}

			newEnd := intervals[i].end
			if newInterval.end > newEnd {
				newEnd = newInterval.end
			}

			newInterval.start = newStart
			newInterval.end = newEnd

			// isMerged = true
		} else {
			minStartInterval := Interval{start: newInterval.start, end: newInterval.end}
			nextInterval := Interval{start: intervals[i].start, end: intervals[i].end}
			if intervals[i].start < newInterval.start {
				minStartInterval.start = intervals[i].start
				minStartInterval.end = intervals[i].end

				nextInterval.start = newInterval.start
				nextInterval.end = newInterval.end
			}

			resultIntervals = append(resultIntervals, minStartInterval)
			newInterval.start = nextInterval.start
			newInterval.end = nextInterval.end
		}
	}

	resultIntervals = append(resultIntervals, newInterval)

	return resultIntervals
}

func isOverlapping(i1 Interval, i2 Interval) bool {
	// Ensure i1 start is always lesser than i2
	if i1.start > i2.start {
		temp := i2
		i2 = i1
		i1 = temp
	}

	if i1.end < i2.start {
		return false
	}

	return true
}
