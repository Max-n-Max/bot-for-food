package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const (
	versionNumber = "0.0.0"
)

var RootCmd = &cobra.Command{
	Use:   "bot",
	Short: fmt.Sprintf("USAGE %s [OPTIONS]", os.Args[0]),
	Long:  fmt.Sprintf(`USAGE %s [OPTIONS] : see --help for details`, os.Args[0]),
	Run:   executeRootCommand,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}






func executeRootCommand(cmd *cobra.Command, args []string) {
	fmt.Printf("Bot For Food v%s\n", versionNumber)
}