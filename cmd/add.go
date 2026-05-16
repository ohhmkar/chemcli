package cmd

/*
Copyright © 2026 Omkar Gajare <theomkargajre@gmail.com>
*/

import (
	"chemcli/internal/db"
	"fmt"

	_ "modernc.org/sqlite"

	"github.com/spf13/cobra"
)

var name string
var qty int
var expiry string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add medicines to the db",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		database, err := db.InitDB()
		if err != nil {
			panic(err)
		}
		defer database.Close()

		_, err = database.Exec(
			`
	INSERT INTO medicines(name, quantity, expiry_date) VALUES(?, ?, ?) ON CONFLICT(name,expiry_date) DO UPDATE SET quantity = quantity + excluded.quantity;
	`,
			name,
			qty,
			expiry,
		)

		if err != nil {
			panic(err)
		}
		fmt.Println("Added", name, qty)
	},
}

func init() {
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Medicine name")
	addCmd.Flags().StringVarP(&expiry, "expiry", "e", "", "Medicine Expiry (YYYY-DD-MM)")
	addCmd.Flags().IntVarP(&qty, "qty", "q", 0, "Medicine quantity")
	err := addCmd.MarkFlagRequired("qty")
	if err != nil {
		return
	}
	err = addCmd.MarkFlagRequired("name")
	if err != nil {
		return
	}
	err = addCmd.MarkFlagRequired("expiry")
	if err != nil {
		return
	}

	rootCmd.AddCommand(addCmd)
}
