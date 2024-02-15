package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "wordtext"}

	var linesCmd = &cobra.Command{
		Use:   "l",
		Short: "Count lines in a file",
		Run: func(cmd *cobra.Command, args []string) {
			file, err := cmd.Flags().GetString("file")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			if file == "" {
				fmt.Println("Please provide a file path using -f or --file flag.")
				return
			}
			countLines(file)
		},
	}

	var wordsCmd = &cobra.Command{
		Use:   "w",
		Short: "Count words in a file",
		Run: func(cmd *cobra.Command, args []string) {
			file, err := cmd.Flags().GetString("file")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			if file == "" {
				fmt.Println("Please provide a file path using -f or --file flag.")
				return
			}
			countWords(file)
		},
	}

	var charactersCmd = &cobra.Command{
		Use:   "c",
		Short: "Count characters in a file",
		Run: func(cmd *cobra.Command, args []string) {
			file, err := cmd.Flags().GetString("file")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			if file == "" {
				fmt.Println("Please provide a file path using -f or --file flag.")
				return
			}
			countCharacters(file)
		},
	}
	var bytesCmd = &cobra.Command{
		Use:   "b",
		Short: "Count bytes in a file",
		Run: func(cmd *cobra.Command, args []string) {
			file, err := cmd.Flags().GetString("file")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			if file == "" {
				fmt.Println("Please provide a file path using -f or --file flag.")
				return
			}
			countBytes(file)
		},
	}

	rootCmd.AddCommand(linesCmd, wordsCmd, charactersCmd, bytesCmd)

	rootCmd.PersistentFlags().StringP("file", "f", "", "Input file")

	rootCmd.Execute()
}

func countLines(file string) {
	// Open the file
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	// Count lines
	scanner := bufio.NewScanner(f)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	fmt.Println("Lines:", lineCount)
}

func countWords(file string) {
	// Open the file
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	// Count words
	scanner := bufio.NewScanner(f)
	wordCount := 0
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}
	fmt.Println("Words:", wordCount)
}

func countCharacters(file string) {
	// Open the file
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	// Count characters
	scanner := bufio.NewScanner(f)
	charCount := 0
	for scanner.Scan() {
		charCount += len(scanner.Text())
	}
	fmt.Println("Characters:", charCount)
}

func countBytes(file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	byteCount := fileInfo.Size()
	fmt.Println("Bytes:", byteCount)
}
