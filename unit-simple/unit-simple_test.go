package main

import (
    "testing"
    "math"
)

func TestTransformed(t *testing.T) {
  var m Unit = NewFundamentalUnit()
  var km Unit = m.ScaleMultiply(1000)
  var cm Unit = m.ScaleDivide(100)
  var cmToKm UnitConverter = cm.GetConverterTo(km)

  if math.Abs(0.00003 - cmToKm.Convert(3.0)) > 1e-10 {
    t.Fatal()
  }

  if math.Abs(3. - cmToKm.Inverse().Convert(0.00003)) > 1e-10 {
    t.Fatal()
  }
}

func TestDerived(t *testing.T) {
  var m Unit = NewFundamentalUnit()
  var km Unit = m.ScaleMultiply(1000)
  var km2 Unit = NewDerivedUnit([]Factor{km.Factor(2, 1)})
  var cm Unit = m.ScaleDivide(100)
  var cm2 Unit = NewDerivedUnit([]Factor{cm.Factor(2, 1)})
  var km2ToCm2 UnitConverter = km2.GetConverterTo(cm2)

  if math.Abs(3e10 - km2ToCm2.Convert(3.)) > 1e-10 {
    t.Fatal()
  }

  if math.Abs(3 - km2ToCm2.Inverse().Convert(3e10)) > 1e-10 {
    t.Fatal()
  }
}

