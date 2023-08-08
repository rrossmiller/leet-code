//Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 < numbers.length.

// Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
// The tests are generated such that there is exactly one solution. You may not use the same element twice.

// Your solution must use only constant extra space.
// use std::cmp::Ordering::{Equal, Greater, Less};
// fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
//     let (mut l, mut r) = (0, numbers.len() - 1);
//
//     while l < r {
//         match (numbers[l] + numbers[r]).cmp(&target) {
//             Less => l += 1,
//             Greater => r -= 1,
//             Equal => return vec![l as i32 + 1, r as i32 + 1],
//         }
//     }
//
//     return vec![0, 0];
// }

fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
    let (mut l, mut r) = (0, numbers.len() - 1);

    while l < r {
        let res = numbers[l] + numbers[r];
        if res < target {
            l += 1;
        } else if res > target {
            r -= 1;
        } else if res == target {
            return vec![l as i32 + 1, r as i32 + 1];
        }
    }

    return vec![0, 0];
}
#[cfg(test)]
mod tests {
    use crate::twosum_2::two_sum;

    #[test]
    fn t1() {
        let numbers = vec![2, 7, 11, 15];
        let target = 9;
        let expected = vec![1, 2];
        let result = two_sum(numbers, target);
        println!("t1 {:?} == {:?}", result, expected);
        assert_eq!(result, expected);
    }

    #[test]
    fn t2() {
        let numbers = vec![2, 3, 4];
        let target = 6;
        let expected = vec![1, 3];
        let result = two_sum(numbers, target);
        println!("t2 {:?} == {:?}", result, expected);
        assert_eq!(result, expected);
    }
}
