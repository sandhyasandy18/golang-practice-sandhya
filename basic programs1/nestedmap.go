package main
import "fmt"
func main() {
	superhero := map[string]map[string]string {

		"Superman" :map[string]string {
		"RealName" :"Clark Kent",
		"City"     :"Metropolis",
		},

		"Batman"   :map[string] string {
		"RealName" :"Bruce Wyane",
		"City"     :"Metropolis",
		},
	}

	if temp, hero :=superhero["Superman"];hero {
		fmt.Println(temp["RealName"],temp["City"])
	}
}
