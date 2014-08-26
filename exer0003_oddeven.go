package main

import "fmt"

func main() {
  for i:=1; i<=10; i++ {
    fmt.Print (i, ": ")
    test_oddeven:=i%2
    if test_oddeven == 0 {
      fmt.Println ("even")
    }else if test_oddeven == 1 {
      fmt.Println ("ODD")
    }
  }
}
