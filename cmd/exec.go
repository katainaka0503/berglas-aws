/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"github.com/katainaka0503/hogehogecli/pkg/exec"
	"github.com/spf13/cobra"
	"os"
)

var unhandleErrorExitCode = 1

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "berglas-aws is a tool inspired by GCP's berglas",
	Long: `berglas-aws is a tool inspired by GCP's berglas(https://github.com/GoogleCloudPlatform/berglas).
With this tool, you can fetch secret value's stored in AWS Secrets Manager by designate in env value as following.

export SOME_ENV_VAR="berglas-aws://<AWS Secrets Manager ARN>"`,
	RunE: execRun(),
}

func execRun() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("must specify command")
		}

		commandSubProcess := args[0]
		var argsSubProcess []string
		if len(args) == 1 {
			argsSubProcess = []string{}
		} else {
			argsSubProcess = args[1:]
		}

		if err := exec.Exec(commandSubProcess, argsSubProcess); err != nil {
			var exitError *exec.ExitError
			if errors.As(err, &exitError) {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(exitError.Code)
			} else {
				fmt.Fprintf(os.Stderr, "unhandled error: %v\n", err)
				os.Exit(unhandleErrorExitCode)
			}
		}
		return nil
	}
}

func init() {
	rootCmd.AddCommand(execCmd)
}
