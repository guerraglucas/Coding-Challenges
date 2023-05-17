// Given the root of a binary search tree and the lowest and highest boundaries as low and high,
// trim the tree so that all its elements lies in [low, high]. Trimming the tree should not change
// the relative structure of the elements that will remain in the tree (i.e., any node's descendant
// should remain a descendant). It can be proven that there is a unique answer.

// Return the root of the trimmed binary search tree. Note that the root may change depending
// on the given bounds.

// Example 1:

// Input: root = [1,0,2], low = 1, high = 2
// Output: [1,null,2]

// Example 2:

// Input: root = [3,0,4,null,2,null,null,1], low = 1, high = 3
// Output: [3,2,null,1]

// Constraints:

// The number of nodes in the tree is in the range [1, 104].
// 0 <= Node.val <= 104
// The value of each node in the tree is unique.
// root is guaranteed to be a valid binary search tree.
// 0 <= low <= high <= 104
package mockAssessment

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TrimBST(root *TreeNode, low int, high int) *TreeNode {

	// if the tree node is nil, return
	if root == nil {
		return nil
	}
	// the tree is guaranteed to be a valid binary search tree
	// it means that the value of the node will be greater than the node.L and smaller than node.R e.g node.L.value < node.value < node.R.value
	// so, if the node.value is less than the 'low' value, we need to update this node to be node.R, to keep the node.L (if exists) less than itself
	// to keep the binary search tree valid
	if root.Val < low {
		return TrimBST(root.Right, low, high)
	}
	// if the node.value is greater than the 'high' value, means that we need to update this node to point to his node.L children
	// to keep the binary search tree valid
	if root.Val > high {
		return TrimBST(root.Left, low, high)
	}

	// if the node is valid, call the function recursively for both its nodes.
	root.Left = TrimBST(root.Left, low, high)
	root.Right = TrimBST(root.Right, low, high)

	// return the updated tree
	return root
}
