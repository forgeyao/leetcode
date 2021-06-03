//https://leetcode-cn.com/problems/duplicate-emails/

select distinct(a.Email) from Person as a, Person as b where a.Id != b.Id and a.Email = b.Email;

// 其他参考方法
select Email from Person group by Email having count(Email) > 1;
select Email from Person group by Email having count(Id) > 1

select Email from (select Email,count(Email) as num1 from Person group by Email ) as num where  num1 > 1
