package main
import "fmt"
func main(){
	rect1:=Rectangle{10,5}
	fmt.Println(rect1.height)
	fmt.Println(rect1.width)
}
type Rectangle struct{
	height float64
	width float64
}
