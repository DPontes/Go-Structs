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
    By specifically defining the variable name in the struct value assignment below,
    we guarantee that the value will be assigned to the desired varaible in the struct,
    regardless of the order of the variables inside the struct
  */
  alex := person{firstName: "Alex", lastName: "Anderson"}
  fmt.Println(alex)
}
