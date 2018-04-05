package main
import "fmt"

func recursiveOut(depth, step, maxdepth int) {
        
	depth = depth + step
	if depth  > maxdepth {
		fmt.Println("Already at the Max!")
		return
	}
        fmt.Printf(" %d\n", depth)
	recursiveOut(depth, step, maxdepth)
}
func recursiveIn(depth, step int) {
	depth = depth - step
	if depth == 0 {
		fmt.Println("Reached Bottom!")
		return
	}
	fmt.Printf("I am at layer %d\n", depth)
	recursiveIn(depth, step)
}
func main() {
	recursiveOut(0, 1, 50)
	recursiveIn(50, 1)
}

          
