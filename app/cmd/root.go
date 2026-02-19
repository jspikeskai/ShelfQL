package cmd

import (
	"ShelfQL/internal/db"
	"ShelfQL/styles"
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var DB *sql.DB
var programStarted = false

var rootCmd = &cobra.Command{
	Use:   "shelfql",
	Short: "ShelfQL is a library inventory CLI",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if DB == nil {
			DB = db.CreateHandle()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if programStarted {
			return
		}
		programStarted = true

		titleRender := styles.TitleStyle.Render("Welcome to ShelfQL!")
		quitInfoRender := styles.InfoStyle.Render("Type quit or q to exit program.")
		fmt.Printf("\n%s\n%s\n\n", titleRender, quitInfoRender)

		scanner := bufio.NewScanner(os.Stdin)

	ExitLoop:
		for {
			fmt.Print("shelfql>")

			if !scanner.Scan() {
				break
			}

			input := strings.TrimSpace(scanner.Text())
			input = strings.ToLower(input)

			switch input {
			case "quit", "q", "exit":
				break ExitLoop
			default:
				cmdArgs := strings.Fields(input)
				cmd.Root().SetArgs(cmdArgs)

				if err := cmd.Root().Execute(); err != nil {
					break
				}
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
