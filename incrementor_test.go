package incrementor

import (
	"sync"
	"testing"
)

func TestIncrement(t *testing.T) {
	n := NewIncrementor()

	if n.GetMaxValue() != 1<<31-1 {
		t.Errorf("Error while initilize max value for counter")
	}

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
	err := n.SetMaximumValue(3)
	if err != nil {
		t.Error(err.Error())
	}
	n.IncrementNumber()
	if n.GetNumber() != 0 {
		t.Error("Error while cleaning up value")
	}
}

func TestMaxValueWithLessThanZero(t *testing.T) {
	n := NewIncrementor()
	err := n.SetMaximumValue(-2)
	if err == nil {
		t.Error("Problem on receive less than zero for maximum value")
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

func TestConcurrencyIncrementing(t *testing.T) {
	inc := NewIncrementor()
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inc.IncrementNumber()
		}()
	}
	wg.Wait()
	if inc.GetNumber() != 1000 {
		t.Errorf("Error while on concurrency incrementing, with answer %d", inc.GetNumber())
	}
}

func TestConcurrencyIncrementingAndMaxValue(t *testing.T) {
	inc := NewIncrementor()
	inc.SetMaximumValue(2)
	wg := &sync.WaitGroup{}
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inc.IncrementNumber()
		}()
	}
	wg.Wait()
	if inc.GetNumber() != 0 {
		t.Errorf("error while on concurrency incrementing and max value, with answer: %d", inc.GetNumber())
	}
}
