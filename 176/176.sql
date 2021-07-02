// https://leetcode-cn.com/problems/second-highest-salary/

/**
 *
{"headers": {"Employee": ["Id", "Salary"]}, "rows": {"Employee": [[1, 100], [2, 200], [3, 300]]}}
{"headers": {"Employee": ["Id", "Salary"]}, "rows": {"Employee": [[1, 100]]}}
{"headers": {"Employee": ["Id", "Salary"]}, "rows": {"Employee": [[1, 100], [2, 100]]}}
 *
 **/
//无法解决 case 2 返回 null
//select Salary as SecondHighestSalary from Employee order by Salary desc limit 1,1

//无法解决 case 3 返回 null
//select (select  Salary  from Employee order by Salary desc limit 1,1) as SecondHighestSalary

// 临时表方案
select (select distinct Salary  from Employee order by Salary desc limit 1,1) as SecondHighestSalary

// ifnull 方案
select ifnull(select distinct Salary from Employee order by Salary desc limit 1,1),null) as SecondHighestSalary

// max 方案
select max(Salary) as SecondHighestSalary from Employee
where Salary < (select max(Salary) from Employee)
