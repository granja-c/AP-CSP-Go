package main
import "fmt"

func main() {
  var e, p, tot float64
  var r, doz int
	fmt.Print("Number of eggs: ")
  fmt.Scanf("%f", e)

  if e <= 47 {
    p = 0.5
  } else if e > 47 && e <= 71 {
    p = 0.45
  } else if e > 71 && e <= 131 {
    p = 0.4
  } else {
    p = 0.35
  }
  doz = int(e) / 12
  r = int(e) % 12
  tot = (e - float64(r)) * p + float64(r) * (p / 12.0)
  fmt.Printf("The total price is: $%g \n", tot)
  tot = 3
  fmt.Print(doz)
  fmt.Print(tot)
}
