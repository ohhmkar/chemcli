package cmd

/*
Copyright © 2026 Omkar Gajare <theomkargajre@gmail.com>
*/
import (
	"chemcli/internal/db"
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches through the db for medicine",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: chem search <medicine name>")
		}
		searchTerm := args[0]

		database, err := db.InitDB()
		if err != nil {
			panic(err)
		}
		defer func(database *sql.DB) {
			err := database.Close()
			if err != nil {

			}
		}(database)

		rows, err := database.Query(`
			SELECT id, name, quantity
			FROM medicines
			WHERE LOWER(name) LIKE LOWER(?)
		`,
			"%"+searchTerm+"%",
		)
		if err != nil {
			panic(err)
		}
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {

			}
		}(rows)
		fmt.Println("#\tID\tNAME\tQTY")

		found := false
		count := 1
		for rows.Next() {

			found = true

			var id int
			var name string
			var qty int

			err = rows.Scan(&id, &name, &qty)
			if err != nil {
				panic(err)
			}

			fmt.Printf("%d\t%d\t%s\t%d\n", count, id, name, qty)
			count++
		}

		if !found {
			fmt.Println("No medicines found")
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
