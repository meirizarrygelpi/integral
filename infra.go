// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package integral

import (
	"fmt"
	"math/big"
	"strings"
)

// An Infra represents an integral infra number (also known as an integral dual
// number).
type Infra struct {
	l, r *big.Int
}

// L returns the left Cayley-Dickson component of z.
func (z *Infra) L() *big.Int {
	return z.l
}

// R returns the right Cayley-Dickson component of z.
func (z *Infra) R() *big.Int {
	return z.r
}

// SetL sets the left Cayley-Dickson component of z equal to a.
func (z *Infra) SetL(a *big.Int) {
	z.l = a
}

// SetR sets the right Cayley-Dickson component of z equal to b.
func (z *Infra) SetR(b *big.Int) {
	z.r = b
}

// Cartesian returns the two cartesian components of z.
func (z *Infra) Cartesian() (a, b *big.Int) {
	a = z.L()
	b = z.R()
	return
}

// String returns the string version of a Infra value.
//
// If z corresponds to a + bε, then the string is "(a+bε)", similar to
// complex128 values.
func (z *Infra) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", z.L())
	if z.R().Sign() == -1 {
		a[2] = fmt.Sprintf("%v", z.R())
	} else {
		a[2] = fmt.Sprintf("+%v", z.R())
	}
	a[3] = "ε"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Infra) Equals(y *Infra) bool {
	if z.L().Cmp(y.L()) != 0 || z.R().Cmp(y.R()) != 0 {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Infra) Copy(y *Infra) *Infra {
	z.SetL(y.L())
	z.SetR(y.R())
	return z
}

// NewInfra returns a pointer to a Infra value made from two given pointers to
// big.Int values.
func NewInfra(a, b *big.Int) *Infra {
	z := new(Infra)
	z.SetL(a)
	z.SetR(b)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Infra) Scal(y *Infra, a *big.Int) *Infra {
	z.SetL(new(big.Int).Mul(y.L(), a))
	z.SetR(new(big.Int).Mul(y.R(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Infra) Neg(y *Infra) *Infra {
	z.SetL(new(big.Int).Neg(y.L()))
	z.SetR(new(big.Int).Neg(y.R()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Infra) Conj(y *Infra) *Infra {
	z.SetL(y.L())
	z.SetR(new(big.Int).Neg(y.R()))
	return z
}

// Add sets z to the sum of x and y, and returns z.
func (z *Infra) Add(x, y *Infra) *Infra {
	z.SetL(new(big.Int).Add(x.L(), y.L()))
	z.SetR(new(big.Int).Add(x.R(), y.R()))
	return z
}

// Sub sets z to the difference of x and y, and returns z.
func (z *Infra) Sub(x, y *Infra) *Infra {
	z.SetL(new(big.Int).Sub(x.L(), y.L()))
	z.SetR(new(big.Int).Sub(x.R(), y.R()))
	return z
}

// Mul sets z to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(ε, ε) = 0
// This binary operation is commutative and associative.
func (z *Infra) Mul(x, y *Infra) *Infra {
	p := new(Infra).Copy(x)
	q := new(Infra).Copy(y)
	z.SetL(
		new(big.Int).Mul(p.L(), q.L()),
	)
	z.SetR(new(big.Int).Add(
		new(big.Int).Mul(q.R(), p.L()),
		new(big.Int).Mul(p.R(), q.L()),
	))
	return z
}

// Quad returns the quadrance of z, a pointer to a big.Int value.
func (z *Infra) Quad() *big.Int {
	return new(big.Int).Mul(z.L(), z.L())
}

// IsZeroDiv returns true if z is a zero divisor. This is equivalent to z being
// nilpotent.
func (z *Infra) IsZeroDiv() bool {
	return z.L().Cmp(big.NewInt(0)) == 0
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Infra) Quo(x, y *Infra) *Infra {
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
