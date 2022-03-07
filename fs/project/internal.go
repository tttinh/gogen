package project

import (
	"fmt"
)

func (b projectBuilder) createInternal() error {
	// create sub folders
	folders := [...]string{
		"adapters",
		"app",
		"domain",
		"ports",
	}
	for _, v := range folders {
		if err := createDir(fmt.Sprintf("%s/internal/%s", b.AbsolutePath, v)); err != nil {
			return err
		}
	}

	return nil
}
