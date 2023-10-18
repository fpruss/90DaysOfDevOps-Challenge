package main
 
import "fmt"
 
func main() {
    // single condition for loop
    m := 1
    for m <= 5 {
        fmt.Printf("Hello %d times.\n",m)
        m = m + 1
    }
    // classic for loop example 
    for i := 6; i <= 10; i++ {
        fmt.Printf("Welcome %d times.\n",i)
    }
}
