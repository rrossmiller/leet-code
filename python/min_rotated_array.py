class Solution:
    def findMin(self, nums: list[int]) -> int:
        # return sorted(nums)[0]
        # return min(nums)
        l, r = 0, len(nums) - 1

        res = nums[0]
        while l <= r:
            # if l-val < r-val, the array is sorted, rtn l-val
            if nums[l] < nums[r]:
                return min(res,nums[l])

            m = (l + r) // 2
            print(nums[l], nums[m], nums[r], f"({l} {m} {r})")
            res = min(res, nums[m])

            # if the middle is greater than the left,
            # search right
            if nums[m] >= nums[l]:
                l = m + 1

            # search left
            else:
                r = m - 1

        return res


if __name__ == "__main__":
    import os

    os.system("clear")
    s = Solution()
    # nums = [3, 1, 2]
    # assert s.findMin(nums) == 1, s.findMin(nums)

    # nums = [3, 4, 5, 1, 2]
    # assert s.findMin(nums) == 1

    nums = [4, 5, 1, 2, 3]
    assert s.findMin(nums) == 1
    exit()

    nums = [4, 5, 6, 7, 0, 1, 2]
    assert s.findMin(nums) == 0

    nums = [11, 13, 15, 17]
    assert s.findMin(nums) == 11
