package main
import "fmt"
import "math"

func main() {
	var p, r, n, t, intr, tot float64
  fmt.Print("Amount saved: ")
  fmt.Scanf("%f", &p)
  fmt.Print("Interest rate: ")
  fmt.Scanf("%f", &r)
  fmt.Print("Number of times compoounded per year: ")
  fmt.Scanf("%f", &n)
  fmt.Print("Number of days at interest: ")
  fmt.Scanf("%f", &t)

  intr = p * (math.Pow(1 + ((0.01 * r) / n), ((n * t) / 365)) - 1)
  tot = intr + p

  fmt.Printf("Interest: $%.2f\n", intr)
  fmt.Printf("Total: $%.2f", tot)
  
}
