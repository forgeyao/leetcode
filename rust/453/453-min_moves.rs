// https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/

impl Solution {
    pub fn min_moves(nums: Vec<i32>) -> i32 {
        let mut min = nums[0];
        for num in &nums {
            if min > *num {
                min = *num;
            }
        }
        let mut count = 0;
        for num in &nums {
            count += *num - min;
        }
        return count;
    }
}

pub struct Solution {}

fn main() {
    let num = vec![1, 2, 3];
    println!("ret:{}", Solution::min_moves(num));
}
