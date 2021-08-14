package main

// Maximum Product of Splitted Binary Tree
//
// Medium
//
// https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/
//
// Given the `root` of a binary tree, split the binary tree into two subtrees by
// removing one edge such that the product of the sums of the subtrees is
// maximized.
//
// Return _the maximum product of the sums of the two subtrees_. Since the
// answer may be too large, return it **modulo** `109 + 7`.
//
// **Note** that you need to maximize the answer before taking the mod and not
// after taking it.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2020/01/21/sample_1_1699.png)
//
// ```
// Input: root = [1,2,3,4,5,6]
// Output: 110
// Explanation: Remove the red edge and get 2 binary trees with sum 11 and 10.
// Their product is 110 (11*10)
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2020/01/21/sample_2_1699.png)
//
// ```
// Input: root = [1,null,2,3,4,null,null,5,6]
// Output: 90
// Explanation: Remove the red edge and get 2 binary trees with sum 15 and
// 6.Their product is 90 (15*6)
//
// ```
//
// **Example 3:**
//
// ```
// Input: root = [2,3,9,10,7,8,6,5,4,11,1]
// Output: 1025
//
// ```
//
// **Example 4:**
//
// ```
// Input: root = [1,1]
// Output: 1
//
// ```
//
// **Constraints:**
//
// - The number of nodes in the tree is in the range `[2, 5 * 104]`.
// - `1 <= Node.val <= 104`
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxProduct(root *TreeNode) int {
	var walk func(n *TreeNode) int

	const mod = 1e9 + 7
	best := 0

	total := 0
	walk = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		cur := n.Val + walk(n.Left) + walk(n.Right)
		if total != 0 {
			v := (total - cur) * cur
			if v > best {
				best = v
			}
		}
		return cur
	}

	total = walk(root)
	walk(root)

	return best % mod
}
