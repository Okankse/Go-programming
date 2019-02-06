package main

import "fmt"

func main() {
  fmt.Print("Please enter a note: ")
  var input float64
  fmt.Scanf("%f", &input)  // entered number

  if input>=85 && input<=100{

  fmt.Println( input,":> Note: A ")
  fmt.Println("PASSED!!!")

}else if input<85 && input>=70 {

    fmt.Println(input,":> Note: B")
    fmt.Println("PASSED!!!")

   }else if input<70 && input>=55 {

    fmt.Println(input,":> Note: C")
    fmt.Println("PASSED!!!")

   }else if input<55 && input>=40 {

    fmt.Println(input,":> Note: D")
    fmt.Println("You should study!!!")

   }else  {

    fmt.Println(input,":> Note: F")
    fmt.Println("NOT PASSED!!!")
   }
}
