/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"os"

	"github.com/sergio/sqlcli/sqlcmd"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query <sql-statement>",
	Short: "Runs a SQL statement",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) < 1 {
			return errors.New("The <sql-statement> argument is required")
		}

		var config sqlcmd.Config
		viper.Unmarshal(&config)

		c := &sqlcmd.QueryCommand{
			Config:       config,
			SQLStatement: args[0],
		}

		result, err := sqlcmd.Run(c)
		if err != nil {
			return err
		}

		os.Stdout.Write(result)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}
