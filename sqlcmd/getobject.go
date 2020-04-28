package sqlcmd

// GetObjectCommand is a struct
type GetObjectCommand struct {
	Config
	ObjectSchema string
	ObjectName   string
}

// GetArgs is a func
func (c *GetObjectCommand) GetArgs() (map[string]string, error) {

	templateText := `
	SET NOCOUNT ON
	GO
	if exists(
		select 1 
		from INFORMATION_SCHEMA.routines
		where
			routine_schema = '{{.ObjectSchema}}' AND
			routine_name = '{{.ObjectName}}'
		)
	begin
		select routine_definition 
		from INFORMATION_SCHEMA.routines
		where
			routine_schema = '{{.ObjectSchema}}' AND
			routine_name = '{{.ObjectName}}'
	end
	if exists(
		select 1
		from INFORMATION_SCHEMA.tables
		where
			table_schema = '{{.ObjectSchema}}' AND
			table_name = '{{.ObjectName}}'
		)
	begin 
		select *
			from information_schema.tables
			where
				table_schema = '{{.ObjectSchema}}' AND
				table_name = '{{.ObjectName}}'
		
		select *
			from information_schema.columns
			where table_schema = '{{.ObjectSchema}}' AND
			table_name = '{{.ObjectName}}'
	end
	`

	args, err := createScriptBasedSubcommand(c.Config, templateText, *c)
	if err != nil {
		return nil, err
	}

	args["h"] = "-1" // no headers
	return args, nil
}
