package main
import "fmt"

func main() {
  var og, w, res string
  var a, b int
	fmt.Print("String: ")
  fmt.Scanln(&og)
  og += " "
  w = ""
  res = ""
  a = 0
  b = 0
  for i := 0; i < len(og); i++{
    if og[i] != ' ' {
      a += 1
    } else {
      w = og[b:a]
      res = w + " " + res
      w = ""
      b = a
      a = 0
    }
  }
  fmt.Printf(res)
}
