/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package wordtext

import (
	"github.com/spf13/cobra"
)

// WordtextCmd represents the wordtext command
var WordtextCmd = &cobra.Command{
	Use:   "wordtext",
	Short: "Word Text is a palette for word document related commands.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(WordtextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wordtextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wordtextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
