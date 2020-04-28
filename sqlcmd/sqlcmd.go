package sqlcmd

import (
	"fmt"
	"os/exec"
)

const executablePath = "/Users/sergio/bin/sqlcmd-fake"

// RunWithArgs runs sqlcmd executable with the provided arguments
func RunWithArgs(args []string) ([]byte, error) {
	return runExternalCommand(args)
}

func runExternalCommand(args []string) ([]byte, error) {
	cmd := exec.Command(executablePath, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, output)
	}
	return output, nil
}
