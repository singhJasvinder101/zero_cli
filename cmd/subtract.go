package cmd

import "github.com/spf13/cobra"


var subtractCmd = &cobra.Command{
	Use:  "subtract",
	Aliases: []string{"minus", "subtraction"},
	Short: "Subtract two numbers",
	Long: `Subtract two numbers. This command takes two arguments and returns their difference.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string){
		result := Subtract(args[0], args[1])
		cmd.Printf("Subtraction of 2 numbers %s and %s is = %s\n", args[0], args[1], result)
	},
}

func init(){
	rootCommand.AddCommand(subtractCmd)
}


