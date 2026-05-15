/*
Copyright © 2026 Omkar Gajare <theomkargajre@gmail.com>
*/
package cmd

import (
	"chemcli/internal/db"
	"fmt"

	_ "modernc.org/sqlite"

	"github.com/spf13/cobra"
)

var name string
var qty int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add medicines to the db",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := db.InitDB()
		if name == "" {
			fmt.Println("Medicine name required")
			return
		}

		if qty <= 0 {
			fmt.Println("Quantity must be positive")
			return
		}
		_, err = db.Exec(
			"INSERT INTO medicines(name, quantity) VALUES(?, ?) ON CONFLICT(name) DO UPDATE SET quantity = quantity + excluded.quantity;",
			name,
			qty,
		)

		if err != nil {
			panic(err)
		}
		fmt.Println("Added", name, qty)
	},
}

func init() {
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Medicine name")
	addCmd.Flags().IntVarP(&qty, "qty", "q", 0, "Medicine quantity")
	err := addCmd.MarkFlagRequired("qty")
	if err != nil {
		return
	}
	err = addCmd.MarkFlagRequired("name")
	if err != nil {
		return
	}

	rootCmd.AddCommand(addCmd)
}
