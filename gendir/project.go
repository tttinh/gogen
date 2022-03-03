package gendir

import (
	"fmt"
)

type Project struct {
	AbsolutePath    string
	PackageName     string
	ApplicationName string
}

func (p Project) Create() error {
	fmt.Printf("%+v\n", p)
	return nil
}
