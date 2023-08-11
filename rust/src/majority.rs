struct Solution {}

use std::collections::HashMap;
impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        let half = nums.len() / 2;
        let mut elems = HashMap::new();
        for n in nums.iter() {
            let mut count: i32 = 1;
            if let Some(x) = elems.get(n) {
                count = *x+1;
            }

            if count as usize > half {
                return *n;
            }

            elems.insert(n, count);
        }

        return -1;
    }
}
