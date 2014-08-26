package main

import "fmt"

func main() {
  fmt.Println("*** Number evenly divisible by 3 ***")
  for i:=1; i<=100; i++ {
    //print if evenly div by 3
    if i%3 == 0 {
      fmt.Println(i)
    }
  }// END: for loop
}// END: main()
