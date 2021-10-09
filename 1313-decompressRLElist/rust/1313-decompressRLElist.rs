// https://leetcode-cn.com/problems/decompress-run-length-encoded-list/

struct Solution {}

impl Solution {
    pub fn decompress_rl_elist(nums: Vec<i32>) -> Vec<i32> {
        let mut v: Vec<i32> = Vec::new();
        //let mut i = 0;
        //while i < nums.len() - 1 {
        for i in (0..nums.len()).step_by(2) {
            for _ in 0..nums[i] {
                v.push(nums[i + 1]);
            }
            //i += 2;
        }
        v
    }
}

fn main() {
    let nums = vec![1, 1, 2, 3];
    println!("{:?}", Solution::decompress_rl_elist(nums));
}
