package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "calc",
		Short: "A simple calculator",
	}

	//Addition
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add two numbers",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			x, y := parseArgs(args)
			add(x, y)
		},
	}

	//Subtraction
	subCmd := &cobra.Command{
		Use:   "sub",
		Short: "Substracts two numbers",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			x, y := parseArgs(args)
			sub(x, y)
		},
	}

	//Division
	divCmd := &cobra.Command{
		Use:   "div",
		Short: "Divides two numbers",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			x, y := parseArgs(args)
			div(x, y)
		},
	}

	//Multipication
	mulCmd := &cobra.Command{
		Use:   "mul",
		Short: "Multiplies two numbers",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			x, y := parseArgs(args)
			div(x, y)
		},
	}

	// Add subcommands
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(subCmd)
	rootCmd.AddCommand(divCmd)
	rootCmd.AddCommand(mulCmd)

	rootCmd.Execute()

}

func add(x, y float64) {
	printAnwser(x+y, "+")
}

func sub(x, y float64) {
	printAnwser(x-y, "-")
}

func mul(x, y float64) {
	printAnwser(x*y, "x")
}

func div(x, y float64) {
	printAnwser(x/y, "/")
}

func printAnwser(value float64, operator string) {
	fmt.Printf("A %v B = %v\n", operator, value)
}

func parseArgs(args []string) (float64, float64) {
	x, y := 0.0, 0.0
	fmt.Sscanf(args[0], "%f", &x)
	fmt.Sscanf(args[1], "%f", &y)
	return x, y
}
