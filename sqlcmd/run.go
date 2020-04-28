package sqlcmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/viper"
)

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
			fmt.Sprintf(`%s`, v))
	}

	sqlCmdBinary := viper.GetString("SqlCmdBinary")
	return runExternalCommand(sqlCmdBinary, commandLine)
}

func runExternalCommand(executablePath string, args []string) ([]byte, error) {
	cmd := exec.Command(executablePath, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		commandLine := fmt.Sprintf("%s %#v", executablePath, args)
		return nil, fmt.Errorf("%v: %s\nCommand line: [%s]", err, output, commandLine)
	}
	return output, nil
}
