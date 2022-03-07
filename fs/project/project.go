package project

import (
	"fmt"
	"os"
	"text/template"

	"github.com/tinhtt/gogen/tpl"
)

type Options struct {
	AbsolutePath    string
	PackageName     string
	ApplicationName string
}

func Create(opts Options) error {
	builder := projectBuilder{opts}
	return builder.build()
}

type projectBuilder struct {
	Options
}

func (b projectBuilder) build() error {
	// root folder
	if err := createDir(b.AbsolutePath); err != nil {
		return err
	}

	// create sub folders
	folders := [...]string{
		"api",
		"cmd",
		"docs",
		"internal",
	}
	for _, v := range folders {
		if err := createDir(fmt.Sprintf("%s/%s", b.AbsolutePath, v)); err != nil {
			return err
		}
	}

	if err := b.createMainFile(); err != nil {
		return err
	}

	if err := b.createDocs(); err != nil {
		return err
	}

	if err := b.createCmd(); err != nil {
		return err
	}

	if err := b.createInternal(); err != nil {
		return err
	}

	return nil
}

func (b projectBuilder) createMainFile() error {
	mainFile, err := os.Create(fmt.Sprintf("%s/main.go", b.AbsolutePath))
	if err != nil {
		return err
	}
	defer mainFile.Close()

	mainTpl := template.Must(template.New("main").Parse(string(tpl.MainTemplate())))
	err = mainTpl.Execute(mainFile, b)
	if err != nil {
		return err
	}

	return nil
}

func (b projectBuilder) createDocs() error {
	f, err := os.Create(fmt.Sprintf("%s/docs/ARCHITECTURE.md", b.AbsolutePath))

	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(tpl.ArchitectureTemplate())
	if err != nil {
		return err
	}

	return nil
}
