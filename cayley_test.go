// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package integral

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestCayleyAddCommutative(t *testing.T) {
	f := func(x, y *Cayley) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Cayley).Add(x, y)
		r := new(Cayley).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCayleyNegConjCommutative(t *testing.T) {
	f := func(x *Cayley) bool {
		// t.Logf("x = %v", x)
		l, r := new(Cayley), new(Cayley)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-commutativity

func TestCayleyMulNonCommutative(t *testing.T) {
	f := func(x, y *Cayley) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Cayley).Commutator(x, y)
		zero := new(Cayley)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestCayleySubAntiCommutative(t *testing.T) {
	f := func(x, y *Cayley) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Cayley), new(Cayley)
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

func TestCayleyAddAssociative(t *testing.T) {
	f := func(x, y, z *Cayley) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Cayley), new(Cayley)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-associativity

func TestCayleyMulNonAssociative(t *testing.T) {
	f := func(x, y, z *Cayley) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l := new(Cayley).Associator(x, y, z)
		zero := new(Cayley)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestCayleyAddZero(t *testing.T) {
	zero := new(Cayley)
	f := func(x *Cayley) bool {
		// t.Logf("x = %v", x)
		l := new(Cayley).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCayleyMulOne(t *testing.T) {
	one := &Hamilton{
		l: Complex{
			l: *big.NewInt(1),
		},
	}
	zero := new(Hamilton)
	f := func(x *Cayley) bool {
		// t.Logf("x = %v", x)
		l := new(Cayley).Mul(x, &Cayley{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCayleyAddNegSub(t *testing.T) {
	f := func(x, y *Cayley) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Cayley), new(Cayley)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCayleyAddScalDouble(t *testing.T) {
	f := func(x *Cayley) bool {
		// t.Logf("x = %v", x)
		l, r := new(Cayley), new(Cayley)
		l.Add(x, x)
		r.Scal(x, big.NewInt(2))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestCayleyNegInvolutive(t *testing.T) {
	f := func(x *Cayley) bool {
		// t.Logf("x = %v", x)
		l := new(Cayley)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCayleyConjInvolutive(t *testing.T) {
	f := func(x *Cayley) bool {
		// t.Logf("x = %v", x)
		l := new(Cayley)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestCayleyMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *Cayley) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Cayley), new(Cayley)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Cayley).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestCayleyAddConjDistributive(t *testing.T) {
	f := func(x, y *Cayley) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Cayley), new(Cayley)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Cayley).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCayleySubConjDistributive(t *testing.T) {
	f := func(x, y *Cayley) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Cayley), new(Cayley)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Cayley).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCayleyAddScalDistributive(t *testing.T) {
	f := func(x, y *Cayley) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewInt(2)
		l, r := new(Cayley), new(Cayley)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(Cayley).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCayleySubScalDistributive(t *testing.T) {
	f := func(x, y *Cayley) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewInt(2)
		l, r := new(Cayley), new(Cayley)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(Cayley).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCayleyAddMulDistributive(t *testing.T) {
	f := func(x, y, z *Cayley) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Cayley), new(Cayley)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Cayley).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCayleySubMulDistributive(t *testing.T) {
	f := func(x, y, z *Cayley) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Cayley), new(Cayley)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Cayley).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestCayleyQuadPositive(t *testing.T) {
	f := func(x *Cayley) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Left-alternativity

func TestCayleyLeftAlternative(t *testing.T) {
	f := func(x, y *Cayley) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Cayley)
		l.Associator(x, x, y)
		zero := new(Cayley)
		return l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Right-alternativity

func TestCayleyRightAlternative(t *testing.T) {
	f := func(x, y *Cayley) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Cayley)
		l.Associator(x, y, y)
		zero := new(Cayley)
		return l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestCayleyComposition(t *testing.T) {
	f := func(x, y *Cayley) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(Cayley)
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
