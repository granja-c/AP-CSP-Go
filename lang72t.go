package main
import "fmt"

func main() {
	var t1, t2, a, b, hrs, min int
  fmt.Print("Time 1: ")
  fmt.Scanf("%d", &t1)
  fmt.Print("Time 2: ")
  fmt.Scanf("%d", &t2)

  if t1 > t2 {
    a = t1
    b = t2
  } else {
    a = t2
    b = t1
  }
  hrs = (a - b) / 100
  min = (a - b) % 100

  fmt.Print(hrs, " hours ", min, " minutes")
}
