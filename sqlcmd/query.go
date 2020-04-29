package sqlcmd

// QueryCommand is a struct
type QueryCommand struct {
	Config
	SQLStatement string
}

// GetArgs is a func
func (c *QueryCommand) GetArgs() (map[string]string, error) {
	args, err := c.Config.GetArgs()
	if err != nil {
		return nil, err
	}
	args["Q"] = c.SQLStatement
	args["Y"] = "80"
	args["y"] = "80"
	return args, nil
}
