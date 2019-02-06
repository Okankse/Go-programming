
package main

import "fmt"

func main() {
  var array [5]float64
      array[0] = 76
      array[1] = 54
      array[2] = 5
      array[3] = 35
      array[4] = 90

      var sum float64 = 0
      for i := 0; i < 5; i++ {
          sum += array[i]
      }

      fmt.Println(sum / 5)
   }
