/* Print average of above scores (use arrays): 98, 93, 77, 82, 83*/

package main

import "fmt"

func main() {
  var x [5]float64
  x[0] = 98
  x[1] = 93
  x[2] = 77
  x[3] = 82
  x[4] = 83

  var total float64;

  for i:=0; i<len(x); i++ {
    total += x[i]    
  }

  fmt.Println("Average is: ", total/float64(len(x)))
}//END: main()
