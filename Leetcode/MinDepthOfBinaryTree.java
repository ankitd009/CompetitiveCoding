// https://leetcode.com/problems/minimum-depth-of-binary-tree/

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public int minDepth(TreeNode root) {
        if(root == null){
            return 0;
        }else{
            int left = minDepth(root.left);
            int right = minDepth(root.right);
            
            int depth = left;
            
            if (left > 0 && right > 0 ){
                depth = Math.min(left, right);
            }else if (left > 0 && right == 0 ){
                depth = left;
            } else if (left == 0 && right > 0 ){
                depth = right;
            }
            
            return 1 + depth;
        }
    }
}
