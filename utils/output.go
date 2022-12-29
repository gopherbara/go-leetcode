package utils

import "fmt"

func Output(name string, output interface{}) {
	fmt.Printf("%s: %v \n", name, output)

}
func PrintFullNode(name string, l *ListNode) {
	fmt.Print(name, ": ")
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
