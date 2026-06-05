package model

import "fmt"

type Increment struct {
	Counter Counter
	Sum     int
}

func (i *Increment) GetIncrementUrl() string {
	return fmt.Sprintf("/increment/%d", i.Counter.Id)
}

func (i *Increment) GetIncrTarget() string {
	return fmt.Sprintf("incr-%d", i.Counter.Id)
}
