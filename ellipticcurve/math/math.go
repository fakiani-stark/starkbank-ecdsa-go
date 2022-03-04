package math

import (
	"math/big"
	"starkbank/ecdsa-go/ellipticcurve/point"
)

type Mather interface {
	Multiply(p point.Point, n *big.Int, N *big.Int, A *big.Int, P *big.Int) point.Point
	Add(p point.Point, q point.Point, A *big.Int, P *big.Int) point.Point
	Inv(x *big.Int, n *big.Int) *big.Int
}

type Math struct{}

//
// Fast way to multily point and scalar in elliptic curves
//
// Params:
//
// - `p`: First Point to mutiply
//
// - `n`: Scalar to mutiply
//
// - `N`: Order of the elliptic curve
//
// - `A`: Coefficient of the first-order term of the equation Y^2 = X^3 + A*X + B (mod P)
//
// - `P`: Prime number in the module of the equation Y^2 = X^3 + A*X + B (mod P)
//
// Returns:
//
// - Point that represents the multiplication of a point and a scalar
//
func Multiply(p point.Point, n *big.Int, N *big.Int, A *big.Int, P *big.Int) point.Point {
	return fromJacobian(jacobianMultiply(toJacobian(p), n, N, A, P), P)
}

//
// Fast way to add two points in elliptic curves
//
// Params:
//
// - `p`: First Point you want to add
//
// - `q`: Second Point you want to add
//
// - `A`: Coefficient of the first-order term of the equation Y^2 = X^3 + A*X + B (mod P)
//
// - `P`: Prime number in the module of the equation Y^2 = X^3 + A*X + B (mod P)
//
// Returns:
//
// - Point that represents the sum of First and Second Point
//
func Add(p point.Point, q point.Point, A *big.Int, P *big.Int) point.Point {
	return fromJacobian(jacobianAdd(toJacobian(p), toJacobian(q), A, P), P)
}

//
// Extended Euclidean Algorithm. It's the 'division' in elliptic curves
//
// Params:
//
// - `x`: Divisor
//
// - `n`: Mod for division
//
// Returns:
//
// - Value representing the division
//
func Inv(x *big.Int, n *big.Int) *big.Int {
	if x.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}
	lm := big.NewInt(1)
	hm := big.NewInt(0)
	high := n
	low := new(big.Int).Mod(x, n)
	var r, nm, nw *big.Int
	for low.Cmp(big.NewInt(1)) > 0 {
		r = new(big.Int).Div(high, low)
		nm = new(big.Int).Sub(hm, new(big.Int).Mul(lm, r))
		nw = new(big.Int).Sub(high, new(big.Int).Mul(low, r))
		high = low
		hm = lm
		low = nw
		lm = nm
	}
	return new(big.Int).Mod(lm, n)
}

//
// Convert point to Jacobian coordinates
//
// Params:
//
// - `p`: the point you want to transform
//
// Returns:
//
// - Point in Jacobian coordinates
//
func toJacobian(p point.Point) point.Point {
	return point.Point{p.X, p.Y, big.NewInt(1)}
}

//
// Convert point to Jacobian coordinates
//
// Params:
//
// - `p`: the point you want to transform
//
// - `P`: Prime number in the module of the equation Y^2 = X^3 + A*X + B (mod P)
//
// Returns:
//
// - Point in Jacobian coordinates
//
func fromJacobian(p point.Point, P *big.Int) point.Point {
	z := Inv(p.Z, P)
	x := new(big.Int).Mod(new(big.Int).Mul(p.X, new(big.Int).Exp(z, big.NewInt(2), nil)), P)
	y := new(big.Int).Mod(new(big.Int).Mul(p.Y, new(big.Int).Exp(z, big.NewInt(3), nil)), P)
	return point.Point{x, y, big.NewInt(0)}
}

