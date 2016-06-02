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

// A Complex represents an integral complex number.
type Complex struct {
	l, r big.Int
}

// Cartesian returns the two integral cartesian components of z.
func (z *Complex) Cartesian() (*big.Int, *big.Int) {
	return &z.l, &z.r
}

// String returns the string version of a Complex value.
//
// If z corresponds to a + bi, then the string is "(a+bi)", similar to
// complex128 values.
func (z *Complex) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", z.l)
	if z.r.Sign() == -1 {
		a[2] = fmt.Sprintf("%v", z.r)
	} else {
		a[2] = fmt.Sprintf("+%v", z.r)
	}
	a[3] = "i"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Complex) Equals(y *Complex) bool {
	if z.l.Cmp(&y.l) != 0 || z.r.Cmp(&y.r) != 0 {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *Complex) Set(y *Complex) *Complex {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewComplex returns a pointer to the Complex value a+bi.
func NewComplex(a, b *big.Int) *Complex {
	z := new(Complex)
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Complex) Scal(y *Complex, a *big.Int) *Complex {
	z.l.Mul(&y.l, a)
	z.r.Mul(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Complex) Neg(y *Complex) *Complex {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Complex) Conj(y *Complex) *Complex {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Complex) Add(x, y *Complex) *Complex {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Complex) Sub(x, y *Complex) *Complex {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(i, i) = -1
// This binary operation is commutative and associative.
func (z *Complex) Mul(x, y *Complex) *Complex {
	a := new(big.Int).Set(&x.l)
	b := new(big.Int).Set(&x.r)
	c := new(big.Int).Set(&y.l)
	d := new(big.Int).Set(&y.r)
	temp := new(big.Int)
	z.l.Sub(
		z.l.Mul(a, c),
		temp.Mul(d, b),
	)
	z.r.Add(
		z.r.Mul(d, a),
		temp.Mul(b, c),
	)
	return z
}

// Quad returns the quadrance of z. If z = a+bi, then the quadrance is
// 		Mul(a, a) + Mul(b, b)
// This is always non-negative.
func (z *Complex) Quad() *big.Int {
	quad := new(big.Int)
	return quad.Add(
		quad.Mul(&z.l, &z.l),
		new(big.Int).Mul(&z.r, &z.r),
	)
}

// Quo sets z equal to the quotient of x and y, and returns z. Note that
// truncated division is used.
func (z *Complex) Quo(x, y *Complex) *Complex {
	zero := new(Complex)
	if y.Equals(zero) {
		panic("zero denominator")
	}
	quad := y.Quad()
	z.Conj(y)
	z.Mul(x, z)
	z.l.Quo(&z.l, quad)
	z.r.Quo(&z.r, quad)
	return z
}

// Generate returns a random Complex value for quick.Check testing.
func (z *Complex) Generate(rand *rand.Rand, size int) reflect.Value {
	randomComplex := &Complex{
		*big.NewInt(rand.Int63()),
		*big.NewInt(rand.Int63()),
	}
	return reflect.ValueOf(randomComplex)
}
