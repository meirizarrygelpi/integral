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

var symbCockle = [4]string{"", "i", "t", "u"}

// A Cockle represents an integral Cockle quaternion.
type Cockle struct {
	l, r Complex
}

// Cartesian returns the four integral Cartesian components of z.
func (z *Cockle) Cartesian() (*big.Int, *big.Int, *big.Int, *big.Int) {
	return &z.l.l, &z.l.r, &z.r.l, &z.r.r
}

// String returns the string representation of a Cockle value.
// If z corresponds to a + bi + ct + du, then the string is "(a+bi+ct+du)",
// similar to complex128 values.
func (z *Cockle) String() string {
	v := make([]*big.Int, 4)
	v[0], v[1] = z.l.Cartesian()
	v[2], v[3] = z.r.Cartesian()
	a := make([]string, 9)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", v[0])
	i := 1
	for j := 2; j < 8; j = j + 2 {
		if v[i].Sign() == -1 {
			a[j] = fmt.Sprintf("%v", v[i])
		} else {
			a[j] = fmt.Sprintf("+%v", v[i])
		}
		a[j+1] = symbCockle[i]
		i++
	}
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Cockle) Equals(y *Cockle) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *Cockle) Set(y *Cockle) *Cockle {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewCockle returns a pointer to the Cockle value a+bi+ct+du.
func NewCockle(a, b, c, d *big.Int) *Cockle {
	z := new(Cockle)
	z.l.l.Set(a)
	z.l.r.Set(b)
	z.r.l.Set(c)
	z.r.r.Set(d)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Cockle) Scal(y *Cockle, a *big.Int) *Cockle {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Cockle) Neg(y *Cockle) *Cockle {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Cockle) Conj(y *Cockle) *Cockle {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Cockle) Add(x, y *Cockle) *Cockle {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Cockle) Sub(x, y *Cockle) *Cockle {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(i, i) = -1
// 		Mul(t, t) = Mul(u, u) = +1
// 		Mul(i, t) = -Mul(t, i) = u
// 		Mul(u, t) = -Mul(t, u) = i
// 		Mul(u, i) = -Mul(i, u) = t
// This binary operation is noncommutative but associative.
func (z *Cockle) Mul(x, y *Cockle) *Cockle {
	a := new(Complex).Set(&x.l)
	b := new(Complex).Set(&x.r)
	c := new(Complex).Set(&y.l)
	d := new(Complex).Set(&y.r)
	temp := new(Complex)
	z.l.Add(
		z.l.Mul(a, c),
		temp.Mul(temp.Conj(d), b),
	)
	z.r.Add(
		z.r.Mul(d, a),
		temp.Mul(b, temp.Conj(c)),
	)
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *Cockle) Commutator(x, y *Cockle) *Cockle {
	return z.Sub(
		z.Mul(x, y),
		new(Cockle).Mul(y, x),
	)
}

// Quad returns the quadrance of z. If z = a+bi+ct+du, then the quadrance is
// 		Mul(a, a) + Mul(b, b) - Mul(c, c) - Mul(d, d)
// This can be positive, negative, or zero.
func (z *Cockle) Quad() *big.Int {
	return new(big.Int).Sub(
		z.l.Quad(),
		z.r.Quad(),
	)
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *Cockle) IsZeroDiv() bool {
	return z.l.Quad().Cmp((&z.r).Quad()) == 0
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Cockle) Quo(x, y *Cockle) *Cockle {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	quad := y.Quad()
	z.Conj(y)
	z.Mul(x, z)
	z.l.l.Quo(&z.l.l, quad)
	z.l.r.Quo(&z.l.r, quad)
	z.r.l.Quo(&z.r.l, quad)
	z.r.r.Quo(&z.r.r, quad)
	return z
}

// IsNilpotent returns true if z raised to the nth power vanishes.
func (z *Cockle) IsNilpotent(n int) bool {
	zero := new(Cockle)
	zeroInt := new(big.Int)
	if z.Equals(zero) {
		return true
	}
	p := NewCockle(big.NewInt(1), zeroInt, zeroInt, zeroInt)
	for i := 0; i < n; i++ {
		p.Mul(p, z)
		if p.Equals(zero) {
			return true
		}
	}
	return false
}

// Generate returns a random Cockle value for quick.Check testing.
func (z *Cockle) Generate(rand *rand.Rand, size int) reflect.Value {
	randomCockle := &Cockle{
		*NewComplex(
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
		),
		*NewComplex(
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
		),
	}
	return reflect.ValueOf(randomCockle)
}
