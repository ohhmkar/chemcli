package cmd

/*
Copyright © 2026 Omkar Anil Gajare <theomkargajre@gmail.com>
*/

import (
	"chemcli/internal/core"
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
			return
		}

		meds, err := core.LowStockMedicines(thresholdQty)
		if err != nil {
			fmt.Printf("Error fetching low stock: %v\n", err)
			return
		}

		fmt.Println("LOW STOCK MEDICINES")
		fmt.Println("ID\tNAME\tQTY")

		if len(meds) == 0 {
			fmt.Println("No low stock medicine found")
			return
		}

		for _, med := range meds {
			fmt.Printf("%d\t%s\t%d\n", med.ID, med.Name, med.Quantity)
		}
	},
}

func init() {
	rootCmd.AddCommand(lowStockCmd)
	lowStockCmd.Flags().IntVarP(
		&thresholdQty, "threshold", "t", 0, "quantity")

}
