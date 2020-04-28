package sqlcmd

import (
	"fmt"
	"os/exec"
)

const executablePath = "/Users/sergio/bin/sqlcmd-fake"

// Run runs sqlcmd executable with the provided arguments
func Run(command ArgsGetter) ([]byte, error) {
	args, err := command.GetArgs()
	if err != nil {
		return nil, err
	}

	commandLine := []string{}
	for k, v := range args {
		commandLine = append(
			commandLine,
			fmt.Sprintf(`-%s`, k),
			fmt.Sprintf(`"%s"`, v))
	}

	return runExternalCommand(commandLine)
}

func runExternalCommand(args []string) ([]byte, error) {
	cmd := exec.Command(executablePath, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, output)
	}
	return output, nil
}
