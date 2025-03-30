package cmd

import "github.com/spf13/cobra"

var isRoundUp bool
var multiplyCmd = &cobra.Command{
	Use:   "multiply",
	Aliases: []string{"times", "multiplication"},
	Short: "Multiply two numbers",
	Long: `Multiply two numbers. This command takes two arguments and returns their product.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string){
		result := Multiply(args[0], args[1], isRoundUp)
		cmd.Printf("Multiplication of 2 numbers %s and %s is = %s\n", args[0], args[1], result)
	},
}


func init(){
	multiplyCmd.Flags().BoolVarP(&isRoundUp, "round-up", "r", false, "Round up the result")
	rootCommand.AddCommand(multiplyCmd)
}


