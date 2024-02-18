package main

import (
    "testing"
    "fmt"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
    var c UnitConverter = &UnitConvert{1, 2}

    fmt.Println(c)
    fmt.Println(c.Scale())
    fmt.Println(c.Offset())
    fmt.Println(c.Convert(4))
    
    
    if c.Convert(4) != 6 {
        t.Fatalf(`fail`)
    }
}
