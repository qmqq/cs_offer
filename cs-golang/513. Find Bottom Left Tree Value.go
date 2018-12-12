package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
	"reflect"
	//"math/rand"
	//"sort"
	//"regexp"
)

/***************************************/
//some common data struct and function using in leetcode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ListNode struct {
	Val  int
	Next *ListNode
}

type Point struct {
	X int
	Y int
}

type Interval struct {
	Start int
	End   int
}

func useLib() {
	fmt.Println(strconv.Itoa(1))
	fmt.Println(strings.Compare("1", "2"))
	fmt.Println(math.Abs(1.0))
	fmt.Println(unicode.IsDigit('1'))
}

func buildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = nums[0]
	ch := make(chan *TreeNode, len(nums))
	ch <- root
	nums = nums[1:]
	for i := 0; i < len(nums); i++ {
		tree := <-ch
		if nums[i] == -1 {
			tree.Left = nil
		} else {
			tree.Left = &TreeNode{
				Val: nums[i],
			}
			ch <- tree.Left
		}
		i++
		if i == len(nums) || nums[i] == -1 {
			tree.Right = nil
		} else {
			tree.Right = &TreeNode{
				Val: nums[i],
			}
			ch <- tree.Right
		}
	}
	return root
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	root := &ListNode{
		Val: nums[0],
	}
	tmp := root
	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{
			Val: nums[i],
		}
		tmp = tmp.Next
	}
	return root

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/**************************************************************/
/*
type IntSlice []int

func (s IntSlice)Less(i,j int)bool{
	return s[i]<s[j]
}
func (s IntSlice)Swap(i,j int){
	s[i],s[j] = s[j],s[i]
}
func (s IntSlice)Len()int{
	return len(s)
}
*/
/**************************************************************/
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := root.Val
	MemBlock := make([][]*TreeNode, 2)
	MemBlock[0] = make([]*TreeNode, 0, 512)
	MemBlock[1] = make([]*TreeNode, 0, 512)
	MemBlock[0] = append(MemBlock[0], root)
	idx := 0
	for len(MemBlock[idx]) > 0 {
		ret = MemBlock[idx][0].Val
		for _, ptr := range MemBlock[idx] {
			if ptr.Left != nil {
				MemBlock[1-idx] = append(MemBlock[1-idx], ptr.Left)
			}
			if ptr.Right != nil {
				MemBlock[1-idx] = append(MemBlock[1-idx], ptr.Right)
			}
		}
		MemBlock[idx] = make([]*TreeNode, 0, 512)
		idx = 1 - idx

	}
	return ret
}

var maps = map[int]string{
	1:"1",
	2:"2",
}

func main() {
	var u ListNode
	t := reflect.Typeof(u)
	for i,n:= 0,t.NumField();i<n;i++{
		f := t.Field(i)
		fmt.Println(f.Name,f.Type)
	}

}
