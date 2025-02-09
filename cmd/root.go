package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "syntax-view",
	Short: "A documentation generator that captures code screenshots",
	Long: `A CLI tool that generates PDF documentation by capturing screenshots 
of your code files or markdown files. It walks through directories recursively
and creates a well-formatted PDF with syntax highlighting.`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringP("output", "o", "documentation.pdf", "Output PDF file path")
}
