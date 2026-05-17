package cmd

/*
Copyright © 2026 Omkar Gajare <theomkargajre@gmail.com>
*/
import (
	"chemcli/internal/core"
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
			fmt.Println("Usage: chemcli search <medicine name>")
			return
		}
		searchTerm := args[0]

		meds, err := core.SearchMedicines(searchTerm)
		if err != nil {
			fmt.Printf("Error searching medicines: %v\n", err)
			return
		}

		fmt.Println("#\tID\tNAME\tQTY")

		if len(meds) == 0 {
			fmt.Println("No medicines found")
			return
		}

		for i, med := range meds {
			fmt.Printf("%d\t%d\t%s\t%d\n", i+1, med.ID, med.Name, med.Quantity)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
