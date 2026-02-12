package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shelfql",
	Short: "ShelfQL is a library inventory CLI",
	Long:  `ShelfQL is a containerized CLI engine for managing high-volume library inventories, featuring a Go-based logic core and MySQL persistence.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ShelfQL (v0.1.0) is running...")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
