package recursion1

/*
Problem Description
Write a recursive function that takes a string, S, as input and prints the characters of S in reverse order.

Problem Constraints
1 <= |s| <= 1000

Input Format
First line of input contains a string S.

Output Format
Print the character of the string S in reverse order.

Example Input
Input 1:
 scaleracademy
Input 2:
 cool

Example Output
Output 1:
 ymedacarelacs
Output 2:
 looc
*/

import (
	"fmt"
)

// Tc is O(N), Sc is O(N) due to recursion stack
func printStringReverse(str string, idx int) {
	if idx == 0 {
		fmt.Printf("%c", str[idx])
		return
	}

	fmt.Printf("%c", str[idx])
	printStringReverse(str, idx-1)
}

func main() {
	// YOUR CODE GOES HERE
	// Please take input and print output to standard input/output (stdin/stdout)
	// E.g. 'fmt.Scanf' for input & 'fmt.Printf' for output

	var str string
	fmt.Scan(&str)

	printStringReverse(str, len(str)-1)

}
