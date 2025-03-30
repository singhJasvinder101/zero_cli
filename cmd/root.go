package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:  "zero",
	Short: "Zero is a CLI tool for mathematical operations",
	Long: `Zero is a CLI tool for mathematical operations. It supports basic operations like addition, subtraction, multiplication, and division.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Zero CLI!")
	},
}

func Execute(){
	if err := rootCommand.Execute(); err != nil{
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
		os.Exit(1)
	}
}
