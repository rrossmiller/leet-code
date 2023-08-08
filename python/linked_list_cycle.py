# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        prev = set()
        while head is not None:
            prev.add(head)
            head = head.next
            if head in prev:
                return True

        return False
