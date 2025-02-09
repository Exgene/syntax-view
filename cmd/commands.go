package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/exgene/sv/internal/capture"
	"github.com/exgene/sv/internal/generate"

	"github.com/spf13/cobra"
)

var (
	inputDir     string
	outputFile   string
	markdownMode bool
)

var captureCmd = &cobra.Command{
	Use:   "capture",
	Short: "Capture code files as PDF",
	Long: `Captures screenshots of code files in the specified directory
and combines them into a single PDF document or Markdown format. Supports recursive
directory traversal and syntax highlighting.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runCapture()
	},
}

func init() {
	RootCmd.AddCommand(captureCmd)

	captureCmd.Flags().StringVarP(&inputDir, "dir", "d", ".", "Directory to capture")
	captureCmd.Flags().BoolVarP(&markdownMode, "markdown", "m", false, "Produces Output in a markdown File")
}

func runCapture() error {
	outputFile, _ = RootCmd.PersistentFlags().GetString("output")

	gen := generate.NewPDFGenerator()

	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if shouldProcessFile(path) {
			img, err := capture.Screenshot(path)
			if err != nil {
				return fmt.Errorf("failed to capture %s: %w", path, err)
			}

			if err := gen.AddPage(path, img); err != nil {
				return fmt.Errorf("failed to add page: %w", path, err)
			}

			fmt.Printf("Processed: %s\n", path)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to process directory: %w", err)
	}

	if err := gen.Save(outputFile); err != nil {
		return fmt.Errorf("failed to save PDF: %w", err)
	}

	fmt.Printf("Successfully generated PDF: %s\n", outputFile)
	return nil
}

func shouldProcessFile(path string) bool {
	ext := filepath.Ext(path)

	codeExts := map[string]bool{
		".go":   true,
		".js":   true,
		".py":   true,
		".java": true,
		".cpp":  true,
		".c":    true,
		".h":    true,
		".rs":   true,
		".rb":   true,
		".php":  true,
		".ts":   true,
	}
	return codeExts[ext]
}
