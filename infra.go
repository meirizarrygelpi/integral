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

// An Infra represents an integral infra number.
type Infra struct {
	l, r big.Int
}

// Cartesian returns the two integral cartesian components of z.
func (z *Infra) Cartesian() (a, b *big.Int) {
	return &z.l, &z.r
}

// String returns the string version of a Infra value.
//
// If z corresponds to a + bε, then the string is "(a+bα)", similar to
// complex128 values.
func (z *Infra) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", z.l)
	if z.r.Sign() == -1 {
		a[2] = fmt.Sprintf("%v", z.r)
	} else {
		a[2] = fmt.Sprintf("+%v", z.r)
	}
	a[3] = "α"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Infra) Equals(y *Infra) bool {
	if z.l.Cmp(&y.l) != 0 || z.r.Cmp(&y.r) != 0 {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *Infra) Set(y *Infra) *Infra {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewInfra returns a pointer to the Infra value a+bα.
func NewInfra(a, b *big.Int) *Infra {
	z := new(Infra)
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Infra) Scal(y *Infra, a *big.Int) *Infra {
	z.l.Mul(&y.l, a)
	z.r.Mul(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Infra) Neg(y *Infra) *Infra {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Infra) Conj(y *Infra) *Infra {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z to the sum of x and y, and returns z.
func (z *Infra) Add(x, y *Infra) *Infra {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z to the difference of x and y, and returns z.
func (z *Infra) Sub(x, y *Infra) *Infra {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(α, α) = 0
// This binary operation is commutative and associative.
func (z *Infra) Mul(x, y *Infra) *Infra {
	a := new(big.Int).Set(&x.l)
	b := new(big.Int).Set(&x.r)
	c := new(big.Int).Set(&y.l)
	d := new(big.Int).Set(&y.r)
	temp := new(big.Int)
	z.l.Mul(a, c)
	z.r.Add(
		z.r.Mul(d, a),
		temp.Mul(b, c),
	)
	return z
}

// Quad returns the quadrance of z. If z = a+bα, then the quadrance is
// 		Mul(a, a)
// This is always non-negative.
func (z *Infra) Quad() *big.Int {
	return new(big.Int).Mul(&z.l, &z.l)
}

// IsZeroDiv returns true if z is a zero divisor. This is equivalent to z being
// nilpotent.
func (z *Infra) IsZeroDiv() bool {
	zero := new(big.Int)
	return z.l.Cmp(zero) == 0
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Infra) Quo(x, y *Infra) *Infra {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	quad := y.Quad()
	z.Conj(y)
	z.Mul(x, z)
	z.l.Quo(&z.l, quad)
	z.r.Quo(&z.r, quad)
	return z
}

// Generate returns a random Infra value for quick.Check testing.
func (z *Infra) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInfra := &Infra{
		*big.NewInt(rand.Int63()),
		*big.NewInt(rand.Int63()),
	}
	return reflect.ValueOf(randomInfra)
}
