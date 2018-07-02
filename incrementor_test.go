package incrementor

import "testing"

func TestIncrement(t *testing.T) {
	n := NewIncrementor()
	if n.GetNumber() != 0 {
		t.Error("Error while initilize incrementor value")
	}
	n.IncrementNumber()
	n.IncrementNumber()
	if n.GetNumber() != 2 {
		t.Error("Error on increment value")
	}
}

func TestMaxValue(t *testing.T) {
	n := NewIncrementor()
	n.IncrementNumber()
	n.IncrementNumber()
	n.SetMaximumValue(3)
	n.IncrementNumber()
	if n.GetNumber() != 0 {
		t.Error("Error while cleaning up value")
	}
}

func TestSetMaxValueGreaterThanCurrentValue(t *testing.T) {
	n := NewIncrementor()
	n.IncrementNumber()
	n.IncrementNumber()
	n.IncrementNumber()
	n.SetMaximumValue(2)
	if n.GetNumber() != 0 {
		t.Error("Error while cleaning up the value on set lower max value")
	}
}
