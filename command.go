package aqualove

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func runCmd(cmd *exec.Cmd) (string, string, error) {
	// cmd := exec.Command("ls", "-l")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating stdout pipe:", err)
		return "", "", err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("Error creating stderr pipe:", err)
		return "", "", err
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting the command:", err)
		return "", "", err
	}

	// Read the output of the command from the pipes
	// You can use a goroutine to read from stdout and stderr concurrently if needed.
	stdoutBytes, _ := io.ReadAll(stdout)
	stderrBytes, _ := io.ReadAll(stderr)

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode := exitError.ExitCode()
			fmt.Printf("Command exited with status code %d\n", exitCode)
		} else {
			// Some other error occurred
			fmt.Println("Error waiting for the command to finish:", err)
		}
	} else {
		fmt.Println("Command exited successfully")
	}

	stdoutStr := strings.TrimSpace(string(stdoutBytes))
	stderrStr := strings.TrimSpace(string(stderrBytes))

	return stdoutStr, stderrStr, err
}
