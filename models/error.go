package models

import (
	"errors"
	"fmt"
)

func Error(v ...interface{}) error {
	return errors.New(fmt.Sprintln(v))
}
