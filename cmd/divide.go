package cmd

import "github.com/spf13/cobra"


var divideCmd = &cobra.Command{
	Use:   "divide",
	Aliases: []string{"division", "div"},
	Short: "Divide two numbers",
	Long: `Divide two numbers. This command takes two arguments and returns their quotient.`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := Divide(args[0], args[1], isRoundUp)
		if err != nil {
			return err
		}
		cmd.Printf("Division of 2 numbers %s and %s is = %s\n", args[0], args[1], result)
		return nil
	},
}

func init(){
	divideCmd.Flags().BoolVarP(&isRoundUp, "round-up", "r", false, "Round up the result")
	rootCommand.AddCommand(divideCmd)
}

