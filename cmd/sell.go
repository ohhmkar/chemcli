package cmd

/*
Copyright © 2026 Omkar Anil Gajare <theomkargajre@gmail.com>
*/

import (
	"chemcli/internal/core"
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

		err := core.SellMedicine(sellName, sellQty)
		if err != nil {
			if err == core.ErrNotFound {
				fmt.Println("Medicine not found")
			} else if err == core.ErrInsufficientStock {
				fmt.Printf("Insufficient stock to sell %d units of %s\n", sellQty, sellName)
			} else {
				fmt.Printf("Error selling medicine: %v\n", err)
			}
			return
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
