package main
import "fmt"


func main() {
	nums := []int{1, 2, 3, 4}
  for n := range nums {
    fmt.Print(n)
  }
}