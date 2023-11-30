struct Solution;

impl Solution {
    fn fizz_buzz(n: i32) -> Vec<String> {
        let mut ans = Vec::new();
        for i in 1..=n {
            let mut s = String::new();
            if i % 3 == 0 {
                s.push_str("Fizz")
            }
            if i % 5 == 0 {
                s.push_str("Buzz")
            }
            if s.is_empty() {
                s.push_str(i.to_string().as_str())
            }

            ans.push(s);
        }
        ans
    }
}
#[cfg(test)]
mod tests {
    use crate::fizz_buzz::Solution;

    #[test]
    fn t1() {
        let ans = vec![
            "1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13",
            "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz", "Fizz", "22", "23", "Fizz",
            "Buzz", "26", "Fizz", "28", "29", "FizzBuzz", "31", "32", "Fizz", "34", "Buzz",
            "Fizz", "37", "38", "Fizz", "Buzz", "41", "Fizz", "43", "44", "FizzBuzz", "46", "47",
            "Fizz", "49", "Buzz",
        ];
        let x = Solution::fizz_buzz(50);
        for (i, v) in x.iter().enumerate() {
            assert_eq!(ans[i], v.as_str())
        }
    }
}
