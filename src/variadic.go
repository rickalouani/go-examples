package main
import "fmt"

func sumNumbers(numbers ...int) int {
	total := 0
	for _,number := range numbers {
		total += number
	}
	return total
}
func main() {
	fmt.Printf("total is %s \n", sumNumbers(1, 2, 4, 23))
}
