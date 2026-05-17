package cmd

/*
Copyright © 2026 Omkar Anil Gajare <theomkargajre@gmail.com>
*/
import (
	"chemcli/internal/core"
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all medicines in stock",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		meds, err := core.ListMedicines()
		if err != nil {
			fmt.Printf("Error listing medicines: %v\n", err)
			return
		}

		fmt.Println("#\tNAME\tQTY\tExpiry")
		for i, med := range meds {
			fmt.Printf("%d\t%s\t%d\t%s\n", i+1, med.Name, med.Quantity, med.ExpiryDate)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
