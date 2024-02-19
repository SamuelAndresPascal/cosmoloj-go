package main

import(
"fmt"
"reflect"
)

type I interface {
Tutu() string
Tata() string
}

type A struct {
}

type B struct {
  A
}

type C struct {
  A
}

func (x *A) Tutu() string {
  return "A"
}

func (x *A) Tata() string {
  return x.Tutu()
}

func (x *B) Tutu() string {
  return "B"
}

func (x *C) Tutu() string {
  return "C"
}

func main() {
  fmt.Println("Hello, World!")

  var a = new(A)
  fmt.Println(reflect.TypeOf(a))
  fmt.Println(a.Tutu())
  fmt.Println(a.Tata())

  var b = new(B)
  fmt.Println(reflect.TypeOf(b))
  fmt.Println(b.Tutu())
  fmt.Println(b.Tata())

  var c = new(C)
  fmt.Println(reflect.TypeOf(c))
  fmt.Println(c.Tutu())
  fmt.Println(c.Tata())
}

