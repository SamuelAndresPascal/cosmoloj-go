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
  GetConverterTo(target Uniter) UnitConverter
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
  ToReference() UnitConverter
  Reference() Uniter
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

type FactorStruct struct {
  unit Uniter
  numerator int32
  denominator int32
}

func NewFactor(unit Uniter, numerator int32, denominator int32) Factorer {
  return &FactorStruct{unit, numerator, denominator}
}

type Unit struct {
  FactorStruct
  this Uniter
}

type FundamentalUnit struct {
  Unit
}

func NewFundamentalUniter() FundamentalUniter {
  var result = new(FundamentalUnit)
  result.this = result
  return result
}

type TransformedUnit struct {
  Unit
  toReference UnitConverter
  reference Uniter
}

func NewTransformedUniter(toReference UnitConverter, reference Uniter) TransformedUniter {
  var result = new(TransformedUnit)
  result.toReference = toReference
  result.this = result
  result.reference = reference
  return result
}

type DerivedUnit struct {
  Unit
}

func (x *UnitConvert) Scale() float64 {
  return x.scale
}

func (x *UnitConvert) Offset() float64 {
  return x.offset
}

func (x *UnitConvert) Inverse() UnitConverter {
  return x.inverse
}

func (x *UnitConvert) Linear() UnitConverter {
  if x.offset == 0 {
    return x
  }
  return NewUnitConverter(x.scale, 0)
}

func (x *UnitConvert) LinearPow(power float64) UnitConverter {
  if x.offset == 0 && power == 1 {
    return x
  }
  return NewUnitConverter(math.Pow(x.scale, power), 0)
}

func (x *UnitConvert) Convert(value float64) float64 {
  return value * x.scale + x.offset
}

func (x *UnitConvert) Concatenate(converter UnitConverter) UnitConverter {
  return NewUnitConverter(converter.Scale() * x.Scale(), x.Convert(converter.Offset()))
}

func (x *FactorStruct) Dim() Uniter {
  return x.unit
}

func (x *FactorStruct) Numerator() int32 {
  return x.numerator
}

func (x *FactorStruct) Denominator() int32 {
  return x.denominator
}

func (x *FactorStruct) Power() float64 {
  return float64(x.numerator) / float64(x.denominator)
}

func (x *Unit) GetConverterTo(target Uniter) UnitConverter {
  return target.ToBase().Inverse().Concatenate(x.this.ToBase())
}

func (x *Unit) Shift(value float64) Uniter {
  return NewTransformedUniter(NewUnitConverter(1, value), x.this)
}

func (x *Unit) ScaleMultiply(value float64) Uniter {
  return NewTransformedUniter(NewUnitConverter(value, 0), x.this)
}

func (x *Unit) ScaleDivide(value float64) Uniter {
  return x.ScaleMultiply(1 / value)
}

func (x *Unit) Factor(numerator int32, denominator int32) Factorer {
  return NewFactor(x.this, numerator, denominator)
}

func (x *FundamentalUnit) ToBase() UnitConverter {
  return NewUnitConverter(1, 0)
}

func (x *TransformedUnit) ToBase() UnitConverter {
  return x.Reference().ToBase().Concatenate(x.ToReference())
}

func (x *TransformedUnit) ToReference() UnitConverter {
  return x.toReference
}

func (x *TransformedUnit) Reference() Uniter {
  return x.reference
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
