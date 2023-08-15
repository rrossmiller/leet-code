struct Solution;

impl Solution {
    pub fn backspace_compare(s: String, t: String) -> bool {
        let mut new_s = Self::process(s);
        let mut new_t = Self::process(t);

        new_s == new_t
    }

    fn process(s: String) -> String {
        let mut new_s = String::from("");

        for c in s.chars() {
            if c == '#' {
                new_s.pop();
            } else {
                new_s.push(c);
            }
        }
        new_s
    }
}

#[cfg(test)]
mod tests {
    use crate::backspace_compare::Solution;

    #[test]
    fn t1() {
        let s = String::from("ab#c");
        let t = String::from("ab#c");

        let r = Solution::backspace_compare(s, t);
        assert!(r);

        let s = String::from("b##c");
        let t = String::from("ab#c");

        let r = Solution::backspace_compare(s, t);
        assert!(!r);
    }
}
