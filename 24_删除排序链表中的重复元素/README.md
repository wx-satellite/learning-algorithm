## 删除排序链表中的重复元素（Remove-Duplicates-From-Sorted-List）
题干：
>给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
示例 1:
&nbsp;&nbsp;输入: 1->1->2
&nbsp;&nbsp;输出: 1->2
示例 2:
&nbsp;&nbsp;输入: 1->1->2->3->3
&nbsp;&nbsp输出: 1->2->3
来源：力扣

还记得删除排序数组的重复项是如何解决的嘛，[戳这里](https://learnku.com/articles/43366 "戳这里")，思路都差不多。
> 其实力扣有很多相似的题目，只不过承载形式不一样而已。比如上面的删除重复项，先前是承载在数组上，这题则承载在链表上。

## 解题思路
初始化指针 **current**，使其指向链表头。

比较 current 的 val 值 与 current.next 的 val 值。如果 current.val 与 current.next.val 相等，则将 current.next 指向 current.next.next，相当于将原本 current.next 抛弃。如果不想等，则 current 指向 current.next。重复上面的操作。

流程图如下：
![](https://cdn.learnku.com/uploads/images/202005/13/21280/OCLACYrif6.jpg!large)


