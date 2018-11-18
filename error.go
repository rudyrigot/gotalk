package gotalk

import (
	"errors"
	"fmt"
)

// Error shows an error case
func Error() error {
	i, err := twoOrFail(true)
	if err != nil {
		return err
	}
	fmt.Println(i)
	return nil
}

func twoOrFail(b bool) (int, error) {
	if b {
		return 2, nil
	} else {
		return 0, errors.New("can't give you that 2")
	}
}
