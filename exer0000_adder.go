package main

import "fmt"

func main() {
  var (
       num1 = 0
       num2 = 0
       sum  = 0
      )

  fmt.Print("Enter num1: ")
  fmt.Scanf("%d", &num1)
  fmt.Print("Enter num2: ")
  fmt.Scanf("%d", &num2)
  sum = num1 + num2
  
  fmt.Println("Sum is: ", sum)
}
