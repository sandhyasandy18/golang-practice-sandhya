package main
import("os" 
       "log")
func main()  {
	file,err:=os.Create("sample.txt")
	if err!=nil{
		log.Fatal(err)
	}
	file.WriteString("Hi my name is sandhya")
	file.Close()
	
}