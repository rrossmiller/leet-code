struct Solution {}

impl Solution {
    pub fn longest_palindrome(s: String) -> i32 {
        let mut palindromes = Vec::new();

        for i in 0..s.chars().count() {
            for j in i + 1..s.chars().count() + 1 {
                let x = &s[i..j];
                if Solution::is_palindrome(x) {
                    palindromes.push(x);
                    println!("{} {}", x, x.chars().count());
                }
            }
        }

        if let Some(n) = palindromes.iter().map(|s| s.chars().count()).max() {
            return n as i32;
        } else {
            return 0;
        };
    }
    pub fn is_palindrome(s: &str) -> bool {
        let chars: Vec<char> = s.chars().collect();
        let len = chars.len();
        for i in 0..len / 2 {
            let a = chars[i];
            let b = chars[len - 1 - i];
            if a != b {
                return false;
            }
        }
        true
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn t1() {
        let s = "abc";
        let p = Solution::is_palindrome(s);
        assert!(!p);
    }

    #[test]
    fn t2() {
        let s = "aba";
        let p = Solution::is_palindrome(s);
        assert!(p);

        let s = "racecar";
        let p = Solution::is_palindrome(s);
        assert!(p)
    }

    #[test]
    fn t3() {
        let s = String::from("racecar");
        let p = Solution::longest_palindrome(s.clone());
        println!("{}", p);
        assert_eq!(p as usize, s.chars().count());

        let s = String::from("a");
        let p = Solution::longest_palindrome(s.clone());
        println!("{}", p);
        assert_eq!(p, 4);

        // let s = String::from("abccccdd");
        // let p = Solution::longest_palindrome(s.clone());
        // println!("{}", p);
        // assert_eq!(p, 7);
    }
}
