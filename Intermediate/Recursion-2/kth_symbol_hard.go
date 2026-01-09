package recursion2

/**
 * @input A : Integer
 * @input B : Long
 *
 * @Output Integer
 */
func solveKthSymbolHard(A int, B int64) int {
	if A == 1 {
		return 0
	}

	parent := solveKthSymbolHard(A-1, B/2)
	// if parent == 0{
	//     if B%2 == 0{
	//         return 0
	//     }
	//     return 1
	// }else if parent == 1{
	//     if B%2 == 0{
	//         return 1
	//     }
	//     return 0
	// }

	if B%2 == 0 {
		return parent
	}

	return 1 - parent
}

// Time Complexity: O(A) -> No.of recursive calls is A
// Space Complexity: O(A) -> Each recursive calls will take the stack space.
