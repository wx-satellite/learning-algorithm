##  实现 strStr（Implement-StrStr）
题干如下：
>实现 strStr() 函数。
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
示例 1:
&nbsp;&nbsp;输入: haystack = "hello", needle = "ll"
&nbsp;&nbsp;输出: 2
示例 2:
&nbsp;&nbsp;输入: haystack = "aaaaa", needle = "bba"
&nbsp;&nbsp;输出: -1
说明:
当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
来源：力扣

本题有很多种解题方式，为了再巩固一下 **双指针解题思路** 这题我们还是选用这种思路来解。
建议再结合《12_移除元素》和 《11_删除排序数组中的重复项》两篇文章仔细体会一下 **双指针思路**。

## 解题思路
初始化 **i** 指向 **needle** 字符串的第一个字符， **j** 指向 **haystack** 字符串的第一个字符。比较 **needle[i]** 与 **haystack[j]** 的字符是否相等：相等则**i++、j++**；不想等则 **i** 重新指向 **needle** 字符串的第一个字符，**j** 指向上一次初始位置的后一个字符。
> 注： j 指向上一次初始位置的后一个字符的计算方式为： j = j - i + 1。可以自己造个数组，推算一下这个规律哟。

流程图如下：
![](https://cdn.learnku.com/uploads/images/202004/21/21280/9yPNvjVMe0.jpg!large)


