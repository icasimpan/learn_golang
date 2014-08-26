package main

import "fmt"

func main() {
  for i:=1; i<=100; i++ {
   // for multiples of 3, "Fizz" instead of number
   // for multiples of 5, "Buzz" instead of number
   // multiple of both 3 & 5, "FizzBuzz"
   if i%3 == 0 && i%5 != 0 {
     fmt.Println(i, "Fizz")
   } else if i%5 == 0 && i%3 != 0 {
     fmt.Println(i, "Buzz")
   } else if i%5 ==0 && i%3 == 0 {
     fmt.Println(i, "FizzBuzz")
   } else {
     fmt.Println(i)
   }
  }// END: for loop
}// END: main()
