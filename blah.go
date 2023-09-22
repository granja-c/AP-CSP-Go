package main
import "fmt"

func main() {
  var a, b float64
  fmt.Print("Enter:")
  fmt.Scanf("%f", &b)
  a = 0.25 / b
  fmt.Print(a)
}
