package main

import "fmt"

func main() {
  var deg_f, deg_c float64

  fmt.Print ("Input Fahrenheit: ")
  fmt.Scanf("%f", &deg_f)
  
  deg_c = (deg_f - 32) * 5/9
  fmt.Print ("Temp C:", deg_c)
}
