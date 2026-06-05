package model

import "fmt"

type Counter struct {
	Id    int
	Name  string
	Owner string
}

func (c *Counter) GetEditUrl() string {
	return fmt.Sprintf("/counter/edit/%d", c.Id)
}

func (c *Counter) GetDeleteUrl() string {
	return fmt.Sprintf("/counter/delete/%d", c.Id)
}

func (c *Counter) GetSaveUrl() string {
	return fmt.Sprintf("/counter/save/%d", c.Id)
}
