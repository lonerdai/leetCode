/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getNum(l*ListNode)int {
     res,t := 0,1
     
     for {
         if l==nil {
             break
         }else{
             res+=l.Val*t
             l = l.Next   
         }
         t*=10
     }
     return res
 }
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    sum,l,t:= getNum(l1)+getNum(l2),new(ListNode),new(ListNode)
    t=l
    for {
        if sum == 0 {
            break
        } else {
            t.Val = sum%10
            sum/=10
            if(sum>0){
                t.Next = new(ListNode)
                t= t.Next
            }
            
        }
    }
    return l
}