package cmd

/*
Copyright © 2026 Omkar Anil Gajare <theomkargajre@gmail.com>
*/
import (
	"chemcli/internal/db"
	"database/sql"
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
		defer func(database *sql.DB) {
			err := database.Close()
			if err != nil {
				return
			}
		}(database)

		rows, err := database.Query("SELECT id,name,quantity,expiry_date from medicines")
		if err != nil {
			panic(err)
		}
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				return
			}
		}(rows)

		fmt.Println("#\tNAME\tQTY\tExpiry")
		count := 1
		for rows.Next() {
			var id int
			var name string
			var qty int
			var expiry string

			err = rows.Scan(&id, &name, &qty, &expiry)
			if err != nil {
				panic(err)
			}

			fmt.Printf("%d\t%s\t%d\t%s\n", count, name, qty, expiry)
			count++
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
