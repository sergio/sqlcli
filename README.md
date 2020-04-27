# SQL Server CLI utility

This CLI provides some quick shortcuts for frequent workflows by wrapping Microsoft's sqlcmd.

## Commands 

*query*: run arbitrary sql against default instance. 

    sqlcli query "select * from table"

*script*: run statements from a .sql file

    sqlcli script myscript.sql

*get*: outputs the definition of a table, view, procedure or function

    sqlcli get dbo.myprocedure

    sqlcli get MyTable

*find*: finds objects (tables, views, stored procs or functions) containing text

    sqlcli find "text or phrase"
