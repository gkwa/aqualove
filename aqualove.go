package aqualove

import (
	"errors"
	"fmt"
	"log/slog"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/manifoldco/promptui"
	mymazda "github.com/taylormonacelli/forestfish/mymazda"
)

func Main() int {
	slog.Debug("aqualove", "test", true)
	test2()

	return 0
}

func test() {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Number",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
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

func test2() error {
	var err error
	path, _ := getProjectBaseDir()
	path, err = mymazda.ExpandTilde(path)
	if err != nil {
		return err
	}

	slog.Debug("path after expansion", "path", path)

	cmd := exec.Command("flashbiter", path)
	stdout, stderr, err := runCmd(cmd)
	if err != nil || stderr != "" {
		slog.Error("runCmd", "error", err)
		return err
	}

	project_name := filepath.Base(stdout)

	url, err := getProjectTemplateURL()
	if err != nil {
		slog.Error("project template", "url", err)
		return err
	}
	slog.Debug("project template", "url", url)

	params := []string{
		"cookiecutter",
		"--no-input",
		url,
		fmt.Sprintf("project_name=%s", project_name),
		fmt.Sprintf("--output-dir=%s", path),
		"--overwrite-if-exists",
	}

	cmd = exec.Command(params[0], params[1:]...)
	slog.Debug("command", "cmd", cmd.String())

	return nil
}
