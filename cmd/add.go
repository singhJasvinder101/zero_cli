package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)


var addCmd = &cobra.Command{
	Use:   "add",
	Aliases: []string{"plus", "addition"},
	Short: "Add two numbers",
	Long: `Add two numbers. This command takes two arguments and returns their sum.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string){
		fmt.Printf("Addition of 2 numbers %s and %s is = %s", args[0], args[1], Add(args[0], args[1]))
	},
}

func init(){
	rootCommand.AddCommand(addCmd)
}

