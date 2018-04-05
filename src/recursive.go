package main
import ( 
	"fmt"
        "bytes"
	"time"
)
func recursiveOut(depth, step, maxdepth int) {
        
        depth = depth + step
        if depth  > maxdepth {
                printDepth(maxdepth)
                return
        }   
        printDepth(depth)
	time.Sleep(300 * time.Millisecond)
        recursiveOut(depth, step, maxdepth)
}
func recursiveIn(depth, step int) {
        depth = depth - step
        if depth == 0 { 
                fmt.Println("")
                return
        }   
        printDepth(depth)
	time.Sleep(300 * time.Millisecond)
        recursiveIn(depth, step)
}
func printDepth(size int) {
    var buffer bytes.Buffer
    for i:= 0; i < size; i++ {
	    buffer.WriteString("*")
    }
    fmt.Println(buffer.String()) 
}
func main() {
	recursiveOut(0, 1, 100)
	recursiveIn(100, 1)

}

