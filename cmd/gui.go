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

		tabs := container.NewAppTabs(
			container.NewTabItem("Inventory", createInventoryTab(w)),
			container.NewTabItem("Add Medicine", createAddTab(w)),
			container.NewTabItem("Sell", createSellTab(w)),
		)
		tabs.SetTabLocation(container.TabLocationTop)

		w.SetContent(tabs)
		w.ShowAndRun()
	},
}

func createInventoryTab(w fyne.Window) fyne.CanvasObject {
	meds, err := core.ListMedicines()
	if err != nil {
		return widget.NewLabel("Error loading inventory")
	}

	if len(meds) == 0 {
		return widget.NewLabel("No medicines in stock.")
	}

	data := [][]string{{"ID", "Name", "Quantity", "Expiry"}}
	for _, m := range meds {
		data = append(data, []string{
			fmt.Sprintf("%d", m.ID),
			m.Name,
			fmt.Sprintf("%d", m.Quantity),
			m.ExpiryDate,
		})
	}

	list := widget.NewTable(
		func() (int, int) { return len(data), len(data[0]) },
		func() fyne.CanvasObject { return widget.NewLabel("Wide Content Here") },
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		},
	)

	return container.NewPadded(list)
}

func createAddTab(w fyne.Window) fyne.CanvasObject {
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
		},
	}

	return container.NewVBox(
		widget.NewLabel("Add New Medicine"),
		form,
	)
}

func createSellTab(w fyne.Window) fyne.CanvasObject {
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