//
// Double a point in elliptic curves
//
// Params:
//
// - `p`: Point you want to double
//
// - `A`: Coefficient of the first-order term of the equation Y^2 = X^3 + A*X + B (mod P)
//
// - `P`: Prime number in the module of the equation Y^2 = X^3 + A*X + B (mod P)
//
// Returns:
//
// - Point that represents the sum of First and Second Point
//
func jacobianDouble(p point.Point, A *big.Int, P *big.Int) point.Point {
	if p.Y == nil || p.Y.Cmp(big.NewInt(0)) == 0 {
		return point.Point{big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	}

	ysq := new(big.Int).Mod(new(big.Int).Exp(p.Y, big.NewInt(2), nil), P)
	S := new(big.Int).Mod(new(big.Int).Mul(new(big.Int).Mul(big.NewInt(4), p.X), ysq), P)
	M := new(big.Int).Mod(new(big.Int).Add(new(big.Int).Mul(big.NewInt(3), new(big.Int).Exp(p.X, big.NewInt(2), nil)), new(big.Int).Mul(A, new(big.Int).Exp(p.Z, big.NewInt(4), nil))), P)
	nx := new(big.Int).Mod(new(big.Int).Sub(new(big.Int).Exp(M, big.NewInt(2), nil), new(big.Int).Mul(big.NewInt(2), S)), P)
	ny := new(big.Int).Mod(new(big.Int).Sub(new(big.Int).Mul(M, new(big.Int).Sub(S, nx)), new(big.Int).Mul(big.NewInt(8), new(big.Int).Exp(ysq, big.NewInt(2), nil))), P)
	nz := new(big.Int).Mod(new(big.Int).Mul(big.NewInt(2), new(big.Int).Mul(p.Y, p.Z)), P)
	return point.Point{nx, ny, nz}
}

//
// Add two points in elliptic curves
//
// Params:
//
// - `p`: First Point you want to add
//
// - `q`: Second Point you want to add
//
// - `A`: Coefficient of the first-order term of the equation Y^2 = X^3 + A*X + B (mod P)
//
// - `P`: Prime number in the module of the equation Y^2 = X^3 + A*X + B (mod P)
//
// Returns:
//
// - Point that represents the sum of First and Second Point
//
func jacobianAdd(p point.Point, q point.Point, A *big.Int, P *big.Int) point.Point {
	if p.Y == nil {
		return q
	}
	if q.Y == nil {
		return p
	}

	U1 := new(big.Int).Mod(new(big.Int).Mul(p.X, new(big.Int).Exp(q.Z, big.NewInt(2), nil)), P)
	U2 := new(big.Int).Mod(new(big.Int).Mul(q.X, new(big.Int).Exp(p.Z, big.NewInt(2), nil)), P)
	S1 := new(big.Int).Mod(new(big.Int).Mul(p.Y, new(big.Int).Exp(q.Z, big.NewInt(3), nil)), P)
	S2 := new(big.Int).Mod(new(big.Int).Mul(q.Y, new(big.Int).Exp(p.Z, big.NewInt(3), nil)), P)

	if U1.Cmp(U2) == 0 {
		if S1.Cmp(S2) != 0 {
			return point.Point{big.NewInt(0), big.NewInt(0), big.NewInt(1)}
		}
		return jacobianDouble(p, A, P)
	}

	H := new(big.Int).Sub(U2, U1)
	R := new(big.Int).Sub(S2, S1)
	H2 := new(big.Int).Mod(new(big.Int).Mul(H, H), P)
	H3 := new(big.Int).Mod(new(big.Int).Mul(H, H2), P)
	U1H2 := new(big.Int).Mod(new(big.Int).Mul(U1, H2), P)
	nx := new(big.Int).Mod(new(big.Int).Sub(new(big.Int).Exp(R, big.NewInt(2), nil), new(big.Int).Add(H3, new(big.Int).Mul(big.NewInt(2), U1H2))), P)
	ny := new(big.Int).Mod(new(big.Int).Sub(new(big.Int).Mul(R, new(big.Int).Sub(U1H2, nx)), new(big.Int).Mul(S1, H3)), P)
	nz := new(big.Int).Mod(new(big.Int).Mul(H, new(big.Int).Mul(p.Z, q.Z)), P)
	return point.Point{nx, ny, nz}
}

//
// Multily point and scalar in elliptic curves
//
// Params:
//
// - `p`: First Point to mutiply
//
// - `n`: Scalar to mutiply
//
// - `N`: Order of the elliptic curve
//
// - `A`: Coefficient of the first-order term of the equation Y^2 = X^3 + A*X + B (mod P)
//
// - `P`: Prime number in the module of the equation Y^2 = X^3 + A*X + B (mod P)
//
// Returns:
//
// - Point that represents the sum of First and Second Point
//
func jacobianMultiply(p point.Point, n *big.Int, N *big.Int, A *big.Int, P *big.Int) point.Point {
	if p.Y.Cmp(big.NewInt(0)) == 0 || n.Cmp(big.NewInt(0)) == 0 {
		return point.Point{big.NewInt(0), big.NewInt(0), big.NewInt(1)}
	}

	if n.Cmp(big.NewInt(1)) == 0 {
		return p
	}
	if n.Cmp(big.NewInt(0)) < 0 || n.Cmp(N) >= 0 {
		return jacobianMultiply(p, big.NewInt(0).Mod(n, N), N, A, P)
	}
	if big.NewInt(0).Mod(n, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		return jacobianDouble(
			jacobianMultiply(
				p,
				new(big.Int).Div(n, big.NewInt(2)),
				N,
				A,
				P),
			A,
			P,
		)
	}

	return jacobianAdd(
		jacobianDouble(
			jacobianMultiply(
				p,
				new(big.Int).Div(n, big.NewInt(2)),
				N,
				A,
				P),
			A,
			P,
		),
		p,
		A,
		P,
	)
}
