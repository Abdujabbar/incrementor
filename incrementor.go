package incrementor

import (
	"sync"
)

//Incrementor  type struct
type Incrementor struct {
	value    int         //current stored value in incrementor
	mutex    *sync.Mutex //mutex for detect from data racing
	maxValue int         //max value for struct
}

//NewIncrementor method for creating new instance of Incrementor
func NewIncrementor() *Incrementor {
	return &Incrementor{
		maxValue: 1<<31 - 1,
		mutex:    &sync.Mutex{},
		value:    0,
	}
}

//GetNumber function for getting current value of Incrementor
func (i *Incrementor) GetNumber() int {
	defer i.mutex.Unlock()
	i.mutex.Lock()
	return i.value
}

//IncrementNumber method for incrementing the current value of Incrementor
//if incrementor value is equal or greater than max value we will set to zero current value again
func (i *Incrementor) IncrementNumber() {
	currentValue := i.GetNumber()
	defer i.mutex.Unlock()
	i.mutex.Lock()
	if currentValue+1 >= i.maxValue {
		i.value = 0
	} else {
		i.value++
	}
}

//SetMaximumValue method for setting max value on Incrementor
func (i *Incrementor) SetMaximumValue(v int) {
	i.maxValue = v
}
