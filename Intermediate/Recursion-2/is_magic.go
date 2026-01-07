package recursion2

/*
Problem Description
Given a number A, check if it is a magic number or not.
A number is said to be a magic number if the sum of its digits is calculated till a single digit recursively by adding the sum of the digits after every addition. If the single digit comes out to be 1, then the number is a magic number.


Problem Constraints
1 <= A <= 109

Input Format
The first and only argument is an integer A.

Output Format
Return an 1 if the given number is magic else return 0.



Example Input
Input 1:
 A = 83557
Input 2:
 A = 1291


Example Output
Output 1:
 1
Output 2:
 0


Example Explanation
Explanation 1:
 Sum of digits of (83557) = 28
 Sum of digits of (28) = 10
 Sum of digits of (10) = 1.
 Single digit is 1, so it's a magic number. Return 1.
Explanation 2:
 Sum of digits of (1291) = 13
 Sum of digits of (13) = 4
 Single digit is not 1, so it's not a magic number. Return 0.
*/

/**
 * @input A : Integer
 *
 * @Output Integer
 */

// Tc : -> At max the sum will be within 3 digits -> 2 digits -> 1 digits. (9999999999), It is almost constant
// Tc: O(log(10) A)
// Sc: O(1) -> Depth is always constant. 



func solve(A int) int {
	sum := 0
	num := A
	for num > 0 {
		r := num % 10
		num = num / 10
		sum += r
	}

	// if sum >= 1 && sum < 10 {   // digits sum will be always positive
	if sum < 10 {
		if sum == 1 {
			return 1
		}
		return 0
		// }else{
		//     return 0
		// }
		// Instead of redundant else case we can just return 0.
	}

	return solve(sum)
}


// Optimized Solution
func solve2(A int )  (int) {
    if A%9 == 1{
        return 1
    }

    return 0
}

// divisibility rule of 9. Sum of all digits divide by 9, then the whole number will be divide by 9. 
// if the sum of all digits is one more than the number divisible by 9, which is the remainder 1, then we can say the number's further sum's also leave 1 at the end. 
// Tc is: O(1), Sc is O(1)


// Doubt arised whenever I am coding
// divisibility rule of 3 is also similar to 9. But that won't work here. 
// Because modulo by 3 wonn't give all the single digits. 
// (0,1,2) alone. using 3 won't work
// (0,1,2,3,4,5,....8) -> mod by 9, 10 (2 digit number gives value 1 when mod by 10). 