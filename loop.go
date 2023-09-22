package main
import "fmt"


func main() {
	nums := []int{5, 6, 7, 8}
  for i, n := range nums {
    fmt.Print(i, n)
  }
}