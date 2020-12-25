package main
import "fmt"
func main() {
	age:=26
	switch age {
	case 17:fmt.Println("not eligible for voting")
	case 18:fmt.Println("eligible for voting")
	case 60:fmt.Println("eligible for senior citizen")
	default:fmt.Println("invalid value try again")		
	}
}