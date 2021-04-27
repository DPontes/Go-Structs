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
   The order of values inside {} corresponds to the order of variables in the struct.
   This is not the best way to assign values to a struct, as the order of the variables
   declaration inside the struct can be changes at some point
  */
  alex := person{"Alex", "Anderson"}

}
