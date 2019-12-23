package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "berglas-aws",
	Short: "berglas-aws is a tool inspired by GCP's berglas",
	Long: `berglas-aws is a tool inspired by GCP's berglas(https://github.com/GoogleCloudPlatform/berglas).
With this tool, you can fetch secret value's stored in AWS Secrets Manager by designate in env value as following.

export SOME_ENV_VAR="berglas-aws://<AWS Secrets Manager ARN>"`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
