from typing import Optional


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def dfs(self, root: TreeNode | None, res: list[int]) -> int:
        if root is None:
            return -1  # height of null tree

        h_left = self.dfs(root.left, res)
        h_right = self.dfs(root.right, res)

        d = 2 + h_left + h_right
        res[0] = max(res[0], d)

        return 1 + max(h_left, h_right)

    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        l = [0]
        self.dfs(root, l)
        return l[0]
