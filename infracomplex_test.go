// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package integral

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestInfraComplexAddCommutative(t *testing.T) {
	f := func(x, y *InfraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(InfraComplex).Add(x, y)
		r := new(InfraComplex).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraComplexNegConjCommutative(t *testing.T) {
	f := func(x *InfraComplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(InfraComplex), new(InfraComplex)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-commutativity

func TestInfraComplexMulNonCommutative(t *testing.T) {
	f := func(x, y *InfraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(InfraComplex).Commutator(x, y)
		zero := new(InfraComplex)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestInfraComplexSubAntiCommutative(t *testing.T) {
	f := func(x, y *InfraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraComplex), new(InfraComplex)
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

func TestInfraComplexAddAssociative(t *testing.T) {
	f := func(x, y, z *InfraComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(InfraComplex), new(InfraComplex)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraComplexMulAssociative(t *testing.T) {
	f := func(x, y, z *InfraComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(InfraComplex), new(InfraComplex)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestInfraComplexAddZero(t *testing.T) {
	zero := new(InfraComplex)
	f := func(x *InfraComplex) bool {
		// t.Logf("x = %v", x)
		l := new(InfraComplex).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraComplexMulOne(t *testing.T) {
	one := &Complex{
		l: *big.NewInt(1),
	}
	zero := new(Complex)
	f := func(x *InfraComplex) bool {
		// t.Logf("x = %v", x)
		l := new(InfraComplex).Mul(x, &InfraComplex{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraComplexAddNegSub(t *testing.T) {
	f := func(x, y *InfraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraComplex), new(InfraComplex)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraComplexAddScalDouble(t *testing.T) {
	f := func(x *InfraComplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(InfraComplex), new(InfraComplex)
		l.Add(x, x)
		r.Scal(x, big.NewInt(2))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestInfraComplexNegInvolutive(t *testing.T) {
	f := func(x *InfraComplex) bool {
		// t.Logf("x = %v", x)
		l := new(InfraComplex)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraComplexConjInvolutive(t *testing.T) {
	f := func(x *InfraComplex) bool {
		// t.Logf("x = %v", x)
		l := new(InfraComplex)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestInfraComplexMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *InfraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraComplex), new(InfraComplex)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(InfraComplex).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestInfraComplexAddConjDistributive(t *testing.T) {
	f := func(x, y *InfraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraComplex), new(InfraComplex)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(InfraComplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraComplexSubConjDistributive(t *testing.T) {
	f := func(x, y *InfraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraComplex), new(InfraComplex)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(InfraComplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraComplexAddScalDistributive(t *testing.T) {
	f := func(x, y *InfraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewInt(2)
		l, r := new(InfraComplex), new(InfraComplex)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(InfraComplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraComplexSubScalDistributive(t *testing.T) {
	f := func(x, y *InfraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewInt(2)
		l, r := new(InfraComplex), new(InfraComplex)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(InfraComplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraComplexAddMulDistributive(t *testing.T) {
	f := func(x, y, z *InfraComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(InfraComplex), new(InfraComplex)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(InfraComplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraComplexSubMulDistributive(t *testing.T) {
	f := func(x, y, z *InfraComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(InfraComplex), new(InfraComplex)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(InfraComplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestInfraComplexComposition(t *testing.T) {
	f := func(x, y *InfraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(InfraComplex)
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
