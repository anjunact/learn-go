package main
import(
	"unicode/utf8"
	"fmt"
	 "github.com/anjunact/learn-go/stack"
)
func main(){
	var   s stack.Stack

	s.Push(25)
	fmt.Printf("statck %v\n",s)
}
func test(n int)  {
	for  i := 0 ; i < n ; i++{
		println(i)
	}
}
func test2(n int)  {
	i:=0
	A: if i<n {
		println(i)
	}
	i++
	goto A
}
func test3(n int)  {
	for i:=0 ;i< n ;i++{
		for j:=0;j<i;j++{
			print("A")
		}
		print("\n")
	}
}
func test4(s string)  {
	c := utf8.RuneCountInString(s)
	println(c)
	a := [...]int{0,1,2,3,4,5}
	fmt.Printf("%v",a)
	println(len([]byte(s)))
}
