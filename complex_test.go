// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package integral

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestComplexAddCommutative(t *testing.T) {
	f := func(x, y *Complex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Complex).Add(x, y)
		r := new(Complex).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexMulCommutative(t *testing.T) {
	f := func(x, y *Complex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Complex).Mul(x, y)
		r := new(Complex).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexNegConjCommutative(t *testing.T) {
	f := func(x *Complex) bool {
		// t.Logf("x = %v", x)
		l, r := new(Complex), new(Complex)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestComplexSubAntiCommutative(t *testing.T) {
	f := func(x, y *Complex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Complex), new(Complex)
		l.Sub(x, y)
		r.Sub(y, x)
		r.Neg(r)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Associativity

func TestComplexAddAssociative(t *testing.T) {
	f := func(x, y, z *Complex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Complex), new(Complex)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexMulAssociative(t *testing.T) {
	f := func(x, y, z *Complex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Complex), new(Complex)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestComplexAddZero(t *testing.T) {
	zero := new(Complex)
	f := func(x *Complex) bool {
		// t.Logf("x = %v", x)
		l := new(Complex).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexMulOne(t *testing.T) {
	one := &Complex{
		l: *big.NewInt(1),
	}
	f := func(x *Complex) bool {
		// t.Logf("x = %v", x)
		l := new(Complex).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexAddNegSub(t *testing.T) {
	f := func(x, y *Complex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Complex), new(Complex)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexAddScalDouble(t *testing.T) {
	f := func(x *Complex) bool {
		// t.Logf("x = %v", x)
		l, r := new(Complex), new(Complex)
		l.Add(x, x)
		r.Scal(x, big.NewInt(2))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestComplexNegInvolutive(t *testing.T) {
	f := func(x *Complex) bool {
		// t.Logf("x = %v", x)
		l := new(Complex)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexConjInvolutive(t *testing.T) {
	f := func(x *Complex) bool {
		// t.Logf("x = %v", x)
		l := new(Complex)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestComplexMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *Complex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Complex), new(Complex)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Complex).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestComplexAddConjDistributive(t *testing.T) {
	f := func(x, y *Complex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Complex), new(Complex)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Complex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexSubConjDistributive(t *testing.T) {
	f := func(x, y *Complex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Complex), new(Complex)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Complex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexAddScalDistributive(t *testing.T) {
	f := func(x, y *Complex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewInt(2)
		l, r := new(Complex), new(Complex)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(Complex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexSubScalDistributive(t *testing.T) {
	f := func(x, y *Complex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewInt(2)
		l, r := new(Complex), new(Complex)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(Complex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexAddMulDistributive(t *testing.T) {
	f := func(x, y, z *Complex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Complex), new(Complex)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Complex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexSubMulDistributive(t *testing.T) {
	f := func(x, y, z *Complex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Complex), new(Complex)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Complex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestComplexQuadPositive(t *testing.T) {
	f := func(x *Complex) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestComplexComposition(t *testing.T) {
	f := func(x, y *Complex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(Complex)
		a, b := new(big.Int), new(big.Int)
		p.Mul(x, y)
		a.Set(p.Quad())
		b.Mul(x.Quad(), y.Quad())
		return a.Cmp(b) == 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
