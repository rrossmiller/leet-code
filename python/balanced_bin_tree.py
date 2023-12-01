class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def height(self, root: TreeNode | None) -> int:
        if root is None:
            return 0  # height of null tree
        l = self.height(root.left)
        r = self.height(root.right)

        # short circuit
        if l == -1 or r == -1:
            return -1

        # if their heights differ by more than one, return -1
        if abs(l - r) > 1:
            return -1

        # Otherwise, return the height of this subtree
        return max(l, r) + 1

    def isBalanced(self, root: TreeNode | None) -> bool:
        if root is None:
            return True
        return self.height(root) != -1
