package main
import "fmt"

func main() {
  var cs, p, tot float64
	fmt.Print("Enter the number of copies: ")
  fmt.Scanf("%f", &cs)
  if cs <= 99 {
    p = 0.3
  } else if cs > 99 && cs <= 499 {
    p = 0.28
  } else if cs > 799 && cs <= 749 {
    p = 0.27
  } else if cs > 749 && cs <= 1000 {
    p = 0.26
  } else {
    p = 0.25
  }
  tot = p * cs
  fmt.Printf("Price per copy: $%g\n", p)
  fmt.Printf("Total: $%g", tot)
}
