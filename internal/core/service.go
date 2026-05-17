package core

import (
	"chemcli/internal/db"
	"database/sql"
	"errors"
)

// Medicine represents a single medicine record
type Medicine struct {
	ID         int
	Name       string
	Quantity   int
	ExpiryDate string
}

var (
	ErrNotFound          = errors.New("medicine not found")
	ErrInsufficientStock = errors.New("insufficient stock")
)

// AddMedicine adds a new medicine or updates quantity if it exists
func AddMedicine(name string, qty int, expiry string) error {
	database, err := db.InitDB()
	if err != nil {
		return err
	}
	defer database.Close()

	_, err = database.Exec(
		`INSERT INTO medicines(name, quantity, expiry_date) VALUES(?, ?, ?) 
		ON CONFLICT(name,expiry_date) DO UPDATE SET quantity = quantity + excluded.quantity;`,
		name, qty, expiry,
	)
	return err
}

// ListMedicines retrieves all medicines in stock
func ListMedicines() ([]Medicine, error) {
	database, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	defer database.Close()

	rows, err := database.Query("SELECT id, name, quantity, expiry_date FROM medicines")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meds []Medicine
	for rows.Next() {
		var med Medicine
		err = rows.Scan(&med.ID, &med.Name, &med.Quantity, &med.ExpiryDate)
		if err != nil {
			return nil, err
		}
		meds = append(meds, med)
	}
	return meds, nil
}

// searchMedicines returns all medicines with the same name
func SearchMedicines(searchTerm string) ([]Medicine, error) {
	database, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	defer database.Close()

	rows, err := database.Query(
		"SELECT id, name, quantity FROM medicines WHERE LOWER(name) LIKE LOWER(?)",
		"%"+searchTerm+"%",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meds []Medicine
	for rows.Next() {
		var med Medicine
		err = rows.Scan(&med.ID, &med.Name, &med.Quantity)
		if err != nil {
			return nil, err
		}
		meds = append(meds, med)
	}
	return meds, nil
}

// lowStock returns medicines below a threshold
func LowStockMedicines(threshold int) ([]Medicine, error) {
	database, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	defer database.Close()

	rows, err := database.Query(
		"SELECT id, name, quantity FROM medicines WHERE quantity <= ? ORDER BY quantity ASC",
		threshold,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meds []Medicine
	for rows.Next() {
		var med Medicine
		err = rows.Scan(&med.ID, &med.Name, &med.Quantity)
		if err != nil {
			return nil, err
		}
		meds = append(meds, med)
	}
	return meds, nil
}

// SellMedicine reduces the quantity of a medicine by name
func SellMedicine(name string, qtyToSell int) error {
	database, err := db.InitDB()
	if err != nil {
		return err
	}
	defer database.Close()

	var currentQty int
	err = database.QueryRow("SELECT quantity FROM medicines WHERE name = ?", name).Scan(&currentQty)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFound
		}
		return err
	}

	if currentQty < qtyToSell {
		return ErrInsufficientStock
	}

	_, err = database.Exec("UPDATE medicines SET quantity = quantity - ? WHERE name = ?", qtyToSell, name)
	return err
}
