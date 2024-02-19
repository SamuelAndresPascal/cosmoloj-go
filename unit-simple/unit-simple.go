package main

import (
"fmt"
"math"
)

type UnitConverter interface {
  Scale() float64
  Offset() float64
  Inverse() UnitConverter
  Linear() UnitConverter
  LinearPow(power float64) UnitConverter
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

type FundamentalUniter interface {
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
  inverse UnitConverter
}

func _NewUnitConverter(scale float64, offset float64, inverse UnitConverter) UnitConverter {
  result := new(UnitConvert)
  result.scale = scale
  result.offset = offset
  if inverse == nil {
    result.inverse = _NewUnitConverter(1 / scale, -offset / scale, result)
  } else {
    result.inverse = inverse
  }
  return result
}

func NewUnitConverter(scale float64, offset float64) UnitConverter {
  return _NewUnitConverter(scale, offset, nil)
}

type AbstractUnit struct {
  Uniter
}

type FundamentalUnit struct {
  FundamentalUniter
  AbstractUnit
}

type TransformedUnit struct {
  TransformedUniter
  AbstractUnit
}

type DerivedUnit struct {
  DerivedUniter
  AbstractUnit
}

func (x UnitConvert) Scale() float64 {
  return x.scale
}

func (x UnitConvert) Offset() float64 {
  return x.offset
}

func (x UnitConvert) Inverse() UnitConverter {
  return x.inverse
}

func (x UnitConvert) Linear() UnitConverter {
  if x.offset == 0 {
    return x
  }
  return NewUnitConverter(x.scale, 0)
}

func (x UnitConvert) LinearPow(power float64) UnitConverter {
  if x.offset == 0 && power == 1 {
    return x
  }
  return NewUnitConverter(math.Pow(x.scale, power), 0)
}

func (x UnitConvert) Convert(value float64) float64 {
  return value * x.scale + x.offset
}

func (x UnitConvert) Concatenate(converter UnitConverter) UnitConverter {
  return NewUnitConverter(converter.Scale() * x.Scale(), x.Convert(converter.Offset()))
}


func main() {
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
