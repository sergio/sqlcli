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
	"strings"

	"github.com/sergio/sqlcli/sqlcmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <schema-qualified-name>",
	Short: "Outputs an object's SQL definition.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("The <schema-qualified-name> argument is required")
		}

		words := strings.SplitN(args[0], ".", 2)
		if len(words) < 2 {
			return errors.New("The <schema-qualified-name> argument should be of the form <schema>.<name>")
		}

		var config sqlcmd.Config
		viper.Unmarshal(&config)

		c := &sqlcmd.GetObjectCommand{
			Config:       config,
			ObjectSchema: words[0],
			ObjectName:   words[1],
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
	rootCmd.AddCommand(getCmd)
}
