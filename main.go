package mathg

import (
	"math"
)

const Pi_2 float64 = math.Pi / 2
const Pi_4 float64 = math.Pi / 4
const Epsilon float64 = float64(7.)/3 - float64(4.)/3 - float64(1.)

func ToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

func ToDegrees(radians float64) float64 {
	return radians * 180.0 / math.Pi
}

func Clamp(value float64, min float64, max float64) float64 {
	if value < min {
		value = min
	} else if value > max {
		value = max
	}
	return value
}

func Clampi(value int, min int, max int) int {
	if value < min {
		value = min
	} else if value > max {
		value = max
	}
	return value
}

func NearlyEqual(a float64, b float64, epsilon float64) bool {
	result := false
	if a == b {
		result = true
	} else if math.Abs(a-b) <= epsilon {
		result = true
	}
	return result
}
