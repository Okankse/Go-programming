package main

import "fmt"

func main() {
  var array [5]float64
      fmt.Println("Please enter 5 numbers: ")
      var sum float64 = 0
      fmt.Scan(&array[0],&array[1],&array[2],&array[3],&array[4],)
      for i := 0; i < 5; i++ {
          sum += array[i]
      }

      fmt.Println("Average: ",sum / 5)
   }
