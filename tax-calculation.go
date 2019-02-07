
package main

import "fmt"
// %10 for under 1.000
// %20 for over 1.000
// %30 for over 10.000
func tax(value int) float64 {
   taxValue :=0.0
  if value<=1000{
    taxValue =float64(value)*0.1

  } else if (value>1000)&&(value<10000){
    taxValue = float64(value)*0.2

  } else {
    taxValue = float64(value)*0.3
  }
    return taxValue
}
func main() {
      var money int
      fmt.Print("Please enter the value of money: ")
      fmt.Scanf("%d",&money)
      //tax:=tax(money)
      fmt.Println(tax(money))

   }
