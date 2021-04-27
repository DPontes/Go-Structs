package main

import (
    "fmt"
)

type person struct {
  firstName string
  lastName string
}

func main() {
  /*
    when declaring an empty variable, Golang assign a zero-value to each of the fields in the struct
    In the case of strings, the 'zero-value' is an empty string ""
  */
  var alex person

  alex.firstName = "Alex"
  alex.lastName = "Andersson"

  fmt.Println(alex)
  fmt.Printf("%+v", alex)   // prints out the names of the variables in the struct 'alex' of type 'person'
}
