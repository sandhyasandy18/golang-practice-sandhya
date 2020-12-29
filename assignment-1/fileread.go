package main
import ("fmt"
		"os"
		"log"
		"io/ioutil")
func main()  {
     file,err:=os.Create("sample.txt")
	if err!=nil{
	log.Fatal(err)
}
	file.WriteString("Hi my name is sandhya")
	file.Close()
	stream, err:=ioutil.ReadFile("sample.txt")
	if err!=nil{
		log.Fatal(err)
	}
	s1:=string(stream)
	fmt.Println(s1)
}