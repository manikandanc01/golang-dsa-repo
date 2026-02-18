package trees

/*
Problem Description
Given a binary tree, return the preorder traversal of its nodes values.

Problem Constraints
1 <= number of nodes <= 105

Input Format
First and only argument is root node of the binary tree, A.

Output Format
Return an integer array denoting the preorder traversal of the given binary tree.

Example Input
Input 1:
   1
    \
     2
    /
   3

Input 2:
   1
  / \
 6   2
    /
   3

Example Output
Output 1:
 [1, 2, 3]
Output 2:
 [1, 6, 2, 3]

Example Explanation
Explanation 1:
 The Preoder Traversal of the given tree is [1, 2, 3].
Explanation 2:
 The Preoder Traversal of the given tree is [1, 6, 2, 3].
*/

func preorderTraversal(A *treeNode) []int {
	// if A == nil{
	//     return []int{}
	// }

	//  arr := []int{A.value}
	//  arr = append(arr,preorderTraversal(A.left)...)
	//  arr = append(arr,preorderTraversal(A.right)...)

	//  return arr

	var res []int
	helper(A, &res)
	return res
}

func helper(A *treeNode, res *[]int) {
	if A == nil {
		return
	}

	*res = append(*res, A.value)
	helper(A.left, res)
	helper(A.right, res)
}

/*
The recursion visits each node once, but because each recursive call builds a new slice and appends
entire subtree results using variadic expansion, the worst-case time complexity becomes O(n²).
 In a balanced tree, it is approximately O(n log n). An optimized approach using a shared slice reduces
 it to O(n).(It will be amortized O(n))

  Sc is O(h) -> Recursive stack height, height of the tree
*/
