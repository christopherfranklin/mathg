package mathg

import "math"

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v *Vec2) ToVec3() *Vec3 {
	return &Vec3{v.X, v.Y, 0}
}

func (v *Vec3) IsZero() bool {
	return math.Abs(v.X) < epsilon && math.Abs(v.Y) < epsilon && math.Abs(v.Z) < epsilon
}

func (v *Vec3) IsEqual(v1 *Vec3) bool {
	return math.Abs(v.X-v1.X) < epsilon && math.Abs(v.Y-v1.Y) < epsilon && math.Abs(v.Z-v1.Z) < epsilon
}

func (v *Vec3) Zero() {
	v.X = 0.
	v.Y = 0.
	v.Z = 0.
}

func (v *Vec3) One() {
	v.X = 1.
	v.Y = 1.
	v.Z = 1.
}

func (v *Vec3) Add(v1 *Vec3) *Vec3 {
	return &Vec3{v.X + v1.X, v.Y + v1.Y, v.Z + v1.Z}
}

func (v *Vec3) AddScalar(scalar float64) *Vec3 {
	return &Vec3{v.X + scalar, v.Y + scalar, v.Z + scalar}
}

func (v *Vec3) Subtract(v1 *Vec3) *Vec3 {
	return &Vec3{v.X - v1.X, v.Y - v1.Y, v.Z - v1.Z}
}

func (v *Vec3) SubtractScalar(scalar float64) *Vec3 {
	return &Vec3{v.X - scalar, v.Y - scalar, v.Z - scalar}
}

func (v *Vec3) Multiply(v1 *Vec3) *Vec3 {
	return &Vec3{v.X * v1.X, v.Y * v1.Y, v.Z * v1.Z}
}

func (v *Vec3) MultiplyScalar(scalar float64) *Vec3 {
	return &Vec3{v.X * scalar, v.Y * scalar, v.Z * scalar}
}

func (v *Vec3) MultiplyMat3(m *Mat3) *Vec3 {
	return &Vec3{m.M11*v.X + m.M12*v.Y + m.M13*v.Z, m.M21*v.X + m.M22*v.Y + m.M23*v.Z, m.M31*v.X + m.M32*v.Y + m.M33*v.Z}
}

func (v *Vec3) Divide(v1 *Vec3) *Vec3 {
	return &Vec3{v.X / v1.X, v.Y / v1.Y, v.Z / v1.Z}
}

func (v *Vec3) DivideScalar(scalar float64) *Vec3 {
	return &Vec3{v.X / scalar, v.Y / scalar, v.Z / scalar}
}

func (v *Vec3) Snap(v1 *Vec3) *Vec3 {
	return &Vec3{math.Floor(v.X/v1.X) * v1.X, math.Floor(v.Y/v1.Y) * v1.Y, math.Floor(v.Z/v1.Z) * v1.Z}
}

func (v *Vec3) Snapf(f float64) *Vec3 {
	return &Vec3{math.Floor(v.X/f) * f, math.Floor(v.Y/f) * f, math.Floor(v.Z/f) * f}
}

func (v *Vec3) Negative() *Vec3 {
	return &Vec3{-1 * v.X, -1 * v.Y, -1 * v.Z}
}

func (v *Vec3) Abs() *Vec3 {
	return &Vec3{math.Abs(v.X), math.Abs(v.Y), math.Abs(v.Z)}
}

func (v *Vec3) Floor() *Vec3 {
	return &Vec3{math.Floor(v.X), math.Floor(v.Y), math.Floor(v.Z)}
}

func (v *Vec3) Ceil() *Vec3 {
	return &Vec3{math.Ceil(v.X), math.Ceil(v.Y), math.Ceil(v.Z)}
}

func (v *Vec3) Round() *Vec3 {
	return &Vec3{math.Round(v.X), math.Round(v.Y), math.Round(v.Z)}
}

func (v *Vec3) Max(v1 *Vec3) *Vec3 {
	return &Vec3{math.Max(v.X, v1.X), math.Max(v.Y, v1.Y), math.Max(v.Z, v1.Z)}
}

func (v *Vec3) Min(v1 *Vec3) *Vec3 {
	return &Vec3{math.Min(v.X, v1.X), math.Min(v.Y, v1.Y), math.Min(v.Z, v1.Z)}
}

