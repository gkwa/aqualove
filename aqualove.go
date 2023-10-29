package aqualove

import (
	"fmt"
	"log/slog"
	"os/exec"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/taylormonacelli/flashbiter"
	mymazda "github.com/taylormonacelli/forestfish/mymazda"
)

func Main() int {
	slog.Debug("aqualove", "test", true)
	path, err := doit()
	if err != nil {
		panic(err)
	}

	fmt.Println(path)

	return 0
}

func getProjectBaseDir() (string, error) {
	prompt := promptui.Select{
		Label: "Select new project base directory",
		Items: []string{
			"~/pdev/taylormonacelli",
			"/tmp",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	fmt.Printf("You choose %q\n", result)
	return result, nil
}

func getProjectTemplateURL() (string, error) {
	prompt := promptui.Select{
		Label: "Select new project base directory",
		Items: []string{
			"https://github.com/taylormonacelli/itsvermont",
			"https://github.com/taylormonacelli/bluesorrow",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil
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
