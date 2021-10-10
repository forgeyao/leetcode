// https://leetcode-cn.com/problems/arranging-coins/
/**
 *
 * 本质是求和公式,已知 n, 求最大的 m
 * 需要主要是的越界问题
 *
 * 求和公式: m*(m+1)/2 <= n
 * 一元二次方程，可以解出 m <= (sqrt(1+8*n) -1 )/2
 */

struct Solution {}

impl Solution {
    pub fn arrange_coins(n: i32) -> i32 {
        // 方法一
        /*let mut m = n;
        let mut i = 0;
        while m >= 0 {
            i += 1;
            m -= i;
        }
        i - 1
        */

        /* 方法二
         * 提升类型避开越界 但超时了
        let mut i: i64 = (n / 2 + 1).into();
        let nn: i64 = 2 * (n as i64);
        let mut ii: i64 = i * (i + 1);
        while ii > nn {
            i -= 1;
            ii = i * (i + 1);
        }
        i as i32
        */

        // 方法三 数学公式
        // 测试提交会报错 expected 'n' to have value from 1 to 2147483647 only
        // 实际提交可以通过
        //((((1 + 8 * (n as i64)) as f64).sqrt() as i64 - 1) / 2) as i32

        // 方法四 没看懂怎么推导出来的
        let nn: i64 = (n as i64) * 2;
        let root: i64 = (nn as f32).sqrt() as i64;
        match root * (root + 1) > nn {
            true => (root as i32) - 1,
            false => (root as i32),
        }
    }
}

fn main() {
    let n: [i32; 4] = [1, 5, 8, 1804289383];
    for i in n {
        println!("{}", Solution::arrange_coins(i));
    }
}