func (v *Vec3) Clamp(min *Vec3, max *Vec3) *Vec3 {
	return &Vec3{Clamp(v.X, min.X, max.X), Clamp(v.Y, min.Y, max.Y), Clamp(v.Z, min.Z, max.Z)}
}

func (v *Vec3) Cross(v1 *Vec3) *Vec3 {
	return &Vec3{v.Y*v1.Z - v.Z*v1.Y, v.Z*v1.X - v.X*v1.Z, v.X*v1.Y - v.Y*v1.X}
}

func (v *Vec3) Dot(v1 *Vec3) float64 {
	return v.X*v1.X + v.Y*v1.Y + v.X*v1.Z
}

func (v *Vec3) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2))
}

func (v *Vec3) LengthSquared() float64 {
	return math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2)
}

func (v *Vec3) Normalize() *Vec3 {
	m := v.Magnitude()
	return &Vec3{v.X / m, v.Y / m, v.Z / m}
}

func (v *Vec3) Project(v1 *Vec3) *Vec3 {
	d := v1.Dot(v1)
	s := v.Dot(v1) / d
	return &Vec3{v.X * s, v.Y * s, v.Z * s}
}

func (v *Vec3) Slide(normal *Vec3) *Vec3 {
	d := v.Dot(normal)
	return &Vec3{v.X - normal.X*d, v.Y - normal.Y*d, v.Z - normal.Z*d}
}

func (v *Vec3) Reflect(normal *Vec3) *Vec3 {
	d := 2. * v.Dot(normal)
	return &Vec3{normal.X*d - v.X, normal.Y*d - v.Y, normal.Z*d - v.Z}
}

func (v *Vec3) Rotate(ra *Vec3, radians float64) *Vec3 {
	cs := math.Cos(radians)
	sn := math.Sin(radians)
	x, y, z := v.X, v.Y, v.Z
	rn := ra.Normalize()
	rx, ry, rz := rn.X, rn.Y, rn.Z
	rfx := x*(cs+rx*rx*(1-cs)) + y*(rx*ry*(1-cs)-rz*sn) + z*(rx*rz*(1-cs)+ry*sn)
	rfy := x*(ry*rx*(1-cs)+rz*sn) + y*(cs+ry*ry*(1-cs)) + z*(ry*rz*(1-cs)-rx*sn)
	rfz := x*(rz*rx*(1-cs)-ry*sn) + y*(rz*ry*(1-cs)+rx*sn) + z*(cs+rz*rz*(1-cs))
	return &Vec3{rfx, rfy, rfz}
}

func (v *Vec3) Lerp(v1 *Vec3, percent float64) *Vec3 {
	return &Vec3{v.X + (v1.X-v.X)*percent, v.Y + (v1.Y-v.Y)*percent, v.Z + (v1.Z-v.Z)*percent}
}

func (v *Vec3) Bezier3(v1 *Vec3, v2 *Vec3, percent float64) *Vec3 {
	t0 := v.Lerp(v1, percent)
	t1 := v1.Lerp(v2, percent)
	return t0.Lerp(t1, percent)
}

func (v *Vec3) Bezier4(v1 *Vec3, v2 *Vec3, v3 *Vec3, percent float64) *Vec3 {
	t0 := v.Lerp(v1, percent)
	t1 := v1.Lerp(v2, percent)
	t2 := v2.Lerp(v3, percent)
	t3 := t0.Lerp(t1, percent)
	t4 := t1.Lerp(t2, percent)
	return t3.Lerp(t4, percent)
}

func (v *Vec3) Distance(v1 *Vec3) float64 {
	return math.Sqrt(math.Pow(v.X-v1.X, 2) + math.Pow(v.Y-v1.Y, 2) + math.Pow(v.Z-v1.Z, 2))
}

func (v *Vec3) LinearIndependent(v1 *Vec3, v2 *Vec3) bool {
	return (v.X*v1.Y*v2.Z + v.Y*v1.Z*v2.X + v.Z*v1.X*v2.Y - v.Z*v1.Y*v2.X - v.Y*v1.Z*v2.X - v.X*v1.Y*v2.Z) != 0
}
