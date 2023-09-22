package main
import "fmt"

func main() {
  var pw, h, l, w, vol float64
	fmt.Print("Enter the weight: ")
  fmt.Scanf("%f", &pw)
  fmt.Print("Enter the height: ")
  fmt.Scanf("%f", &h)
  fmt.Print("Enter the length: ")
  fmt.Scanf("%f", &l)
  fmt.Print("Enter the width: ")
  fmt.Scanf("%f", &w)
  vol = h * l * w

  if pw > 27 && vol > 100000 {
    fmt.Print("Too heavy and too large")
  } else if pw > 27 {
    fmt.Print("Too heavy")
  } else if vol > 100000 {
    fmt.Print("Too large")
  }
}