// /*
// Copyright © 2026 NAME HERE <EMAIL ADDRESS>

// */
// package cmd

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"

// 	"goku/converter"

// 	"github.com/spf13/cobra"
// )


// // rootCmd represents the base command when called without any subcommands
// var rootCmd = &cobra.Command{
// 	Use:   "goku",
// 	Short: "A brief description of your application",
// 	Long: `A longer description that spans multiple lines and likely contains
// examples and usage of using your application. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	// Uncomment the following line if your bare application
// 	// has an action associated with it:
// 	// Run: func(cmd *cobra.Command, args []string) { },
// }

// // Execute adds all child commands to the root command and sets flags appropriately.
// // This is called by main.main(). It only needs to happen once to the rootCmd.
// func Execute() {
// 	err := rootCmd.Execute()
// 	if err != nil {
// 		os.Exit(1)
// 	}
// }

// func init() {
// 	// Here you will define your flags and configuration settings.
// 	// Cobra supports persistent flags, which, if defined here,
// 	// will be global for your application.

// 	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goku.yaml)")

// 	// Cobra also supports local flags, which will only run
// 	// when this action is called directly.
// 	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }


package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"goku/converter"

	"github.com/spf13/cobra"
)

var inputFile string
var outputFile string

var rootCmd = &cobra.Command{
	Use:   "goku",
	Short: "JSON <-> YAML converter",
	Run: func(cmd *cobra.Command, args []string) {

		inputExt := filepath.Ext(inputFile)
		outputExt := filepath.Ext(outputFile)

		// Edge Case Check
		if inputExt == outputExt {
			fmt.Println("Error: Requested output must be in different format than input file")
			os.Exit(1)
		}

		// JSON -> YAML
		if inputExt == ".json" && (outputExt == ".yaml" || outputExt == ".yml") {
			err := converter.JSONToYAML(inputFile, outputFile)

			if err != nil {
				fmt.Println("Conversion Error:", err)
				return
			}

			fmt.Println("Successfully converted JSON to YAML")
			return
		}

		// YAML -> JSON
		if (inputExt == ".yaml" || inputExt == ".yml") && outputExt == ".json" {
			err := converter.YAMLToJSON(inputFile, outputFile)

			if err != nil {
				fmt.Println("Conversion Error:", err)
				return
			}

			fmt.Println("Successfully converted YAML to JSON")
			return
		}

		fmt.Println("Unsupported conversion format")
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file")

	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("output")
}