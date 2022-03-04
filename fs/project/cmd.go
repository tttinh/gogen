package project

import (
	"fmt"
	"os"
	"text/template"

	"github.com/tinhtt/gogen/tpl"
)

func (b projectBuilder) createCmd() error {
	if err := createDir(fmt.Sprintf("%s/cmd", b.AbsolutePath)); err != nil {
		return err
	}

	rootFile, err := os.Create(fmt.Sprintf("%s/cmd/root.go", b.AbsolutePath))
	if err != nil {
		return err
	}
	defer rootFile.Close()

	rootTpl := template.Must(template.New("root").Parse(string(tpl.RootTemplate())))
	err = rootTpl.Execute(rootFile, b)
	if err != nil {
		return err
	}

	return nil
}
