package subarraysandsubsequences

/*
Problem Description
Little Ponny has been given a string A, and he wants to find out the lexicographically minimum subsequence from it of size >= 2. Can you help him?
A string a is lexicographically smaller than string b, if the first different letter in a and b is smaller in a. For example, "abc" is lexicographically smaller than "acc" because the first different letter is 'b' and 'c' which is smaller in "abc".

Problem Constraints
1 <= |A| <= 105
A only contains lowercase alphabets.

Input Format
The first and the only argument of input contains the string, A.

Output Format
Return a string representing the answer.

Example Input
Input 1:
 A = "abcdsfhjagj"
Input 2:
 A = "ksdjgha"

Example Output
Output 1:
 "aa"
Output 2:
 "da"

Example Explanation
Explanation 1:
 "aa" is the lexicographically minimum subsequence from A.
Explanation 2:
 "da" is the lexicographically minimum subsequence from A.
*/

/**
 * @input A : String
 *
 * @Output string.
 */

func solveLittlePonyProblem(A string) string {
	// As per the question the answer will always be of length 2
	// Last character can't be a first character so we can exclude that
	// In the rest of the string find the first occurence of the smallest character
	// and before that the strings are invalid we can exclude, we find the second occurence after that point
	// Tc is O(n + n), Sc is O(1)

	smallestCharacter := ""
	firstOccurencePosition := -1

	for i := 0; i < len(A)-1; i++ {
		if smallestCharacter == "" {
			// smallestCharacter = A[i] //  A[i] is byte. Cannot assign byte to a string
			smallestCharacter = string(A[i])
		} else if A[i] < smallestCharacter[0] {
			//  smallestCharacter[0] = A[i]   // Cannot assign to smallestCharacter[0] (string are immutable)
			smallestCharacter = string(A[i])
			firstOccurencePosition = i
		}
	}

	secondCharacter := ""
	for j := firstOccurencePosition + 1; j < len(A); j++ {
		if secondCharacter == "" {
			// secondCharacter = A[j] // A[j] is byte. Cannot assign byte to a string
			secondCharacter = string(A[j])
		} else if A[j] < secondCharacter[0] {
			//    secondCharacter[0] = A[j]
			secondCharacter = string(A[j])
		}
	}

	return smallestCharacter + secondCharacter

	/*
	   // Brute force approach
	   // TLE, Tc is O(n^2), Sc is O(1)
	   n := len(A)
	    var minimumLengthSubsequence string

	    for i:=0; i<n; i++{
	        for j:= i+1; j<n; j++{
	            currString := string(A[i]) + string(A[j])
	            if len(minimumLengthSubsequence) == 0{
	                minimumLengthSubsequence = currString
	            }else{
	                if compareTwoStrings(minimumLengthSubsequence, currString){
	                    minimumLengthSubsequence = currString
	                }
	            }
	        }
	    }

	    return minimumLengthSubsequence*/
}

func compareTwoStrings(old, curr string) bool {
	n := len(old)
	if len(curr) < len(old) {
		n = len(curr)
	}

	for i := 0; i < n; i++ {
		if curr[i] < old[i] {
			return true
		}
	}

	return false
}
