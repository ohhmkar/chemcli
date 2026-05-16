package cmd

/*
Copyright © 2026 Omkar Anil Gajare <theomkargajre@gmail.com>
*/

import (
	"chemcli/internal/db"
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"
)

var thresholdQty int

// lowStockCmd represents the lowStock command
var lowStockCmd = &cobra.Command{
	Use:   "lowStock",
	Short: "lists all medicines with quantity below a threshold",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if thresholdQty <= 0 {
			fmt.Println("Threshold should be positive")
		}

		database, err := db.InitDB()
		if err != nil {
			panic(err)
		}
		defer func(database *sql.DB) {
			err := database.Close()
			if err != nil {

			}
		}(database)

		rows, err := database.Query(
			`SELECT id, name, quantity
FROM medicines
where quantity <= ?
ORDER BY quantity ASC 
`, thresholdQty)

		if err != nil {
			panic(err)
		}
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {

			}
		}(rows)

		fmt.Println("LOW STOCK MEDICINES")
		fmt.Println("ID\tNAME\tQTY")

		found := false

		for rows.Next() {

			found = true

			var id int
			var name string
			var qty int

			err = rows.Scan(&id, &name, &qty)
			if err != nil {
				panic(err)
			}

			fmt.Printf(
				"%d\t%s\t%d\n",
				id,
				name,
				qty,
			)
		}

		if !found {
			fmt.Println("No low stock medicine found")
		}
	},
}

func init() {
	rootCmd.AddCommand(lowStockCmd)
	lowStockCmd.Flags().IntVarP(
		&thresholdQty, "threshold", "t", 0, "quantity")

}
