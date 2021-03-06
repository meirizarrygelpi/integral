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

var symbSupra = [4]string{"", "α", "β", "γ"}

// A Supra represents a rational supra number.
type Supra struct {
	l, r Infra
}

// Cartesian returns the four rational Cartesian components of z.
func (z *Supra) Cartesian() (*big.Int, *big.Int, *big.Int, *big.Int) {
	return &z.l.l, &z.l.r, &z.r.l, &z.r.r
}

// String returns the string representation of a Supra value.
//
// If z corresponds to a + bα + cβ + dγ, then the string is "(a+bα+cβ+dγ)",
// similar to complex128 values.
func (z *Supra) String() string {
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
		a[j+1] = symbSupra[i]
		i++
	}
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Supra) Equals(y *Supra) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *Supra) Set(y *Supra) *Supra {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewSupra returns a pointer to the Supra value a+bα+cβ+dγ.
func NewSupra(a, b, c, d *big.Int) *Supra {
	z := new(Supra)
	z.l.l.Set(a)
	z.l.r.Set(b)
	z.r.l.Set(c)
	z.r.r.Set(d)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Supra) Scal(y *Supra, a *big.Int) *Supra {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Supra) Neg(y *Supra) *Supra {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Supra) Conj(y *Supra) *Supra {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Supra) Add(x, y *Supra) *Supra {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Supra) Sub(x, y *Supra) *Supra {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(α, α) = Mul(β, β) = Mul(γ, γ) = 0
// 		Mul(α, β) = -Mul(β, α) = γ
// 		Mul(β, γ) = Mul(γ, β) = 0
// 		Mul(γ, α) = Mul(α, γ) = 0
// This binary operation is noncommutative but associative.
func (z *Supra) Mul(x, y *Supra) *Supra {
	a := new(Infra).Set(&x.l)
	b := new(Infra).Set(&x.r)
	c := new(Infra).Set(&y.l)
	d := new(Infra).Set(&y.r)
	temp := new(Infra)
	z.l.Mul(a, c)
	z.r.Add(
		z.r.Mul(d, a),
		temp.Mul(b, temp.Conj(c)),
	)
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *Supra) Commutator(x, y *Supra) *Supra {
	return z.Sub(
		z.Mul(x, y),
		new(Supra).Mul(y, x),
	)
}

// Quad returns the quadrance of z. If z = a+bα+cβ+dγ, then the quadrance is
// 		Mul(a, a)
// This is always non-negative.
func (z *Supra) Quad() *big.Int {
	return z.l.Quad()
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *Supra) IsZeroDiv() bool {
	return z.l.IsZeroDiv()
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Supra) Quo(x, y *Supra) *Supra {
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

// Generate returns a random Supra value for quick.Check testing.
func (z *Supra) Generate(rand *rand.Rand, size int) reflect.Value {
	randomSupra := &Supra{
		*NewInfra(
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
		),
		*NewInfra(
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
		),
	}
	return reflect.ValueOf(randomSupra)
}
