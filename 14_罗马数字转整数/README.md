## 罗马数字转整数（Roman-To-Integer）
题干：
>罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
字符 &nbsp;&nbsp;  数值 
I     &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;  &nbsp;&nbsp;&nbsp;    1
V     &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;     5
X     &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;      10
L     &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;    50
C     &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;       100
D     &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;         500
M     &nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;        1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。
通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。
示例 1:
&nbsp;&nbsp;输入: "LVIII"
&nbsp;&nbsp;输出: 58
&nbsp;&nbsp;解释: L = 50, V= 5, III = 3.
示例 2:
&nbsp;&nbsp;输入: "MCMXCIV"
&nbsp;&nbsp;输出: 1994
&nbsp;&nbsp;解释: M = 1000, CM = 900, XC = 90, IV = 4.
来源：力扣

题目有点长，耐心读一下哦，整体来说是比较简单的。

## 解题思路
根据题意我们可以知道单个罗马字母和数字的映射关系，例如 **V** 表示 **数字 5**，只是这里有一个问题： **罗马字母的组合形式对应的数字并不是单个罗马字母对应的数字之和**。 例如 **CD** 对应的数字是 **400**，而单个 **C** 对应的数字为 **100**，单个 **D** 对应的数字为 **500**。

根据规律可以发现（当然题干也明确说明了），当罗马数字中小数在大数左边时，所表示的数等于大数减去小数的值。 例如 **CD** 对应的数字是 **单个 D 对应的数值减去单个 C 对应的数值，即 500 - 100 = 400**。所以我们只需要遍历目标字符串，判断当前的罗马字符对应的数值 **cur** 与下一个罗马字符对应的数值 **next** 的大小，**cur** 小于 **next** 则 **减去 cur 的值**，反之则 **加上 cur 的值**。

流程图如下：
![](https://cdn.learnku.com/uploads/images/202004/23/21280/6Ys5DzmU6K.jpg!large)

