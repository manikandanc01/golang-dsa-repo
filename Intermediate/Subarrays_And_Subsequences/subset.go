package subarraysandsubsequences

/**
 * @input A : Integer array
 *
 * @Output 2D int array.
 */

/*
Problem Description
Given a set of distinct integers A, return all possible subsets.

NOTE:
Elements in a subset must be in non-descending order.
The solution set must not contain duplicate subsets.
Also, the subsets should be sorted in ascending ( lexicographic ) order.
The initial list is not necessarily sorted.

Problem Constraints
1 <= |A| <= 16
INTMIN <= A[i] <= INTMAX


Input Format
First and only argument of input contains a single integer array A.

Output Format
Return a vector of vectors denoting the answer.

Example Input
Input 1:
A = [1]

Input 2:
A = [1, 2, 3]

Example Output
Output 1:
[
    []
    [1]
]

Output 2:
[
 []
 [1]
 [1, 2]
 [1, 2, 3]
 [1, 3]
 [2]
 [2, 3]
 [3]
]


Example Explanation
Explanation 1:
 You can see that these are all possible subsets.

Explanation 2:
You can see that these are all possible subsets.
*/

import "sort"

func buildSubsets(idx, n int, A []int, solutionArray *[][]int, currArray []int) {
	if idx == n {
		snapshot := make([]int, len(currArray))
		copy(snapshot, currArray)
		*solutionArray = append(*solutionArray, snapshot)
		return
	}

	currElement := A[idx]

	// Select the current element and go on further wards
	currArray = append(currArray, currElement)
	buildSubsets(idx+1, n, A, solutionArray, currArray)

	// not selecting current element
	currArray = currArray[:len(currArray)-1]
	buildSubsets(idx+1, n, A, solutionArray, currArray)

}

func subsets(A []int) [][]int {
	sort.Ints(A)

	subsets := make([][]int, 0)
	currArray := []int{}

	buildSubsets(0, len(A), A, &subsets, currArray)

	sort.Slice(subsets, func(i, j int) bool {
		a, b := subsets[i], subsets[j]

		minLen := len(a)
		if len(b) < len(a) {
			minLen = len(b)
		}

		for k := 0; k < minLen; k++ {
			if a[k] < b[k] {
				return true
			} else if a[k] > b[k] {
				return false
			}
		}

		// If all elements are equal, return the min length array
		return len(a) < len(b)
	})
	return subsets
}

// This is the brutte force backtracking algortithm
// Every element have 2 choices. So there are 2^n recursive calls.
// The copying of element happens O(n) at the base case. Total 2^n times it will happen.
// Total subsets: 2^n
// Sorting of 2^n element with O(n) comparison. n.2^nlog 2^n, (n^2. 2^n)
// Time complexity is: O(n.2^n + n^2.2^n)

// In this way we can't maintain the order. We should sort the final result to bring the lexicographical order. 


/*

METHOD 2: To maintain the natural order without sorting

import "sort"

var a []int  
var ans [][]int

func buildSubsets(temp []int, index int) {
    snapshot := make([]int, len(temp))
    copy(snapshot,temp)
    ans = append(ans, snapshot)

    // ans = append(ans, temp)
    for i:=index; i<len(a); i++{
        temp = append(temp, a[i])
        buildSubsets(temp, i+1)
        temp = temp[:len(temp)-1]

        // var cur []int
        // for j:=0; j<len(temp)-1; j++{
        //     cur = append(cur,temp[j])
        // }
        // temp = cur

       
        //   temp = temp[:len(temp)-1]  

        //   in the above line we are just reducing the slot lenght. 
        //   the underlying array will be same. 
        //   next append will overwrite the existing element in the underlying array. 
        //   this will modify the already stored element. Because of shared memory.
        //   so creating new slice and this will point to different underlying array. 
        //   before appending new element and inserting into ans. 

        //   This is also not efficient. Because os manual copying everytime. 
        //   Best way is take the snapshot before adding to the solution. 

        //   And keeping the gloabal variables are also not good practice. 

    }
}

func subsets(A []int )  ([][]int) {
    sort.Ints(A)
    ans = nil // Without this submission fails. Because ans is a global variable. For each test case this should be reset to nil. Otherwisesome weird output will come. 
    a = A  
    var temp []int
    buildSubsets(temp,0)
    return ans
}

// This method helps to maintain the lexicographical order. But before that entire slice should be sorted. 
// Time complexity: O(n.2^n) , for copying to snapshot at worst case takes O(n) Time
// Space complexity: To store 2^n subsets. O(2^n) and recursive stack space is O(n)
// And also we needs to copy the slice to snapshot. This also takes O(n) space.
// SC is O(n.2^n + n) 

*/

/*

METHOD 3: We can use the bitwise logic. 

Input: [1,2,3]

Algorithm:
1. Sort the input array
2. Iterate 2^n times outer loop(i), inner loop(j)  n times if i&(1<<j) we can include the A[j]th element in the result.
3. Once this is done, sort the result again to lexicographical order. 

[0,0,0], [0,0,1], [0,1,0],[0,1,1],[1,0,0], [1,0,1], [1,1,1]

Just populate the elements where 1 is there and skip 

sample code:
   sort(A)
   for(int i=0;i<end;i++)
    {
        vector<int>temp;
        for(int j=0;j<n;j++)
        {
            if((i&(1<<j))!=0)
            {
                temp.push_back(A[j]);
            }
        }
        res.push_back(temp);
    }
	sort(res)

Time complexity: O(n log n + 2^n.n + n^2.2^n) (n^2.2^n) -> for final sorting as we see in the first method. 
*/
