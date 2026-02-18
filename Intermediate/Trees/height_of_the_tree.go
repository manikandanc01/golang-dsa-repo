package trees

/*
Problem Description
You are given the root node of a binary tree A. You have to find the height of the given tree.
A binary tree's height is the number of nodes along the longest path from the root node down to the farthest leaf node.

Problem Constraints
1 <= Number of nodes in the tree <= 105
0 <= Value of each node <= 109

Input Format
The first and only argument is a tree node A.

Output Format
Return an integer denoting the height of the tree.

Example Input
Input 1:
 Values =  1
          / \
         4   3
Input 2:
 Values =  1
          / \
         4   3
        /
       2

Example Output
Output 1:
 2
Output 2:
 3
*/

func solve(A *treeNode) int {
	if A == nil {
		return 0
	}

	leftSubtreeHeight := solve(A.left)
	rightSubTreeHeight := solve(A.right)

	// return max(lh,rh)+1   // go 1:21+ have this built in function max. But below that don't have.
	maxHeight := leftSubtreeHeight
	if rightSubTreeHeight > leftSubtreeHeight {
		maxHeight = rightSubTreeHeight
	}

	return maxHeight + 1
}

// Tc is O(n), Sc is O(log n)
