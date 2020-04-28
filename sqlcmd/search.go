package sqlcmd

// SearchCommand is a struct
type SearchCommand struct {
	Config
	SearchTerm string
}

// GetArgs is a func
func (c *SearchCommand) GetArgs() (map[string]string, error) {

	templateText := `
	DECLARE @Text nvarchar(4000);
	SET @Text = '{{.SearchTerm}}';
	
	-- Get the schema name, table name, and table type for:
	
	-- Table names
	SELECT
		   TABLE_SCHEMA  AS 'Object Schema'
		  ,TABLE_NAME    AS 'Object Name'
		  ,TABLE_TYPE    AS 'Object Type'
		  ,'Table Name'  AS 'TEXT Location'
	FROM  INFORMATION_SCHEMA.TABLES
	WHERE TABLE_NAME LIKE '%'+@Text+'%'
	UNION
	 --Column names
	SELECT
		  TABLE_SCHEMA   AS 'Object Schema'
		  ,COLUMN_NAME   AS 'Object Name'
		  ,'COLUMN'      AS 'Object Type'
		  ,'Column Name' AS 'TEXT Location'
	FROM  INFORMATION_SCHEMA.COLUMNS
	WHERE COLUMN_NAME LIKE '%'+@Text+'%'
	UNION
	-- Function or procedure bodies
	SELECT
		  SPECIFIC_SCHEMA     AS 'Object Schema'
		  ,ROUTINE_NAME       AS 'Object Name'
		  ,ROUTINE_TYPE       AS 'Object Type'
		  ,ROUTINE_DEFINITION AS 'TEXT Location'
	FROM  INFORMATION_SCHEMA.ROUTINES 
	WHERE ROUTINE_DEFINITION LIKE '%'+@Text+'%'
		  AND (ROUTINE_TYPE = 'function' OR ROUTINE_TYPE = 'procedure');`

	args, err := createScriptBasedSubcommand(c.Config, templateText, *c)
	if err != nil {
		return nil, err
	}
	return args, nil
}