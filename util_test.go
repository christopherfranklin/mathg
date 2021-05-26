package mathg_test

import (
	"math"
	"testing"

	"github.com/christopherfranklin/mathg"
)

const epsilon float64 = float64(7.)/3 - float64(4.)/3 - float64(1.)

func TestToRadians(t *testing.T) {
	deg := 180.0
	rad := mathg.ToRadians(deg)
	if rad != math.Pi {
		t.Fatalf(`ToRadians(%f) = %f`, deg, rad)
	}
}

func TestToDegrees(t *testing.T) {
	deg := mathg.ToDegrees(math.Pi)
	if deg != 180.0 {
		t.Fatalf(`ToDegrees(%f) = %f`, math.Pi, deg)
	}
}

func TestClampMax(t *testing.T) {
	val := mathg.Clamp(50.0, 10.0, 20.1)
	if val != 20.1 {
		t.Fatal("Clamp max failed")
	}
}

func TestClampMin(t *testing.T) {
	val := mathg.Clamp(5.2, 10.0, 20.5)
	if val != 10.0 {
		t.Fatal("Clamp min failed")
	}
}

func TestClampVal(t *testing.T) {
	val := mathg.Clamp(30.0, 10.0, 50.0)
	if val != 30.0 {
		t.Fatal("Clamp value failed")
	}
}

func TestEqual(t *testing.T) {
	if !mathg.NearlyEqual(0., 0., epsilon) {
		t.Fatal("Values are not equal")
	}
}

func TestNotEqual(t *testing.T) {
	if mathg.NearlyEqual(0., 0.00000001, epsilon) {
		t.Fatal("Values are equal")
	}
}
