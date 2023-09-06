package main
import "fmt"

func main(){
  var leng, wid, area, perim int64
  fmt.Print("Length: ")
  fmt.Scanf("%d", &leng)
  fmt.Print("Width: ")
  fmt.Scanf("%d", &wid)

  area = leng * wid
  perim = leng * 2 + wid * 2

  fmt.Println("Area = ", area)
  fmt.Println("Perimeter = ", perim)
}