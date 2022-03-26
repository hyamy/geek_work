package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main1() {
	fmt.Printf("%+v", errors.Cause(top()))
}

func top() error {
	if err := middle(); err != nil {
		return errors.Wrap(err, "top error")
	}

	return nil
}

func middle() error {
	if err := bottom(); err != nil {
		return errors.Wrap(err, "middle error")
	}

	return nil
}

func bottom() error {
	return errors.New("bottom error")
}
