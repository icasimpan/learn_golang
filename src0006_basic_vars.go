package main

import "fmt"

func main() {
  var x string = "Hello"
  y:=" World"
  fmt.Println(x,y)
  x+=y
  fmt.Println("Written another way: ",x)
}
