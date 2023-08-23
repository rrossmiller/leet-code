class palindromeLinkedList {
    public static void main(String[] args) {

    }

}

class ListNode {
    int val;
    ListNode next;

    ListNode() {
    }

    ListNode(int val) {
        this.val = val;
    }

    ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }
}

class Solution {
    public boolean isPalindrome(ListNode head) {
        ListNode fast = head;
        ListNode slow = head;
        while (fast != null && fast.next != null) {
            fast = fast.next.next;
            slow = slow.next;
        }
        ListNode tail = revese(slow);

        while (tail != null) {
            if (head.val != tail.val) {
                return false;

            }

            head = head.next;
            tail = tail.next;
        }
        return true;
    }

    private ListNode revese(ListNode node) {

        ListNode prev = null;
        ListNode current = node;
        ListNode tmp = null;
        while (current != null) {
            tmp = current.next;
            current.next = prev;
            prev = current;
            current = tmp;
        }
        node = prev;

        return node;
    }
}
