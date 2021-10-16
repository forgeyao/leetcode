struct Solution;

impl Solution {
    pub fn count_and_say(n: i32) -> String {
        if n == 1 {
            return "1".to_string();
        }

        let s = Solution::count_and_say(n - 1);
        let mut c = s.chars().nth(0).unwrap();
        let mut count = 1;
        let mut ret = String::new();
        for i in 1..s.len() {
            if s.chars().nth(i).unwrap() == c {
                count += 1;
                continue;
            }
            ret += &(count.to_string() + &c.to_string());
            count = 1;
            c = s.chars().nth(i).unwrap();
        }
        ret + &(count.to_string() + &c.to_string())
    }

    pub fn count_and_say2(n: i32) -> String {
        let mut ret: String = "1".to_string();
        if n == 1 {
            return ret;
        }

        let count_str = |s: String| -> String {
            let mut c = s.chars().nth(0).unwrap();
            let mut count = 1;
            let mut ret = String::new();
            for i in 1..s.len() {
                if s.chars().nth(i).unwrap() == c {
                    count += 1;
                    continue;
                }
                ret += &(count.to_string() + &c.to_string());
                count = 1;
                c = s.chars().nth(i).unwrap();
            }
            ret + &(count.to_string() + &c.to_string())
        };

        for _ in 1..n {
            ret = count_str(ret);
        }
        ret
    }
}

fn main() {
    for i in 1..=7 {
        println!("ret: {:}, {:?}", i, Solution::count_and_say(i));
        println!("ret: {:}, {:?}", i, Solution::count_and_say2(i));
    }
}
