class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isSameTree(self, p: TreeNode | None, q: TreeNode | None) -> bool:
        if p is None and q is None:
            return True
        if p is None or q is None or p.val != q.val:
            return False

        return self.isSameTree(q.left, p.left) and self.isSameTree(q.right, p.right)

    def isSubtree(self, root: TreeNode | None, subRoot: TreeNode | None) -> bool:
        """
        Use dfs and same_tree
        if the roots are the same, check if the tree is the same, otherwise continue
        """
        # TreeNode.__str__ = lambda x: str(x.val)
        if root is None and subRoot is None:
            return True
        if root is None or subRoot is None:
            return False

        if root.val == subRoot.val and self.isSameTree(root, subRoot):
            return True

        return self.isSubtree(root.left, subRoot) or self.isSubtree(root.right, subRoot)
