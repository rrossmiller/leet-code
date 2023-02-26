from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        try:
            return nums.index(target) # lol
        except:
            return -1

if __name__ == "__main__":
    nums = [-1, 0, 3, 5, 9, 12]
    target = 9
    print(Solution().search(nums, target))
    assert Solution().search(nums, target) == 4

    nums = [-1, 0, 3, 5, 9, 12]
    target = 2
    print(Solution().search(nums, target))
