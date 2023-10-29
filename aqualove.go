package aqualove

import (
	"fmt"
	"log/slog"
	"os/exec"
	"path/filepath"

	"github.com/taylormonacelli/aeryavenue"
	"github.com/taylormonacelli/flashbiter"
	mymazda "github.com/taylormonacelli/forestfish/mymazda"
)

func Main() int {
	slog.Debug("aqualove", "test", true)
	projectPath, err := doit()
	if err != nil {
		panic(err)
	}

	fmt.Println(projectPath)

	return 0
}

func doit() (string, error) {
	var err error
	path, _ := getProjectBaseDir()
	path, err = mymazda.ExpandTilde(path)
	if err != nil {
		return "", err
	}

	slog.Debug("path after expansion", "path", path)

	projectPath, err := flashbiter.GetUniquePath()
	if err != nil {
		panic(err)
	}

	project_name := filepath.Base(projectPath)

	url, err := getProjectTemplateURL()
	if err != nil {
		slog.Error("project template", "url", err)
		return "", err
	}
	slog.Debug("project template", "url", url)

	params := []string{
		"cookiecutter",
		"--no-input",
		url,
		fmt.Sprintf("project_name=%s", project_name),
		fmt.Sprintf("--output-dir=%s", path),
	}

	cmd := exec.Command(params[0], params[1:]...)
	slog.Debug("command", "cmd", cmd.String())
	stdout, stderr, err := runCmd(cmd)
	slog.Debug("runCmd", "stdout", stdout)
	slog.Debug("runCmd", "stderr", stderr)

	if err != nil {
		panic(err)
	}

	path = filepath.Join(path, projectPath)
	return path, nil
}

func getProjectBaseDir() (string, error) {
	baseDirOptions := map[string]string{
		"~/pdev/taylormonacelli": "~/pdev/taylormonacelli",
		"/tmp":                   "/tmp",
	}
	inputSelector := aeryavenue.GetInputSelector()
	baseDir, err := aeryavenue.SelectItem(baseDirOptions, inputSelector)
	if err != nil {
		slog.Error("selectItem failed", "error", err)
	}

	return baseDir, nil
}

func getProjectTemplateURL() (string, error) {
	baseDirOptions := map[string]string{
		"https://github.com/taylormonacelli/itsvermont/archive/refs/heads/master.zip": "https://github.com/taylormonacelli/itsvermont/archive/refs/heads/master.zip",
		"https://github.com/taylormonacelli/bluesorrow/archive/refs/heads/master.zip": "https://github.com/taylormonacelli/bluesorrow/archive/refs/heads/master.zip",
	}
	inputSelector := aeryavenue.GetInputSelector()
	template, err := aeryavenue.SelectItem(baseDirOptions, inputSelector)
	if err != nil {
		slog.Error("selectItem failed", "error", err)
	}

	return template, nil
}
