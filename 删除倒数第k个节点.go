//给定一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
package main
import "fmt"
type ListNode struct {
  Val int
  Next *ListNode
}
/*
思路：
利用快慢指针，先让fast指针指向第k个节点,此时fast和slow间隔k个节点，然后fast和slow同时位移，当fast走到链表尾时，
slow在倒数第k个节点
*/
//生成链表
func makeListNode(s []int) *ListNode {
  if len(s)==0 {
    return nil
  }
  node:=&ListNode{
    Val:s[0],
  }
  temp:=node
  for i:=0;i<len(s);i++ {
    temp.Next=&ListNode{
      Val:s[i], 
    }
    temp=temp.Next
  }
  return node
}
func delete(head *ListNode, k int) *ListNode {
	fast, slow := &ListNode{}, &ListNode{} //定义快慢指针
	fast, slow = head, head
	//当链表为空时
	if head == nil {
		return nil
	}
	//当链表仅有一个节点时，删除的就是该节点
	if head.Next == nil {
		return nil
	}
	//将快指针移动到第k个节点处
	for i := 0; i < k-1; i++ {
		if fast.Next != nil {
			fast = fast.Next
		}
	}
	//定义一个新链表帮助储存slow前的节点
	helper := &ListNode{}

	//同时移动快慢指针,直到快指针移动到链表尾
	for fast.Next != nil {
		//更新helper的节点
		helper = slow
		fast = fast.Next
		slow = slow.Next
	}
	//直接将slow的前一个节点连接到slow的后一个节点，即删除slow
	helper.Next = slow.Next
	return head
}
//显示链表信息
func List(node *ListNode, a []int) {
	for node != nil {
		a = append(a, node.Val)
		node = node.Next
	}
	fmt.Println(a)
}
func main() {
  var s []int={1,2,3,4,5}
  var a []int
  var k int
  scanln(&k)
  node:=makeListNode(s)//创建链表
  delete(node,k)
  List(node,a)
  
}
