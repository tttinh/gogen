package fs

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/tinhtt/gogen/fs/project"
)

func Generate() (string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	moduleName, err := getModuleImportPath()
	if err != nil {
		return "", err
	}

	projectOpts := project.Options{
		AbsolutePath:    workingDir,
		PackageName:     moduleName,
		ApplicationName: path.Base(moduleName),
	}

	if err := project.Create(projectOpts); err != nil {
		return "", err
	}

	return workingDir, nil
}

func getModuleImportPath() (string, error) {
	mod, err := parseModuleInfo()
	if err != nil {
		return "", err
	}

	pkg, err := parsePackageInfo()
	if err != nil {
		return "", err
	}

	return path.Join(mod.Path, fileToURL(strings.TrimPrefix(pkg.Dir, mod.Dir))), nil
}

func fileToURL(in string) string {
	i := strings.Split(in, string(filepath.Separator))
	return path.Join(i...)
}

func parseModuleInfo() (*moduleInfo, error) {
	var mod moduleInfo

	info, err := getModuleData("-m")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(info, &mod); err != nil {
		return nil, err
	}

	if mod.Path == "command-line-arguments" {
		return nil, errors.New("please run `go mod init <MODULENAME>` before `gogen init`")
	}

	return &mod, nil

}

func parsePackageInfo() (*packageInfo, error) {
	var pkg packageInfo

	info, err := getModuleData("-e")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(info, &pkg); err != nil {
		return nil, err
	}

	return &pkg, nil
}

func getModuleData(args ...string) ([]byte, error) {
	cmdArgs := append([]string{"list", "-json"}, args...)
	return exec.Command("go", cmdArgs...).Output()
}

// moduleInfo: the output of `go list` when using -m flag, omitted non-relevent fields.
type moduleInfo struct {
	Path, Dir, GoMod string
}

// packageInfo: `go list` output when using -e flag, , omitted non-relevent fields.
type packageInfo struct {
	Dir string
}
