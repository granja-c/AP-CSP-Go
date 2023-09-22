package main
import "fmt"

func main() {
  var egg, doz, ext int
  var price, tot float64
  fmt.Print("Number of eggs: ")
  fmt.Scanf("%d", egg)

  doz = egg / 12
  ext = egg % 12
  if egg <= 47 {
    price = 0.5
  } else if egg > 47 && egg <= 71 {
    price = 0.45
  } else if egg > 71 && egg <= 131 {
    price = 0.4
  } else {
    price = 0.35
  }
  tot = float64(doz) * price + float64(ext) * (price * 1/12)
  fmt.Println(tot)
  fmt.Printf("Total: $%g\n", tot)
  fmt.Print("blahblahblah")
}
