package main

import (
"fmt"
. "com.cosmoloj.go/unit_simple"
)

func main() {
  var i UnitConverter = Identity()
  fmt.Println(i)
  fmt.Println(i.Inverse())
  var c UnitConverter = NewUnitConverter(1, 2)

  fmt.Println(c)
  fmt.Println(c.Scale())
  fmt.Println(c.Offset())
  fmt.Println(c.Convert(4))
  fmt.Println(c.Inverse().Scale())
  fmt.Println(c.Inverse().Offset())
  fmt.Println(c.Inverse().Convert(4))

  fmt.Println("Hello, World!")
}
