
class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode() {
    }

    TreeNode(int val) {
        this.val = val;
    }

    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }

    public String toString() {
        StringBuilder sb = new StringBuilder();
        boolean ok = true;
        sb.append(this.val + " ");
        if (this.left != null) {
            ok = this.left.val < this.val;
            sb.append(this.left.val + " " + ok + " ");
        }
        if (this.right != null) {
            ok = this.right.val > this.val;
            sb.append(this.right.val + " " + ok + " ");
        }
        return sb.toString();
    }
}

class Solution {
    public boolean isValidBST(TreeNode root) {
        return isValid(root, Long.MIN_VALUE, Long.MAX_VALUE);
    }

    private boolean isValid(TreeNode node, long left, long right) {
        if (node == null) {
            return true;
        }
        System.out.println(node + " " + left + " " + right);
        if (node.val <= left || node.val >= right) {
            return false;
        }
        return isValid(node.left, left, node.val) && isValid(node.right, node.val, right);
    }
}
