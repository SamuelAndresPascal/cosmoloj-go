package main

import "fmt"




type UnitConverter interface {
    Scale() float64
    Offset() float64
    Inverse() UnitConverter
    Linear() UnitConverter
    LinearPow() UnitConverter
    Convert(value float64) float64
    Concatenate(converter UnitConverter) UnitConverter
}

type Factorer interface {
    Dim() Uniter
    Numerator() int32
    Denominator() int32
    Power() float64
}

type Uniter interface {
    Factorer
    GetConverterTo(unit Uniter) UnitConverter
    ToBase() UnitConverter
    Shift(value float64) Uniter
    ScaleMultiply(value float64) Uniter
    ScaleDivide(value float64) Uniter
    Factor(numerator int32, denominator int32) Factorer
}

type FuntamentalUniter interface {
    Uniter
}

type TransformedUniter interface {
    Uniter
}

type DerivedUniter interface {
    Uniter
}


type UnitConvert struct {
    scale float64
    offset float64

}

func (x UnitConvert) Scale() float64 {
    return x.scale
}

func (x UnitConvert) Offset() float64 {
    return x.offset
}


func main() {
    c := UnitConvert{1, 2}

    fmt.Println(c)
    fmt.Println(c.scale)
    fmt.Println(c.offset)
    fmt.Println(c.Scale())
    fmt.Println(c.Offset())

    fmt.Println("Hello, World!")
}
