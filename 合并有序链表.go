//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
  res:=&ListNode {}
  temp:=res
  for list1!=nil && list2!=nil {
      if list1.Val<list2.Val {
          temp.Next=list1
          list1=list1.Next
      }else {
          temp.Next=list2
          list2=list2.Next
      }
      temp=temp.Next  //继续向下一个节点推进
  }
  if list1!=nil {
      temp.Next=list1
  }else {
      temp.Next=list2 
  }
  return res.Next
}
