package mathg_test

import (
	"math"
	"testing"

	"github.com/christopherfranklin/mathg"
)

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

func TestClampiMax(t *testing.T) {
	val := mathg.Clampi(50, 10, 20)
	if val != 20 {
		t.Fatal("Clampi max failed")
	}
}

func TestClampiMin(t *testing.T) {
	val := mathg.Clampi(5, 10, 20)
	if val != 10 {
		t.Fatal("Clampi min failed")
	}
}

func TestClampiVal(t *testing.T) {
	val := mathg.Clampi(30, 10, 50)
	if val != 30 {
		t.Fatal("Clampi value failed")
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
	val := mathg.Clampi(30.0, 10.0, 50.0)
	if val != 30.0 {
		t.Fatal("Clamp value failed")
	}
}
