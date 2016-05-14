// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package integral

import (
	"fmt"
	"math/big"
	"strings"
)

// A Complex represents an integral complex number (also known as a Gaussian
// integer).
type Complex struct {
	l, r *big.Int
}

// L returns the left Cayley-Dickson component of z.
func (z *Complex) L() *big.Int {
	return z.l
}

// R returns the right Cayley-Dickson component of z.
func (z *Complex) R() *big.Int {
	return z.r
}

// SetL sets the left Cayley-Dickson component of z equal to a.
func (z *Complex) SetL(a *big.Int) {
	z.l = a
}

// SetR sets the right Cayley-Dickson component of z equal to b.
func (z *Complex) SetR(b *big.Int) {
	z.r = b
}

// Cartesian returns the two cartesian components of z.
func (z *Complex) Cartesian() (a, b *big.Int) {
	a = z.L()
	b = z.R()
	return
}

// String returns the string version of a Complex value.
//
// If z corresponds to a + bi, then the string is "(a+bi)", similar to
// complex128 values.
func (z *Complex) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", z.L())
	if z.R().Sign() == -1 {
		a[2] = fmt.Sprintf("%v", z.R())
	} else {
		a[2] = fmt.Sprintf("+%v", z.R())
	}
	a[3] = "i"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Complex) Equals(y *Complex) bool {
	if z.L().Cmp(y.L()) != 0 || z.R().Cmp(y.R()) != 0 {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Complex) Copy(y *Complex) *Complex {
	z.SetL(y.L())
	z.SetR(y.R())
	return z
}

// NewComplex returns a pointer to a Complex value made from two given pointers
// to big.Int values.
func NewComplex(a, b *big.Int) *Complex {
	z := new(Complex)
	z.SetL(a)
	z.SetR(b)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Complex) Scal(y *Complex, a *big.Int) *Complex {
	z.SetL(new(big.Int).Mul(y.L(), a))
	z.SetR(new(big.Int).Mul(y.R(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Complex) Neg(y *Complex) *Complex {
	z.SetL(new(big.Int).Neg(y.L()))
	z.SetR(new(big.Int).Neg(y.R()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Complex) Conj(y *Complex) *Complex {
	z.SetL(y.L())
	z.SetR(new(big.Int).Neg(y.R()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Complex) Add(x, y *Complex) *Complex {
	z.SetL(new(big.Int).Add(x.L(), y.L()))
	z.SetR(new(big.Int).Add(x.R(), y.R()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Complex) Sub(x, y *Complex) *Complex {
	z.SetL(new(big.Int).Sub(x.L(), y.L()))
	z.SetR(new(big.Int).Sub(x.R(), y.R()))
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(i, i) = -1
// This binary operation is commutative and associative.
func (z *Complex) Mul(x, y *Complex) *Complex {
	p := new(Complex).Copy(x)
	q := new(Complex).Copy(y)
	z.SetL(new(big.Int).Sub(
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
func (z *Complex) Quad() *big.Int {
	return new(big.Int).Add(
		new(big.Int).Mul(z.L(), z.L()),
		new(big.Int).Mul(z.R(), z.R()),
	)
}

// Quo sets z equal to the quotient of x and y, and returns z. Note that
// truncated division is used.
func (z *Complex) Quo(x, y *Complex) *Complex {
	quad := y.Quad()
	z.Conj(y)
	z.Mul(x, z)
	z.SetL(new(big.Int).Quo(z.L(), quad))
	z.SetR(new(big.Int).Quo(z.R(), quad))
	return z
}
