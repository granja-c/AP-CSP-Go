package main
import "fmt"

func main() {
	var n1, n2, n3, n4, sum int64
  var avg float64
  fmt.Print("Number 1: ")
  fmt.Scanf("%d", &n1)
  fmt.Print("Number 2: ")
  fmt.Scanf("%d", &n2)
  fmt.Print("Number 3: ")
  fmt.Scanf("%d", &n3)
  fmt.Print("Number 4: ")
  fmt.Scanf("%d", &n4)

  sum = n1 + n2 + n3 + n4
  var sum2 = float64(sum)
  avg = sum2 / 4.0

  fmt.Print("Sum = ", sum)
  fmt.Print("\nAverage = ", avg)
}
