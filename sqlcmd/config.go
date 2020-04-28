package sqlcmd

// Config is a struct
type Config struct {
	Server   string
	Database string
	Username string
	Password string
}

// GetArgs is a func
func (c *Config) GetArgs() (map[string]string, error) {
	return map[string]string{
		"S": c.Server,
		"U": c.Username,
		"P": c.Password,
		"d": c.Database,
	}, nil
}
