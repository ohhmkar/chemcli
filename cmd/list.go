/*
Copyright © 2026 Omkar Anil Gajare <theomkargajre@gmail.com>
*/
package cmd

import (
	"chemcli/internal/db"
	"fmt"

	_ "modernc.org/sqlite"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all medicines in stock",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		database, err := db.InitDB()
		if err != nil {
			panic(err)
		}
		defer database.Close()

		rows, err := database.Query("SELECT id,name,quantity from medicines")
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		fmt.Println("#\tNAME\tQTY")
		count := 1
		for rows.Next() {
			var id int
			var name string
			var qty int

			err = rows.Scan(&id, &name, &qty)
			if err != nil {
				panic(err)
			}

			fmt.Printf("%d\t%s\t%d\n", count, name, qty)
			count++
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
