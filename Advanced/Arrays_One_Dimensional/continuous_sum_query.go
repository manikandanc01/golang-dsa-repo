package arraysonedimensional

/*
Problem Description
There are A beggars sitting in a row outside a temple. Each beggar initially has an empty pot.
When the devotees come to the temple, they donate some amount of coins to these beggars.
Each devotee gives a fixed amount of coin(according to their faith and ability) to some K beggars sitting
next to each other.

Given the amount P donated by each devotee to the beggars ranging from L to R index,
where 1 <= L <= R <= A, find out the final amount of money in each beggar's pot at the end of the day,
provided they don't fill their pots by any other means.

For ith devotee B[i][0] = L, B[i][1] = R, B[i][2] = P, given by the 2D array B

Problem Constraints
1 <= A <= 2 * 105
1 <= L <= R <= A
1 <= P <= 103
0 <= len(B) <= 105

Input Format
The first argument is a single integer A.
The second argument is a 2D integer array B.

Output Format
Return an array(0 based indexing) that stores the total number of coins in each beggars pot.

Example Input
Input 1:-
A = 5
B = [[1, 2, 10], [2, 3, 20], [2, 5, 25]]

Example Output
Output 1:-
10 55 45 25 25

Example Explanation
Explanation 1:-
First devotee donated 10 coins to beggars ranging from 1 to 2. Final amount in each beggars pot after first devotee: [10, 10, 0, 0, 0]
Second devotee donated 20 coins to beggars ranging from 2 to 3. Final amount in each beggars pot after second devotee: [10, 30, 20, 0, 0]
Third devotee donated 25 coins to beggars ranging from 2 to 5. Final amount in each beggars pot after third devotee: [10, 55, 45, 25, 25]
*/

/**
 * @input A : Integer
 * @input B : 2D integer array
 *
 * @Output Integer array.
 */
func solveContinuousSumQuery(A int, B [][]int) []int {
	// Brute Force: Iterate over B, for each  starting and ending index in B[i] iterate from B[i][0] to B[i][1] index
	// Sum to the answer array
	// Time complexity: O(m * n) , m is the size of array B, and n is the size of the array with length A



	// Optimal solution will be using Prefix sum technique. 
	// in prefix sum arr[0] is added to the result from the index 0 to n-1 
	// arr[1] element is added to all the elements from the index 1 to n-1
	// We need to add the values in the specifix range. 
	// initially the all of the beggarsPot is empty. Let's say [0,0,0,0,0]
	// For eg: B[i][0] = 1 (leftIdx), B[i][1] = 2 (rightIdx), B[i][2] = 10 (amount)
	// We should add the amount 10 to beggarsPot from the index 0 to 1 (Because B values are 1 base indexing)
	// what we can do is we add beggarsPot[0] = 10, and beggarsPot[2] as -10. 
	// Finally we will take the prefix sum of the beggarsPot. 
	// Prefix sum is 0 to ith index sum. 
	// While we are doing the first amount 10 will be added to all the beggarsPot. 
	// To correct this we subtracted the amount 10 from the next beggarsPot to the last beggar. 
	// Time Complexity: O(m+n) ,  m is the size of the B array, n is the number of beggars. 

	beggarsPot := make([]int, A)
	for _, innerArr := range B {
		leftIdx := innerArr[0] - 1
		rightIdx := innerArr[1]
		amount := innerArr[2]

		beggarsPot[leftIdx] += amount
		if rightIdx < A {
			beggarsPot[rightIdx] -= amount
		}
	}

	finalBeggarsPotAmounts := make([]int, A)
	for i, amount := range beggarsPot {
		if i == 0 {
			finalBeggarsPotAmounts[i] = amount
		} else {
			finalBeggarsPotAmounts[i] = finalBeggarsPotAmounts[i-1] + amount
		}
	}

	return finalBeggarsPotAmounts
}
