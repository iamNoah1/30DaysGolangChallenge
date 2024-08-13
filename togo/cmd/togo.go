/*
Copyright © 2024 Noah Ispas <noahispas.public@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// togoCmd represents the togo command
var togoCmd = &cobra.Command{
	Use:   "togo",
	Short: "Root command to work with todo items",
	Long:  `Root command to work with todo items`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("togo called")
	},
}

func init() {
	rootCmd.AddCommand(togoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// togoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// togoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
