package mathg

import "math"

type Vec4 struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (v *Vec3) ToVec4() *Vec4 {
	return &Vec4{v.X, v.Y, v.Z, 0}
}

func (v *Vec4) IsZero() bool {
	return math.Abs(v.X) < epsilon && math.Abs(v.Y) < epsilon && math.Abs(v.Z) < epsilon && math.Abs(v.W) < epsilon
}

func (v *Vec4) IsEqual(v1 *Vec4) bool {
	return math.Abs(v.X-v1.X) < epsilon && math.Abs(v.Y-v1.Y) < epsilon && math.Abs(v.Z-v1.Z) < epsilon && math.Abs(v.W-v1.W) < epsilon
}

func (v *Vec4) Zero() {
	v.X = 0.
	v.Y = 0.
	v.Z = 0.
	v.W = 0.
}

func (v *Vec4) One() {
	v.X = 1.
	v.Y = 1.
	v.Z = 1.
	v.W = 1.
}

func (v *Vec4) Add(v1 *Vec4) *Vec4 {
	return &Vec4{v.X + v1.X, v.Y + v1.Y, v.Z + v1.Z, v.W + v1.W}
}

func (v *Vec4) AddScalar(scalar float64) *Vec4 {
	return &Vec4{v.X + scalar, v.Y + scalar, v.Z + scalar, v.W + scalar}
}

func (v *Vec4) Subtract(v1 *Vec4) *Vec4 {
	return &Vec4{v.X - v1.X, v.Y - v1.Y, v.Z - v1.Z, v.W + v1.W}
}

func (v *Vec4) SubtractScalar(scalar float64) *Vec4 {
	return &Vec4{v.X - scalar, v.Y - scalar, v.Z - scalar, v.W - scalar}
}

func (v *Vec4) Multiply(v1 *Vec4) *Vec4 {
	return &Vec4{v.X * v1.X, v.Y * v1.Y, v.Z * v1.Z, v.W * v1.W}
}

func (v *Vec4) MultiplyScalar(scalar float64) *Vec4 {
	return &Vec4{v.X * scalar, v.Y * scalar, v.Z * scalar, v.W * scalar}
}

func (v *Vec4) MultiplyMat4(m *Mat4) *Vec4 {
	return &Vec4{
		m.M11*v.X + m.M12*v.Y + m.M13*v.Z + m.M14*v.W,
		m.M21*v.X + m.M22*v.Y + m.M23*v.Z + m.M24*v.W,
		m.M31*v.X + m.M32*v.Y + m.M33*v.Z + m.M34*v.W,
		m.M41*v.X + m.M42*v.Y + m.M43*v.Z + m.M44*v.W,
	}
}

func (v *Vec4) Divide(v1 *Vec4) *Vec4 {
	return &Vec4{v.X / v1.X, v.Y / v1.Y, v.Z / v1.Z, v.W / v1.W}
}

func (v *Vec4) DivideScalar(scalar float64) *Vec4 {
	return &Vec4{v.X / scalar, v.Y / scalar, v.Z / scalar, v.W / scalar}
}

func (v *Vec4) Snap(v1 *Vec4) *Vec4 {
	return &Vec4{
		math.Floor(v.X/v1.X) * v1.X,
		math.Floor(v.Y/v1.Y) * v1.Y,
		math.Floor(v.Z/v1.Z) * v1.Z,
		math.Floor(v.W/v1.W) * v1.W,
	}
}

func (v *Vec4) Snapf(f float64) *Vec4 {
	return &Vec4{math.Floor(v.X/f) * f, math.Floor(v.Y/f) * f, math.Floor(v.Z/f) * f, math.Floor(v.W/f) * f}
}

func (v *Vec4) Negative() *Vec4 {
	return &Vec4{-1 * v.X, -1 * v.Y, -1 * v.Z, -1 * v.W}
}

func (v *Vec4) Abs() *Vec4 {
	return &Vec4{math.Abs(v.X), math.Abs(v.Y), math.Abs(v.Z), math.Abs(v.W)}
}

func (v *Vec4) Floor() *Vec4 {
	return &Vec4{math.Floor(v.X), math.Floor(v.Y), math.Floor(v.Z), math.Floor(v.W)}
}

func (v *Vec4) Ceil() *Vec4 {
	return &Vec4{math.Ceil(v.X), math.Ceil(v.Y), math.Ceil(v.Z), math.Ceil(v.W)}
}

func (v *Vec4) Round() *Vec4 {
	return &Vec4{math.Round(v.X), math.Round(v.Y), math.Round(v.Z), math.Round(v.W)}
}

func (v *Vec4) Max(v1 *Vec4) *Vec4 {
	return &Vec4{math.Max(v.X, v1.X), math.Max(v.Y, v1.Y), math.Max(v.Z, v1.Z), math.Max(v.W, v1.W)}
}

func (v *Vec4) Min(v1 *Vec4) *Vec4 {
	return &Vec4{math.Min(v.X, v1.X), math.Min(v.Y, v1.Y), math.Min(v.Z, v1.Z), math.Min(v.W, v1.W)}
}

func (v *Vec4) Clamp(min *Vec4, max *Vec4) *Vec4 {
	return &Vec4{
		Clamp(v.X, min.X, max.X),
		Clamp(v.Y, min.Y, max.Y),
		Clamp(v.Z, min.Z, max.Z),
		Clamp(v.W, min.W, max.W),
	}
}

func (v *Vec4) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2) + math.Pow(v.W, 2))
}

func (v *Vec4) Normalize() *Vec4 {
	m := v.Magnitude()
	return &Vec4{v.X / m, v.Y / m, v.Z / m, v.W / m}
}

func (v *Vec4) Lerp(v1 *Vec4, percent float64) *Vec4 {
	return &Vec4{
		v.X + (v1.X-v.X)*percent,
		v.Y + (v1.Y-v.Y)*percent,
		v.Z + (v1.Z-v.Z)*percent,
		v.W + (v1.W-v.W)*percent,
	}
}
