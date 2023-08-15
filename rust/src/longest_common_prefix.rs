struct Solution;

impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        String::from("")
    }
}

#[cfg(test)]
mod tests {
    use crate::longest_common_prefix::Solution;

    #[test]
    fn t1() {
        let i = vec![
            "flower".to_string(),
            "flow".to_string(),
            "flight".to_string(),
        ];

        let r = Solution::longest_common_prefix(i);
        assert_eq!(r,"fl");
    }
}
