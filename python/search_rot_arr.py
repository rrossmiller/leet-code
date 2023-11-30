class Solution:
    def search(self, nums: list[int], tgt: int) -> int:
        l, r = 0, len(nums) - 1

        while l <= r:
            m = (l + r) // 2
            if nums[m] == tgt:
                return m
            elif nums[l] == tgt:
                return l
            elif nums[r] == tgt:
                return r

            # if l-val < middle
            if nums[l] < nums[m]:
                # tgt < l-val, search right
                if tgt < nums[l] or tgt > nums[m]:
                    l = m + 1
                # tgt < middle, search left
                # elif tgt < nums[m]:
                else:
                    r = m - 1

            # l-val > middle
            else:
                #  tgt < middle, search right
                if tgt < nums[l] and tgt > nums[m]:
                    l = m + 1
                # search left
                else:
                    r = m - 1

            # if l-val < middle and tgt > middle,

        return -1


if __name__ == "__main__":
    import os

    os.system("clear")
    s = Solution()

    t = 1
    nums = [5, 1, 2, 3, 4]
    ans = s.search(nums, t)
    # print(nums)
    assert ans == 1, ans

    t = 6
    nums = [8, 1, 2, 3, 4, 5, 6, 7]
    ans = s.search(nums, t)
    # print(nums)
    assert ans == 6, ans

    nums = [1, 3]
    t = 3
    # print(nums)
    ans = s.search(nums, t)
    assert ans == 1, ans

    t = 4
    nums = [1, 2, 3, 4, 5, 6]
    # print(nums)
    ans = s.search(nums, t)
    assert ans == 3, ans

    nums = [4, 5, 6, 7, 0, 1, 2]
    t = 0
    # print(nums)
    ans = s.search(nums, t)
    assert ans == 4, ans

    nums = [4, 5, 6, 7, 0, 1, 2]
    t = 3
    assert s.search(nums, t) == -1
