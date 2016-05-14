// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package integral

import (
	"fmt"
	"math/big"
	"strings"
)

// A Perplex represents an integral perplex number.
type Perplex struct {
	l, r *big.Int
}

// L returns the left Cayley-Dickson component of z.
func (z *Perplex) L() *big.Int {
	return z.l
}

// R returns the right Cayley-Dickson component of z.
func (z *Perplex) R() *big.Int {
	return z.r
}

// SetL sets the left Cayley-Dickson component of z equal to a.
func (z *Perplex) SetL(a *big.Int) {
	z.l = a
}

// SetR sets the right Cayley-Dickson component of z equal to b.
func (z *Perplex) SetR(b *big.Int) {
	z.r = b
}

// Cartesian returns the two cartesian components of z.
func (z *Perplex) Cartesian() (a, b *big.Int) {
	a = z.L()
	b = z.R()
	return
}

// String returns the string version of a Perplex value.
//
// If z corresponds to a + bs, then the string is "(a+bs)", similar to
// complex128 values.
func (z *Perplex) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", z.L())
	if z.R().Sign() == -1 {
		a[2] = fmt.Sprintf("%v", z.R())
	} else {
		a[2] = fmt.Sprintf("+%v", z.R())
	}
	a[3] = "s"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Perplex) Equals(y *Perplex) bool {
	if z.L().Cmp(y.L()) != 0 || z.R().Cmp(y.R()) != 0 {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Perplex) Copy(y *Perplex) *Perplex {
	z.SetL(y.L())
	z.SetR(y.R())
	return z
}

// NewPerplex returns a pointer to a Perplex value made from two given pointers
// to big.Int values.
func NewPerplex(a, b *big.Int) *Perplex {
	z := new(Perplex)
	z.SetL(a)
	z.SetR(b)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Perplex) Scal(y *Perplex, a *big.Int) *Perplex {
	z.SetL(new(big.Int).Mul(y.L(), a))
	z.SetR(new(big.Int).Mul(y.R(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Perplex) Neg(y *Perplex) *Perplex {
	z.SetL(new(big.Int).Neg(y.L()))
	z.SetR(new(big.Int).Neg(y.R()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Perplex) Conj(y *Perplex) *Perplex {
	z.SetL(y.L())
	z.SetR(new(big.Int).Neg(y.R()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Perplex) Add(x, y *Perplex) *Perplex {
	z.SetL(new(big.Int).Add(x.L(), y.L()))
	z.SetR(new(big.Int).Add(x.R(), y.R()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Perplex) Sub(x, y *Perplex) *Perplex {
	z.SetL(new(big.Int).Sub(x.L(), y.L()))
	z.SetR(new(big.Int).Sub(x.R(), y.R()))
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(s, s) = +1
// This binary operation is commutative and associative.
func (z *Perplex) Mul(x, y *Perplex) *Perplex {
	p := new(Perplex).Copy(x)
	q := new(Perplex).Copy(y)
	z.SetL(new(big.Int).Add(
		new(big.Int).Mul(p.L(), q.L()),
		new(big.Int).Mul(q.R(), p.R()),
	))
	z.SetR(new(big.Int).Add(
		new(big.Int).Mul(q.R(), p.L()),
		new(big.Int).Mul(p.R(), q.L()),
	))
	return z
}

// Quad returns the quadrance of z, a pointer to a big.Int value.
func (z *Perplex) Quad() *big.Int {
	return new(big.Int).Sub(
		new(big.Int).Mul(z.L(), z.L()),
		new(big.Int).Mul(z.R(), z.R()),
	)
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *Perplex) IsZeroDiv() bool {
	if z.L().Cmp(z.R()) == 0 {
		return true
	}
	if z.L().Cmp(new(big.Int).Neg(z.R())) == 0 {
		return true
	}
	return false
}

// Quo sets z equal to the quotient of x and y, and returns z. Note that
// truncated division is used.
func (z *Perplex) Quo(x, y *Perplex) *Perplex {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	quad := y.Quad()
	z.Conj(y)
	z.Mul(x, z)
	z.SetL(new(big.Int).Quo(z.L(), quad))
	z.SetR(new(big.Int).Quo(z.R(), quad))
	return z
}
