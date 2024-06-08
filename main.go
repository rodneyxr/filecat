package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "filecat",
	Short: "A tool for concatenating files with filtering",
	Long: `filecat is a command-line utility that allows you to concatenate 
	the contents of files within a directory and its subdirectories.`,
}

// Global variables to store flag values
var (
	excludeDirs          []string
	excludeExts          []string
	includeDirs          []string
	includeExts          []string
	ignoreExtensionless  bool
	includeExtensionless bool
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [directory]",
	Short: "Concatenate files in a directory with filtering",
	Long: `Recursively walks through the specified directory 
	(or the current directory if none is provided) and concatenates the content of 
	files based on the provided filtering options.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		directory := "."
		if len(args) > 0 {
			directory = args[0]
		}

		// If ignoring extensionless files, ensure "" is in excludeExts
		if ignoreExtensionless && !containsExact(excludeExts, "") {
			excludeExts = append(excludeExts, "")
		}

		// If including extensionless files, ensure "" is in includeExts (and takes precedence over excludeExts)
		if includeExtensionless && !containsExact(includeExts, "") {
			includeExts = append(includeExts, "")
		}

		return filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			// Skip excluded directories
			if d.IsDir() && containsExact(excludeDirs, d.Name()) {
				return fs.SkipDir
			}

			// Read and print the file content
			if !d.IsDir() {
				// Get file extension (without the leading dot)
				ext := strings.TrimPrefix(filepath.Ext(path), ".")

				// Check if the file matches the inclusion/exclusion criteria
				includeDir := len(includeDirs) == 0 || containsExact(includeDirs, filepath.Dir(path))
				includeExt := len(includeExts) == 0 || containsExact(includeExts, ext)
				excludeExt := len(excludeExts) > 0 && containsExact(excludeExts, ext)

				if includeDir && includeExt && !excludeExt {
					content, err := os.ReadFile(path)
					if err != nil {
						return err
					}

					fmt.Printf("## FILE: \"%s\"\n", path) // Markdown header for file path
					fmt.Printf("```%s\n", ext)            // Markdown code block with file extension
					fmt.Println(string(content))
					fmt.Println("```")
					fmt.Println()
				}
			}
			return nil
		})
	},
}

// init function is used to initialize the Cobra command and flags
func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringSliceVarP(&excludeDirs, "exclude-dirs", "d", nil, "Directories to exclude")
	runCmd.Flags().StringSliceVarP(&excludeExts, "exclude-exts", "e", nil, "File extensions to exclude (without the dot)")
	runCmd.Flags().StringSliceVarP(&includeDirs, "include-dirs", "D", nil, "Directories to include")
	runCmd.Flags().StringSliceVarP(&includeExts, "include-exts", "E", nil, "File extensions to include (without the dot)")
	runCmd.Flags().BoolVarP(&ignoreExtensionless, "ignore-extensionless", "i", false, "Ignore files without extensions")
	runCmd.Flags().BoolVarP(&includeExtensionless, "include-extensionless", "I", false, "Include files without extensions")
}

// containsExact checks if a slice contains an exact string match
func containsExact(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
