package mathg_test

import (
	"testing"

	"github.com/christopherfranklin/mathg"
)

func TestVec2(t *testing.T) {
	v := mathg.Vec2{1.0, 2.0}
	if v.X != 1.0 && v.Y != 2.0 {
		t.Fatal("Failed to create a Vec2")
	}
}

func TestIsEqual(t *testing.T) {
	v, w := &mathg.Vec2{}, &mathg.Vec2{}
	if v.IsEqual(w) != true {
		t.Fatal("Vec2 empty are not equal")
	}
}

func TestIsZero(t *testing.T) {
	v := &mathg.Vec2{}
	if v.IsZero() != true {
		t.Fatal("Vec2 IsZero failed")
	}
}

func TestZeroOut(t *testing.T) {
	v := mathg.Vec2{1.0, 2.0}
	v.Zero()
	if v.IsZero() != true {
		t.Fatal("Vec2 Zero failed")
	}
}
