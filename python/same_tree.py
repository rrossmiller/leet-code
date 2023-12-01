# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isSameTree(self, p: TreeNode, q: TreeNode) -> bool:
        if p is None and q is None:
            return True
        if p is None or q is None or p.val != q.val:
            return False

        return self.isSameTree(q.left, p.left) and self.isSameTree(q.right, p.right)
