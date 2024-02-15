package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func main() {
	// cmd.Execute()
	var rootCmd = &cobra.Command{Use: "wordtext"}

    var linesCmd = &cobra.Command{
        Use:   "lines [file]",
        Short: "Count lines in a file",
        Run: func(cmd *cobra.Command, args []string) {
            file, err := os.Open(args[0])
            if err != nil {
                fmt.Println("Error:", err)
                return
            }
            defer file.Close()

            scanner := bufio.NewScanner(file)
            lineCount := 0
            for scanner.Scan() {
                lineCount++
            }
            fmt.Println("Lines:", lineCount)
        },
    }

    var wordsCmd = &cobra.Command{
        Use:   "words [file]",
        Short: "Count words in a file",
        Run: func(cmd *cobra.Command, args []string) {
            file, err := os.Open(args[0])
            if err != nil {
                fmt.Println("Error:", err)
                return
            }
            defer file.Close()

            scanner := bufio.NewScanner(file)
            wordCount := 0
            for scanner.Scan() {
                words := strings.Fields(scanner.Text())
                wordCount += len(words)
            }
            fmt.Println("Words:", wordCount)
        },
    }

    var charactersCmd = &cobra.Command{
        Use:   "characters [file]",
        Short: "Count characters in a file",
        Run: func(cmd *cobra.Command, args []string) {
            file, err := os.Open(args[0])
            if err != nil {
                fmt.Println("Error:", err)
                return
            }
            defer file.Close()

            scanner := bufio.NewScanner(file)
            charCount := 0
            for scanner.Scan() {
                charCount += len(scanner.Text())
            }
            fmt.Println("Characters:", charCount)
        },
    }

    rootCmd.AddCommand(linesCmd, wordsCmd, charactersCmd)

    rootCmd.Flags().BoolP("lines", "l", false, "Count lines")
    rootCmd.Flags().BoolP("words", "w", false, "Count words")
    rootCmd.Flags().BoolP("characters", "c", false, "Count characters")

    rootCmd.Run = func(cmd *cobra.Command, args []string) {
        lines, _ := cmd.Flags().GetBool("lines")
        words, _ := cmd.Flags().GetBool("words")
        characters, _ := cmd.Flags().GetBool("characters")

        if lines {
            linesCmd.Run(cmd, args)
        }
        if words {
            wordsCmd.Run(cmd, args)
        }
        if characters {
            charactersCmd.Run(cmd, args)
        }
    }

    rootCmd.Execute()
}
