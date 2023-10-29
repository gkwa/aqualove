package aqualove

import (
	"fmt"
	"io"
	"log/slog"
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
		slog.Error("error creating stderr pipe", "error", err)
		return "", "", err
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		slog.Error("error starting command", "cmd", cmd.String(), "error", err)
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
			slog.Error("command exited", "status code", exitCode)
		} else {
			// Some other error occurred
			slog.Error("error waiting for command to finish", "cmd", cmd.String(), "error", err)
		}
	} else {
		slog.Debug("command exited successfully")
	}

	stdoutStr := strings.TrimSpace(string(stdoutBytes))
	stderrStr := strings.TrimSpace(string(stderrBytes))

	return stdoutStr, stderrStr, err
}
