package project

import (
	"fmt"
)

func (b projectBuilder) createInternal() error {
	// create sub folders
	folders := [...]string{
		"adapter",
		"app",
		"domain",
		"port",
	}
	for _, v := range folders {
		if err := createDir(fmt.Sprintf("%s/internal/%s", b.AbsolutePath, v)); err != nil {
			return err
		}
	}

	return nil
}
