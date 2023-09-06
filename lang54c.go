package main
import "fmt"
import "math"

func main() {
  var dia, r, circ, area float64
  var pi = 3.14159

  fmt.Print("Enter the diameter: ")
  fmt.Scanf("%f", &dia)

  r = dia / 2
  circ = 2 * r * pi
  area = pi * (math.Pow(r, 2))

  fmt.Printf("Radius: %g\n", r)
  fmt.Printf("Area: %.3f\n", area)
  fmt.Printf("Circumference: %.3f\n", circ)
}
