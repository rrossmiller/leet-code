struct Solution;
impl Solution {
    pub fn move_zeroes(nums: &mut Vec<i32>) {
        let mut snow_ball_size = 0;
        for i in 0..nums.len() {
            if nums[i] == 0 {
                snow_ball_size += 1;
            } else if snow_ball_size > 0 {
                let t = nums[i];
                nums[i] = 0;
                nums[i - snow_ball_size] = t;
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use crate::move_zeroes::Solution;

    #[test]
    fn t1() {
        let mut i = vec![0, 1, 0, 3, 12];
        println!("{:?}", i);
        Solution::move_zeroes(&mut i);
        let o = vec![1, 3, 12, 0, 0];

        println!("{:?}", i);
        println!("{:?}", o);
        assert_eq!(i, o);
    }
    #[test]
    fn t2() {
        let mut i = vec![0, 0, 1];
        println!("{:?}", i);

        Solution::move_zeroes(&mut i);
        let o = vec![1, 0, 0];

        println!("{:?}", i);
        println!("{:?}", o);
        assert_eq!(i, o);
    }
}
