package sqlcmd

import (
	"bytes"
	"html/template"
	"io/ioutil"
)

// ArgsGetter is an interface specifying a method that generates command line args
type ArgsGetter interface {
	GetArgs() (map[string]string, error)
}

func createScriptBasedSubcommand(config Config, templateText string, data interface{}) (map[string]string, error) {

	tmpl, err := template.New("sql-statement").Parse(templateText)
	if err != nil {
		return nil, err
	}
	var sqlStatement bytes.Buffer
	err = tmpl.Execute(&sqlStatement, data)
	if err != nil {
		return nil, err
	}

	scriptFile, err := writeTempFile(sqlStatement.Bytes())
	if err != nil {
		return nil, err
	}

	subcommand := ScriptCommand{
		Config:    config,
		InputFile: scriptFile,
	}
	args, err := subcommand.GetArgs()
	if err != nil {
		return nil, err
	}
	return args, nil
}

func writeTempFile(content []byte) (string, error) {
	tmpfile, err := ioutil.TempFile("", "sqlcli.*.sql")
	if err != nil {
		return "", err
	}

	if _, err := tmpfile.Write(content); err != nil {
		tmpfile.Close()
		return "", err
	}
	if err := tmpfile.Close(); err != nil {
		return "", err
	}
	return tmpfile.Name(), nil
}
