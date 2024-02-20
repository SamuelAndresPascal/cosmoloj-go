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
  var km Unit = m.ScaleMultiply(1000.)
  var km2 Unit = NewDerivedUnit([]Factor{km.Factor(2, 1)})
  var cm Unit = m.ScaleDivide(100.)
  var cm2 Unit = NewDerivedUnit([]Factor{cm.Factor(2, 1)})
  var km2ToCm2 UnitConverter = km2.GetConverterTo(cm2)

  if math.Abs(3e10 - km2ToCm2.Convert(3.)) > 1e-10 {
    t.Fatal()
  }

  if math.Abs(3. - km2ToCm2.Inverse().Convert(3e10)) > 1e-10 {
    t.Fatal()
  }
}

func TestCombinedDimensionDerived(t *testing.T) {

  var m Unit = NewFundamentalUnit()
  var kg Unit = NewFundamentalUnit()
  var g Unit = kg.ScaleDivide(1000.)
  var ton Unit = kg.ScaleMultiply(1000.)
  var gPerM2 = NewDerivedUnit([]Factor{g, m.Factor(-2, 1)})
  var km = m.ScaleMultiply(1000.)
  var tonPerKm2 = NewDerivedUnit([]Factor{ton, km.Factor(-2, 1)})
  var cm = m.ScaleDivide(100.)
  var tonPerCm2 = NewDerivedUnit([]Factor{ton, cm.Factor(-2, 1)})
  var gPerM2ToTonPerKm2 = gPerM2.GetConverterTo(tonPerKm2)
  var gPerM2ToTonPerCm2 = gPerM2.GetConverterTo(tonPerCm2)

  if math.Abs(1. - gPerM2ToTonPerKm2.Convert(1.)) > 1e-10 {
    t.Fatal()
  }

  if math.Abs(3. - gPerM2ToTonPerKm2.Convert(3.)) > 1e-10 {
    t.Fatal()
  }

  if math.Abs(1e-10 - gPerM2ToTonPerCm2.Convert(1.)) > 1e-20 {
    t.Fatal()
  }

  if math.Abs(3e-10 - gPerM2ToTonPerCm2.Convert(3.)) > 1e-20 {
    t.Fatal()
  }

  if 0. != gPerM2ToTonPerCm2.Offset() {
    t.Fatal()
  }

  if 1e-10 != gPerM2ToTonPerCm2.Scale() {
    t.Fatal()
  }

  if 0. != gPerM2ToTonPerCm2.Inverse().Offset() {
    t.Fatal()
  }

  if math.Abs(3. - gPerM2ToTonPerCm2.Inverse().Convert(3e-10)) > 1e-10 {
    t.Fatal()
  }
}

func TestTemperatures(t *testing.T) {

  var k Unit = NewFundamentalUnit()
  var c Unit = k.Shift(273.15)
  var kToC UnitConverter = k.GetConverterTo(c)

  if math.Abs(-273.15 - kToC.Convert(0.)) > 1e-10 {
    t.Fatal()
  }
  
  if math.Abs(273.15 - kToC.Inverse().Convert(0.)) > 1e-10 {
    t.Fatal()
  }
  
  // en combinaison avec d'autres unites, les conversions d'unites de temperatures doivent devenir lineaires
  var m Unit = NewFundamentalUnit()
  var cPerM Unit = NewDerivedUnit([]Factor{c, m.Factor(-1, 1)})
  var kPerM Unit = NewDerivedUnit([]Factor{k, m.Factor(-1, 1)})
  var kPerMToCPerM UnitConverter = kPerM.GetConverterTo(cPerM)
  
  if math.Abs(3. - kPerMToCPerM.Convert(3.)) > 1e-10 {
    t.Fatal()
  }
  
  if math.Abs(3. - kPerMToCPerM.Inverse().Convert(3.)) > 1e-10 {
    t.Fatal()
  }
}

func TestSpeed(t *testing.T) {

  var m Unit = NewFundamentalUnit()
  var km Unit = m.ScaleMultiply(1000.)
  
  var s Unit = NewFundamentalUnit()
  var h Unit = s.ScaleMultiply(3600.)
  
  var mPerS Unit = NewDerivedUnit([]Factor{m, s.Factor(-1, 1)})
  var kMPerH Unit = NewDerivedUnit([]Factor{km, h.Factor(-1, 1)})
  
  var mPerSToKMPerH UnitConverter = mPerS.GetConverterTo(kMPerH)
  
  if math.Abs(360. - mPerSToKMPerH.Convert(100.)) > 1e-10 {
    t.Fatal()
  }
  
  if math.Abs(5. - mPerSToKMPerH.Inverse().Convert(18.)) > 1e-10 {
    t.Fatal()
  }
}

