package db

import (
	"fmt"
)

type Todo struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

func (t *Todo) Validate() (err error) {
	if len(t.Content) > 140 {
		err = fmt.Errorf("Todo must be less than or equal to 140 characters, was %d", len(t.Content))
	}
	if len(t.Content) == 0 {
		err = fmt.Errorf("Todo is empty")
	}
	return
}
