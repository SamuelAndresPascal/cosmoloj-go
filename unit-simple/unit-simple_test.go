package main

import (
    "testing"
    "math"
)

func TestTransformed(t *testing.T) {
    var m = NewFundamentalUnit()
    var km = m.ScaleMultiply(1000)
    var cm = m.ScaleDivide(100)
    var cmToKm = cm.GetConverterTo(km)
    
    if math.Abs(0.00003 - cmToKm.Convert(3.0)) > 1e-10 {
        t.Fatal()
    }
    
    if math.Abs(3. - cmToKm.Inverse().Convert(0.00003)) > 1e-10 {
        t.Fatal()
    }
}
