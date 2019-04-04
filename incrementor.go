//Package incrementor implements simple counter structure, which supports concurrency
package incrementor

import (
	"fmt"
	"sync"
)

//Incrementor  type struct
type Incrementor struct {
	currentValue int         //current stored value in incrementor
	mutex        *sync.Mutex //mutex for detect from data racing
	maxValue     int         //max value for struct
}

//NewIncrementor method for creating new instance of Incrementor
func NewIncrementor() *Incrementor {
	return &Incrementor{
		maxValue:     1<<31 - 1,
		mutex:        &sync.Mutex{},
		currentValue: 0,
	}
}

//GetNumber function for getting current value of Incrementor
func (i *Incrementor) GetNumber() int {
	defer i.mutex.Unlock()
	i.mutex.Lock()
	return i.currentValue
}

//IncrementNumber method for incrementing the current value of Incrementor
//if incrementor value is equal or greater than max value we will set to zero current value again
func (i *Incrementor) IncrementNumber() {
	if i.GetNumber()+1 >= i.maxValue {
		i.setZeroValue()
	} else {
		defer i.mutex.Unlock()
		i.mutex.Lock()
		i.currentValue++
	}
}

//setZeroValue method for setting zero current value
func (i *Incrementor) setZeroValue() {
	defer i.mutex.Unlock()
	i.mutex.Lock()
	i.currentValue = 0
}

//SetMaximumValue method for setting max value on Incrementor
func (i *Incrementor) SetMaximumValue(v int) error {
	if v >= 0 {
		if i.GetNumber() >= v {
			i.setZeroValue()
		}
		i.maxValue = v
		return nil
	}
	return fmt.Errorf("You cannot set maximum value less then zero")
}
