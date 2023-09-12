package main
import "fmt"
import "math"

func main() {
  var a, b, sum, diff, prod, avg, abs, max, min float64
	fmt.Print("Number 1: ")
  fmt.Scanf("%f", &a)
  fmt.Print("Number 2: ")
  fmt.Scanf("%f", &b)

  sum = a + b
  diff = a - b
  prod = a * b
  avg = sum / 2
  abs = math.Abs(diff)
  if a > b {
    max = a
    min = b
  } else {
    max = b
    min = a
  }
  fmt.Printf("Sum: %g\n", sum)
  fmt.Printf("Difference: %g\n", diff)
  fmt.Printf("Product: %g\n", prod)
  fmt.Printf("Average: %g\n", avg)
  fmt.Printf("Absolute value: %g\n", abs)
  fmt.Printf("Minimum: %g\n", min)
  fmt.Printf("Maximum: %g\n", max)
}
