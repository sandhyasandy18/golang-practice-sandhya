package main
import (
  "fmt"
  "regexp"
)

func main() { 
  re := regexp.MustCompile("p([a-z]+)ch")
  fmt.Println(re.FindStringSubmatch("peach punch"))
  fmt.Println(re.FindStringSubmatch("cricket"))
}