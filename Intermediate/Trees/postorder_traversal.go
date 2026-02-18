package trees

/*
Problem Description
Given a binary tree, return the Postorder traversal of its nodes values.

Problem Constraints
1 <= number of nodes <= 105

Input Format
First and only argument is root node of the binary tree, A.

Output Format
Return an integer array denoting the Postorder traversal of the given binary tree.

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
 [3, 2, 1]
Output 2:
 [6, 3, 2, 1]

Example Explanation
Explanation 1:
 The Preoder Traversal of the given tree is [3, 2, 1].
Explanation 2:
 The Preoder Traversal of the given tree is [6, 3, 2, 1].
*/

func postorderTraversal(A *treeNode) []int {
	var res []int
	helperForPostOrder(A, &res)
	return res
}

func helperForPostOrder(A *treeNode, res *[]int) {
	if A == nil {
		return
	}

	helperForPostOrder(A.left, res)
	helperForPostOrder(A.right, res)
	*res = append(*res, A.value)
}

/*
The recursion visits each node once, but because each recursive call builds a new slice and appends
entire subtree results using variadic expansion, the worst-case time complexity becomes O(n²).
 In a balanced tree, it is approximately O(n log n). An optimized approach using a shared slice reduces
 it to O(n).(It will be amortized O(n))

  Sc is O(h) -> Recursive stack height, height of the tree
*/
