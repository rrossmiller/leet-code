struct Solution {}

impl Solution {
    pub fn sorted_squares(nums: Vec<i32>) -> Vec<i32> {
        let mut x: Vec<i32> = nums.iter().map(|x| x.pow(2)).collect();
        x.sort();
        x
    }
}

#[cfg(test)]
mod tests {
    use std::vec;

    use super::Solution;

    #[test]
    fn t1() {
        let nums = vec![-4, -1, 0, 3, 10];
        let out = Solution::sorted_squares(nums);
        let soln = vec![0, 1, 9, 16, 100];
        assert_eq!(soln, out);

        let nums = vec![-7, -3, 2, 3, 11];
        let out = Solution::sorted_squares(nums);
        let soln = vec![4, 9, 9, 49, 121];
        assert_eq!(soln, out);
    }
}
