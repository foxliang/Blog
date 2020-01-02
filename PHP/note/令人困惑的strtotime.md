经常会有人被strtotime结合-1 month, +1 month, next month的时候搞得很困惑, 然后就会觉得这个函数有点不那么靠谱, 动不动就出问题. 用的时候就会很慌…

这不, 刚刚就有人在微博上又问我:

鸟哥，今天是2018-07-31 执行代码:

```
date("Y-m-d",strtotime("-1 month"))
怎么输出是2018-07-01?
```

好的吧, 虽然这个问题看起来很迷惑, 但从内部逻辑上来说呢, 其实是”对”的, 你先别着急哈, 让我慢慢讲:

我们来模拟下date内部的对于这种事情的处理逻辑:

1. 先做-1 month, 那么当前是07-31, 减去一以后就是06-31.

2. 再做日期规范化, 因为6月没有31号, 所以就好像2点60等于3点一样, 6月31就等于了7月1

是不是逻辑很”清晰”呢? 我们也可以手动验证第二个步骤, 比如:

```
var_dump(date("Y-m-d", strtotime("2017-06-31")));
//输出2017-07-01
```
也就是说, 只要涉及到大小月的最后一天, 都可能会有这个迷惑, 我们也可以很轻松的验证类似的其他月份, 印证这个结论:

```
var_dump(date("Y-m-d", strtotime("-1 month", strtotime("2017-03-31"))));
//输出2017-03-03
var_dump(date("Y-m-d", strtotime("+1 month", strtotime("2017-08-31"))));
//输出2017-10-01
var_dump(date("Y-m-d", strtotime("next month", strtotime("2017-01-31"))));
//输出2017-03-03
var_dump(date("Y-m-d", strtotime("last month", strtotime("2017-03-31"))));
//输出2017-03-03
```
那怎么办呢?

从PHP5.3开始呢, date新增了一系列修正短语, 来明确这个问题, 那就是”first day of” 和 “last day of”, 也就是你可以限定好不要让date自动”规范化”:

```
var_dump(date("Y-m-d", strtotime("last day of -1 month", strtotime("2017-03-31"))));
//输出2017-02-28
var_dump(date("Y-m-d", strtotime("first day of +1 month", strtotime("2017-08-31"))));
////输出2017-09-01
var_dump(date("Y-m-d", strtotime("first day of next month", strtotime("2017-01-31"))));
////输出2017-02-01
var_dump(date("Y-m-d", strtotime("last day of last month", strtotime("2017-03-31"))));
////输出2017-02-28
```

那如果是5.3之前的版本(还有人用么?), 你可以使用mktime之类的, 把所有的日子忽略掉, 比如都限定为每月1号就可以了, 只不过就不如直接用first day来的更加优雅.

源：http://www.laruence.com/2018/07/31/3207.html
