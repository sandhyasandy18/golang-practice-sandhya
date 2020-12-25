package main
import "fmt"
func main() {
	var name string="Sandhya"
	const pi float64=3.14412435
	win:=true
	x:=10
	
	fmt.Printf("%.3f  \n", pi) //precision
	fmt.Printf("%T \n",name)   //type 
	fmt.Printf("%t \n",win)    //boolean
	fmt.Printf("%d \n",x)      //integer 
	fmt.Printf("%b \n",20)     //binary value
	fmt.Printf("%c \n",33)     //ASCII value
	fmt.Printf("%x \n",15)     //Hex value
	fmt.Printf("%e \n",pi)     //Scientific notation
}