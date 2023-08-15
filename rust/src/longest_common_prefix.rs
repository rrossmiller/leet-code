struct Solution;

impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        if strs[0] == "" {
            return String::from("");
        } else if strs.len() == 1 {
            return strs[0].clone();
        }
        let mut i = 0;
        let mut is_common = true;

        // while the starting letters of the other words are same as the first
        while is_common {
            for j in 1..strs.len() {
                if strs[j].chars().count() == i + 1
                    || strs[0].chars().count() == i + 1
                    || &strs[0][0..i + 1] != &strs[j][0..i + 1]
                {
                    is_common = false;
                    break;
                }
            }
            i += 1;
        }

        String::from(&strs[0][0..i - 1])
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
        assert_eq!(r, "fl");
    }
}
