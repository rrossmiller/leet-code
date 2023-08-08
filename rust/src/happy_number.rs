struct Solution {}

impl Solution {
    pub fn is_happy(n: i32) -> bool {
        let mut n = Solution::get_sum(n);
        while n != 1 {
            n = Solution::get_sum(n);
            if n < 10 && n != 1{
                return false;
            }
        }
        true
    }

    fn get_sum(n: i32) -> i32 {
        let n: u32 = n
            .to_string()
            .chars()
            .map(|d| d.to_digit(10).unwrap())
            .map(|d| d.pow(2))
            .sum();
        n as i32
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn t1() {
        let n = Solution::get_sum(55);
        assert_eq!(50, n);
    }

    #[test]
    fn t2() {
        let happy = Solution::is_happy(19);

        assert!(happy);
    }
}
