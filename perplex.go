// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package integral

import (
	"fmt"
	"math/big"
	"math/rand"
	"reflect"
	"strings"
)

// A Perplex represents an integral perplex number.
type Perplex struct {
	l, r big.Int
}

// Cartesian returns the two cartesian components of z.
func (z *Perplex) Cartesian() (*big.Int, *big.Int) {
	return &z.l, &z.r
}

// String returns the string version of a Perplex value.
//
// If z corresponds to a + bs, then the string is "(a+bs)", similar to
// complex128 values.
func (z *Perplex) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", z.l)
	if z.r.Sign() == -1 {
		a[2] = fmt.Sprintf("%v", z.r)
	} else {
		a[2] = fmt.Sprintf("+%v", z.r)
	}
	a[3] = "s"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Perplex) Equals(y *Perplex) bool {
	if z.l.Cmp(&y.l) != 0 || z.r.Cmp(&y.r) != 0 {
		return false
	}
	return true
}

// Set copies y onto z, and returns z.
func (z *Perplex) Set(y *Perplex) *Perplex {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewPerplex returns a pointer to the Perplex value a+bs.
func NewPerplex(a, b *big.Int) *Perplex {
	z := new(Perplex)
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Perplex) Scal(y *Perplex, a *big.Int) *Perplex {
	z.l.Mul(&y.l, a)
	z.r.Mul(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Perplex) Neg(y *Perplex) *Perplex {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Perplex) Conj(y *Perplex) *Perplex {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Perplex) Add(x, y *Perplex) *Perplex {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Perplex) Sub(x, y *Perplex) *Perplex {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(s, s) = +1
// This binary operation is commutative and associative.
func (z *Perplex) Mul(x, y *Perplex) *Perplex {
	a := new(big.Int).Set(&x.l)
	b := new(big.Int).Set(&x.r)
	c := new(big.Int).Set(&y.l)
	d := new(big.Int).Set(&y.r)
	temp := new(big.Int)
	z.l.Add(
		z.l.Mul(a, c),
		temp.Mul(d, b),
	)
	z.r.Add(
		z.r.Mul(d, a),
		temp.Mul(b, c),
	)
	return z
}

// Quad returns the quadrance of z. If z = a+bs, then the quadrance is
// 		Mul(a, a) - Mul(b, b)
// This can be positive, negative, or zero.
func (z *Perplex) Quad() *big.Int {
	quad := new(big.Int)
	return quad.Sub(
		quad.Mul(&z.l, &z.l),
		new(big.Int).Mul(&z.r, &z.r),
	)
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *Perplex) IsZeroDiv() bool {
	if z.l.Cmp(&z.r) == 0 {
		return true
	}
	if z.l.Cmp(new(big.Int).Neg(&z.r)) == 0 {
		return true
	}
	return false
}

// Quo sets z equal to the quotient of x and y, and returns z. Note that
// truncated division is used.
func (z *Perplex) Quo(x, y *Perplex) *Perplex {
	if y.IsZeroDiv() {
		panic("zero divisor denominator")
	}
	quad := y.Quad()
	z.Conj(y)
	z.Mul(x, z)
	z.l.Quo(&z.l, quad)
	z.r.Quo(&z.r, quad)
	return z
}

// Generate returns a random Perplex value for quick.Check testing.
func (z *Perplex) Generate(rand *rand.Rand, size int) reflect.Value {
	randomPerplex := &Perplex{
		*big.NewInt(rand.Int63()),
		*big.NewInt(rand.Int63()),
	}
	return reflect.ValueOf(randomPerplex)
}
