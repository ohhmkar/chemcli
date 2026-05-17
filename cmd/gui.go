package cmd

import (
	"fmt"
	"strconv"

	"chemcli/internal/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/cobra"
)

var guiCmd = &cobra.Command{
	Use:   "gui",
	Short: "Launch the native desktop companion app",
	Run: func(cmd *cobra.Command, args []string) {
		a := app.New()
		w := a.NewWindow("ChemCLI - Inventory Manager")
		w.Resize(fyne.NewSize(600, 400))

		var inventoryData [][]string

		list := widget.NewTable(
			func() (int, int) { return len(inventoryData), len(inventoryData[0]) },
			func() fyne.CanvasObject { return widget.NewLabel("Wide Content Here") },
			func(i widget.TableCellID, o fyne.CanvasObject) {
				if i.Row < len(inventoryData) && i.Col < len(inventoryData[0]) {
					o.(*widget.Label).SetText(inventoryData[i.Row][i.Col])
				}
			},
		)

		refreshInventory := func() {
			meds, err := core.ListMedicines()
			inventoryData = [][]string{{"ID", "Name", "Quantity", "Expiry"}}
			if err == nil {
				for _, m := range meds {
					inventoryData = append(inventoryData, []string{
						fmt.Sprintf("%d", m.ID),
						m.Name,
						fmt.Sprintf("%d", m.Quantity),
						m.ExpiryDate,
					})
				}
			}
			list.Refresh()
		}

		refreshInventory()

		tabs := container.NewAppTabs(
			container.NewTabItem("Inventory", container.NewPadded(list)),
			container.NewTabItem("Add Medicine", createAddTab(w, refreshInventory)),
			container.NewTabItem("Sell", createSellTab(w, refreshInventory)),
		)
		tabs.SetTabLocation(container.TabLocationTop)

		w.SetContent(tabs)
		w.ShowAndRun()
	},
}

func createAddTab(w fyne.Window, onSuccess func()) fyne.CanvasObject {
	nameEntry := widget.NewEntry()
	qtyEntry := widget.NewEntry()
	expiryEntry := widget.NewEntry()
	expiryEntry.SetPlaceHolder("YYYY-MM-DD")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: nameEntry},
			{Text: "Quantity", Widget: qtyEntry},
			{Text: "Expiry Date", Widget: expiryEntry},
		},
		OnSubmit: func() {
			q, err := strconv.Atoi(qtyEntry.Text)
			if err != nil {
				dialog.ShowError(fmt.Errorf("invalid quantity"), w)
				return
			}
			err = core.AddMedicine(nameEntry.Text, q, expiryEntry.Text)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			dialog.ShowInformation("Success", "Medicine added!", w)
			nameEntry.SetText("")
			qtyEntry.SetText("")
			expiryEntry.SetText("")
			
			// Refresh inventory table
			onSuccess()
		},
	}

	return container.NewVBox(
		widget.NewLabel("Add New Medicine"),
		form,
	)
}

func createSellTab(w fyne.Window, onSuccess func()) fyne.CanvasObject {
	nameEntry := widget.NewEntry()
	qtyEntry := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: nameEntry},
			{Text: "Quantity", Widget: qtyEntry},
		},
		OnSubmit: func() {
			q, err := strconv.Atoi(qtyEntry.Text)
			if err != nil {
				dialog.ShowError(fmt.Errorf("invalid quantity"), w)
				return
			}
			err = core.SellMedicine(nameEntry.Text, q)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			dialog.ShowInformation("Success", "Medicine sold!", w)
			nameEntry.SetText("")
			qtyEntry.SetText("")
			
			// Refresh inventory table
			onSuccess()
		},
	}

	return container.NewVBox(
		widget.NewLabel("Sell Medicine"),
		form,
	)
}

func init() {
	rootCmd.AddCommand(guiCmd)
}
