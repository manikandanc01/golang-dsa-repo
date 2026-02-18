package trees

/*
Problem Description
Given a binary tree, return the inorder traversal of its nodes' values.

Problem Constraints
1 <= number of nodes <= 105

Input Format
First and only argument is root node of the binary tree, A.

Output Format
Return an integer array denoting the inorder traversal of the given binary tree.

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
 [1, 3, 2]

Output 2:
 [6, 1, 3, 2]


Example Explanation
Explanation 1:
 The Inorder Traversal of the given tree is [1, 3, 2].
Explanation 2:
 The Inorder Traversal of the given tree is [6, 1, 3, 2].
*/

func inorderTraversal(A *treeNode) []int {
	if A == nil {
		return []int{}
	}

	arr := inorderTraversal(A.left)
	arr = append(arr, A.value)
	arr = append(arr, inorderTraversal(A.right)...)

	return arr
}

// Tc is O(n), n -> size of the tree. (No.of nodes in the tree)
// Space complexity is O(n + h), height of the recursive stack if we are not considering the output spce -> O(1)

//---------------------x-----------------------x-------------------x---------

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func treeNode_new(value int) *treeNode {
	var node *treeNode = new(treeNode)
	node.value = value
	node.left = nil
	node.right = nil
	return node
}

// --------------------x------------------------------------x---------------------x--------
