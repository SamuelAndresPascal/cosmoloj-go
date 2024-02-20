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

type Factor interface {
  Dim() Unit
  Numerator() int32
  Denominator() int32
  Power() float64
}

type Unit interface {
  Factor
  GetConverterTo(target Unit) UnitConverter
  ToBase() UnitConverter
  Shift(value float64) Unit
  ScaleMultiply(value float64) Unit
  ScaleDivide(value float64) Unit
  Factor(numerator int32, denominator int32) Factor
}

type FundamentalUnit interface {
  Unit
}

type TransformedUnit interface {
  Unit
  ToReference() UnitConverter
  Reference() Unit
}

type DerivedUnit interface {
  Unit
  Definition() []Factor
}

type UnitConverterImpl struct {
  scale float64
  offset float64
  inverse UnitConverter
}

func _NewUnitConverter(scale float64, offset float64, inverse UnitConverter) UnitConverter {
  result := new(UnitConverterImpl)
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

type FactorImpl struct {
  unit Unit
  numerator int32
  denominator int32
}

func NewFactor(unit Unit, numerator int32, denominator int32) Factor {
  return &FactorImpl{unit, numerator, denominator}
}

type UnitImpl struct {
  FactorImpl
  this Unit
}

type FundamentalUnitImpl struct {
  UnitImpl
}

func NewFundamentalUnit() FundamentalUnit {
  var result = new(FundamentalUnitImpl)
  result.this = result
  result.unit = result.this
  result.numerator = 1
  result.denominator = 1
  return result
}

type TransformedUnitImpl struct {
  UnitImpl
  toReference UnitConverter
  reference Unit
}

func NewTransformedUnit(toReference UnitConverter, reference Unit) TransformedUnit {
  var result = new(TransformedUnitImpl)
  result.this = result
  result.unit = result.this
  result.numerator = 1
  result.denominator = 1
  result.toReference = toReference
  result.reference = reference
  return result
}

type DerivedUnitImpl struct {
  UnitImpl
  definition []Factor
}

func NewDerivedUnit(definition []Factor) DerivedUnit {
  var result = new(DerivedUnitImpl)
  result.this = result
  result.unit = result.this
  result.numerator = 1
  result.denominator = 1
  result.definition = definition
  return result
}

func (x *UnitConverterImpl) Scale() float64 {
  return x.scale
}

func (x *UnitConverterImpl) Offset() float64 {
  return x.offset
}

func (x *UnitConverterImpl) Inverse() UnitConverter {
  return x.inverse
}

func (x *UnitConverterImpl) Linear() UnitConverter {
  if x.offset == 0 {
    return x
  }
  return NewUnitConverter(x.scale, 0)
}

func (x *UnitConverterImpl) LinearPow(power float64) UnitConverter {
  if x.offset == 0 && power == 1 {
    return x
  }
  return NewUnitConverter(math.Pow(x.scale, power), 0)
}

func (x *UnitConverterImpl) Convert(value float64) float64 {
  return value * x.scale + x.offset
}

func (x *UnitConverterImpl) Concatenate(converter UnitConverter) UnitConverter {
  return NewUnitConverter(converter.Scale() * x.Scale(), x.Convert(converter.Offset()))
}

func (x *FactorImpl) Dim() Unit {
  return x.unit
}

func (x *FactorImpl) Numerator() int32 {
  return x.numerator
}

func (x *FactorImpl) Denominator() int32 {
  return x.denominator
}

func (x *FactorImpl) Power() float64 {
  return float64(x.numerator) / float64(x.denominator)
}

func (x *UnitImpl) GetConverterTo(target Unit) UnitConverter {
  return target.ToBase().Inverse().Concatenate(x.this.ToBase())
}

func (x *UnitImpl) Shift(value float64) Unit {
  return NewTransformedUnit(NewUnitConverter(1, value), x.this)
}

func (x *UnitImpl) ScaleMultiply(value float64) Unit {
  return NewTransformedUnit(NewUnitConverter(value, 0), x.this)
}

func (x *UnitImpl) ScaleDivide(value float64) Unit {
  return x.ScaleMultiply(1 / value)
}

func (x *UnitImpl) Factor(numerator int32, denominator int32) Factor {
  return NewFactor(x.this, numerator, denominator)
}

func (x *FundamentalUnitImpl) ToBase() UnitConverter {
  return NewUnitConverter(1, 0)
}

func (x *TransformedUnitImpl) ToBase() UnitConverter {
  return x.Reference().ToBase().Concatenate(x.ToReference())
}

func (x *TransformedUnitImpl) ToReference() UnitConverter {
  return x.toReference
}

func (x *TransformedUnitImpl) Reference() Unit {
  return x.reference
}

func (x *DerivedUnitImpl) Definition() []Factor {
  return x.definition
}

func (x *DerivedUnitImpl) ToBase() UnitConverter {
  var transform = NewUnitConverter(1, 0)
  for _, factor := range x.definition {
    transform = factor.Dim().ToBase().LinearPow(factor.Power()).Concatenate(transform) 
  }
  return transform
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
