package main

import (
    "fmt"
)

type contactInfo struct {
  email   string
  zipCode int
}

type person struct {
  firstName string
  lastName  string
  contactInfo       // declaring this way is the same as declaring "contactInfo contactInfo"
}

func (p person) print() {
  fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName (newFirstName string) {
  (*pointerToPerson).firstName = newFirstName
}

func main() {
  jim := person {
    firstName: "Jim",
    lastName: "Party",
    contactInfo: contactInfo {
      email: "jim@gmail.com",
      zipCode: 94000,
    },
  }

  jim.updateName("Jimbo")   // Golang allows the calling of receiver functions with either a pointer or with the root type (in this case, type `person`
  jim.print()
}
