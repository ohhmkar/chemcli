package cmd

/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/

import (
	"chemcli/internal/db"
	"database/sql"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var sellName string
var sellQty int

// sellCmd represents the sell command
var sellCmd = &cobra.Command{
	Use:   "sell",
	Short: "cmd to sell stock",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if sellName == "" {
			fmt.Println("Medicine name required")
			return
		}

		if sellQty <= 0 {
			fmt.Println("Quantity must be positive")
			return
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

		var currentQty int

		err = database.QueryRow(
			"SELECT quantity FROM medicines WHERE name = ?",
			sellName,
		).Scan(&currentQty)

		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Medicine not found")
			return
		}

		if err != nil {
			panic(err)
		}

		if currentQty < sellQty {
			fmt.Printf(
				"Insufficient stock. Available: %d\n",
				currentQty,
			)
			return
		}

		_, err = database.Exec(
			`
			UPDATE medicines
			SET quantity = quantity - ?
			WHERE name = ?
			`,
			sellQty,
			sellName,
		)

		if err != nil {
			panic(err)
		}

		fmt.Printf(
			"Sold %d units of %s\n",
			sellQty,
			sellName,
		)
	},
}

func init() {
	rootCmd.AddCommand(sellCmd)

	sellCmd.Flags().StringVarP(
		&sellName, "name", "n", "", "Medicine name",
	)

	sellCmd.Flags().IntVarP(
		&sellQty,
		"qty", "q", 0, "Quantity to sell",
	)

	err := sellCmd.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}

	err = sellCmd.MarkFlagRequired("qty")
	if err != nil {
		panic(err)
	}
}
