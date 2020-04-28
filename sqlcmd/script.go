package sqlcmd

// ScriptCommand is a struct
type ScriptCommand struct {
	Config
	InputFile string
}

// GetArgs is a func
func (c *ScriptCommand) GetArgs() (map[string]string, error) {
	args, err := c.Config.GetArgs()
	if err != nil {
		return nil, err
	}
	args["i"] = c.InputFile
	return args, nil
}
