package mathg

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func (v *Vec2) IsZero() bool {
	return math.Abs(v.X) < epsilon && math.Abs(v.Y) < epsilon
}

func (v *Vec2) IsEqual(v2 *Vec2) bool {
	return math.Abs(v.X-v2.X) < epsilon && math.Abs(v.Y-v2.Y) < epsilon
}

func (v *Vec2) Zero() {
	v.X = 0.0
	v.Y = 0.0
}

func (v *Vec2) One() {
	v.X = 1.0
	v.Y = 1.0
}

func (v *Vec2) Add(v2 *Vec2) *Vec2 {
	return &Vec2{v.X + v2.X, v.Y + v2.Y}
}

func (v *Vec2) AddScalar(scalar float64) *Vec2 {
	return &Vec2{v.X + scalar, v.Y + scalar}
}

func (v *Vec2) Subtract(v2 *Vec2) *Vec2 {
	return &Vec2{v.X - v2.X, v.Y - v2.Y}
}

func (v *Vec2) SubtractScalar(scalar float64) *Vec2 {
	return &Vec2{v.X - scalar, v.Y - scalar}
}

func (v *Vec2) Multiply(v2 *Vec2) *Vec2 {
	return &Vec2{v.X * v2.X, v.Y * v2.Y}
}

func (v *Vec2) MultiplyScalar(scalar float64) *Vec2 {
	return &Vec2{v.X * scalar, v.Y * scalar}
}

func (v *Vec2) MultiplyMat2(m *Mat2) *Vec2 {
	return &Vec2{m.M11*v.X + m.M12*v.Y, m.M21*v.X + m.M22*v.Y}
}

func (v *Vec2) Divide(v1 *Vec2) *Vec2 {
	return &Vec2{v.X / v1.X, v.Y / v1.Y}
}

func (v *Vec2) DivideScalar(scalar float64) *Vec2 {
	return &Vec2{v.X / scalar, v.Y / scalar}
}

func (v *Vec2) Snap(v1 *Vec2) *Vec2 {
	return &Vec2{math.Floor(v.X/v1.X) * v1.X, math.Floor(v.Y/v1.Y) * v1.Y}
}

func (v *Vec2) Snapf(f float64) *Vec2 {
	return &Vec2{math.Floor(v.X/f) * f, math.Floor(v.Y/f) * f}
}

func (v *Vec2) Negative() *Vec2 {
	return &Vec2{-1 * v.X, -1 * v.Y}
}

func (v *Vec2) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func (v *Vec2) Dot(v2 *Vec2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v *Vec2) Cross(v2 *Vec2) float64 {
	return (v.X * v2.Y) - (v.Y * v2.X)
}

func (v *Vec2) Abs() *Vec2 {
	return &Vec2{math.Abs(v.X), math.Abs(v.Y)}
}

func (v *Vec2) Floor() *Vec2 {
	return &Vec2{math.Floor(v.X), math.Floor(v.Y)}
}

func (v *Vec2) Ceil() *Vec2 {
	return &Vec2{math.Ceil(v.X), math.Ceil(v.Y)}
}

func (v *Vec2) Round() *Vec2 {
	return &Vec2{math.Round(v.X), math.Round(v.Y)}
}

func (v *Vec2) Max(v2 *Vec2) *Vec2 {
	return &Vec2{math.Max(v.X, v2.X), math.Max(v.Y, v2.Y)}
}

func (v *Vec2) Min(v2 *Vec2) *Vec2 {
	return &Vec2{math.Min(v.X, v2.X), math.Min(v.Y, v2.Y)}
}

func (v *Vec2) Normalize() *Vec2 {
	m := v.Magnitude()
	return &Vec2{v.X / m, v.Y / m}
}

func (v *Vec2) Clamp(min *Vec2, max *Vec2) *Vec2 {
	return &Vec2{Clamp(v.X, min.X, max.X), Clamp(v.Y, min.Y, max.Y)}
}

func (v *Vec2) Project(v1 *Vec2) *Vec2 {
	d := v1.Dot(v1)
	s := v.Dot(v1) / d
	return &Vec2{v1.X * s, v1.Y * s}
}

func (v *Vec2) Slide(normal *Vec2) *Vec2 {
	d := v.Dot(normal)
	return &Vec2{v.X - normal.X*d, v.Y - normal.Y*d}
}

func (v *Vec2) Reflect(normal *Vec2) *Vec2 {
	d := 2. * v.Dot(normal)
	return &Vec2{normal.X*d - v.X, normal.Y*d - v.Y}
}

func (v *Vec2) Tangent() *Vec2 {
	return &Vec2{v.Y, -1 * v.X}
}

func (v *Vec2) Rotate(radians float64) *Vec2 {
	cs := math.Cos(radians)
	sn := math.Sin(radians)
	return &Vec2{v.X*cs - v.Y*sn, v.X*sn + v.Y*cs}
}

func (v *Vec2) Lerp(v1 *Vec2, percent float64) *Vec2 {
	return &Vec2{v.X + (v1.X-v.X)*percent, v.Y + (v1.Y-v.Y)*percent}
}

func (v *Vec2) Bezier3(v1 *Vec2, v2 *Vec2, percent float64) *Vec2 {
	t0 := v.Lerp(v1, percent)
	t1 := v1.Lerp(v2, percent)
	return t0.Lerp(t1, percent)
}

func (v *Vec2) Bezier4(v1 *Vec2, v2 *Vec2, v3 *Vec2, percent float64) *Vec2 {
	t0 := v.Lerp(v1, percent)
	t1 := v1.Lerp(v2, percent)
	t2 := v2.Lerp(v3, percent)
	t3 := t0.Lerp(t1, percent)
	t4 := t1.Lerp(t2, percent)
	return t3.Lerp(t4, percent)
}

func (v *Vec2) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

func (v *Vec2) Distance(v1 *Vec2) float64 {
	return math.Sqrt(math.Pow((v.X-v1.X), 2) + math.Pow((v.Y-v1.Y), 2))
}

func (v *Vec2) LinearIndependent(v1 *Vec2) bool {
	return (v.X*v1.Y - v1.X*v.Y) != 0
}
