# Simple Unit (implémentation Go)


## Utilisation

Utilisation des unités transformées :

```go
var m Unit = NewFundamentalUnit()
var km Unit = m.ScaleMultiply(1000)
var cm Unit = m.ScaleDivide(100)
var cmToKm UnitConverter = cm.GetConverterTo(km)

cmToKm.Convert(3.0) // 0.00003
cmToKm.Inverse().Convert(0.00003) // 3.
```

Utilisation des unités dérivées :

```go
var m Unit = NewFundamentalUnit()
var km Unit = m.ScaleMultiply(1000.)
var km2 Unit = NewDerivedUnit(km.Factor(2))
var cm Unit = m.ScaleDivide(100.)
var cm2 Unit = NewDerivedUnit(cm.Factor(2))
var km2ToCm2 UnitConverter = km2.GetConverterTo(cm2)

km2ToCm2.Convert(3.) // 3e10
km2ToCm2.Inverse().Convert(3e10) // 3.
```

Utilisation des unités dérivées en combinant les dimensions :

```go
var m Unit = NewFundamentalUnit()
var kg Unit = NewFundamentalUnit()
var g Unit = kg.ScaleDivide(1000.)
var ton Unit = kg.ScaleMultiply(1000.)
var gPerM2 = NewDerivedUnit(g, m.Factor(-2))
var km = m.ScaleMultiply(1000.)
var tonPerKm2 = NewDerivedUnit(ton, km.Factor(-2))
var cm = m.ScaleDivide(100.)
var tonPerCm2 = NewDerivedUnit(ton, cm.Factor(-2))
var gPerM2ToTonPerKm2 = gPerM2.GetConverterTo(tonPerKm2)
var gPerM2ToTonPerCm2 = gPerM2.GetConverterTo(tonPerCm2)

gPerM2ToTonPerKm2.Convert(1.) // 1.
gPerM2ToTonPerKm2.Convert(3.) // 3.
gPerM2ToTonPerCm2.Convert(1.) // 1e-10
gPerM2ToTonPerCm2.Convert(3.) // 3e-10
gPerM2ToTonPerCm2.Offset() // 0.
gPerM2ToTonPerCm2.Scale() // 1e-10
gPerM2ToTonPerCm2.Inverse().Offset() // 0.
gPerM2ToTonPerCm2.Inverse().Convert(3e-10) // 3.
```

Utilisation des températures (conversions affines et linéaires) :

```go
var k Unit = NewFundamentalUnit()
var c Unit = k.Shift(273.15)
var kToC UnitConverter = k.GetConverterTo(c)

kToC.Convert(0.) // -273.15
kToC.Inverse().Convert(0.) // 273.15

// en combinaison avec d'autres unites, les conversions d'unites de temperatures doivent devenir lineaires
var m Unit = NewFundamentalUnit()
var cPerM Unit = NewDerivedUnit(c, m.Factor(-1))
var kPerM Unit = NewDerivedUnit(k, m.Factor(-1))
var kPerMToCPerM UnitConverter = kPerM.GetConverterTo(cPerM)

kPerMToCPerM.Convert(3.)) > 1e-10 // 3.
kPerMToCPerM.Inverse().Convert(3.) // 3.
```

Utilisation des conversions non décimales :

```go
var m Unit = NewFundamentalUnit()
var km Unit = m.ScaleMultiply(1000.)

var s Unit = NewFundamentalUnit()
var h Unit = s.ScaleMultiply(3600.)

var mPerS Unit = NewDerivedUnit(m, s.Factor(-1))
var kMPerH Unit = NewDerivedUnit(km, h.Factor(-1))

var mPerSToKMPerH UnitConverter = mPerS.GetConverterTo(kMPerH)

mPerSToKMPerH.Convert(100.) // 360.
mPerSToKMPerH.Inverse().Convert(18.) // 5.
```

