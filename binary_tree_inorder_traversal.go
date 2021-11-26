/**
94. Binary Tree Inorder Traversal

Given the root of a binary tree, return the inorder traversal of its nodes' values. 

Example 1:
Input: root = [1,null,2,3]
Output: [1,3,2]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [1]
Output: [1]

Example 4:
Input: root = [1,2]
Output: [2,1]

Example 5:
Input: root = [1,null,2]
Output: [1,2]
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// recursive implementation
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    
    left := inorderTraversal(root.Left)
    middle := append(left, root.Val)
    right := inorderTraversal(root.Right)
    return append(middle, right...)
}

// iteratively implementation
func inorderTraversal(root *TreeNode) []int {
    var result []int
    var stack []*TreeNode
    
    for root != nil || len(stack) > 0 {
        if root != nil {
            stack = append(stack, root)
            root = root.Left
            continue
        }
        
        stackLen := len(stack)
        root = stack[stackLen-1]
        stack = stack[0:stackLen-1]
        result = append(result, root.Val)
        root = root.Right
    }
    return result
}
